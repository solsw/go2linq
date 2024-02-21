package go2linq

import (
	"fmt"
	"iter"
	"testing"
)

func TestChunk_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
		size   int
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[[]int]
		wantErr     bool
		expectedErr error
	}{
		{name: "01",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "02",
			args: args{
				size: 2,
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "03",
			args: args{
				source: Empty[int](),
				size:   0,
			},
			wantErr:     true,
			expectedErr: ErrSizeOutOfRange,
		},
		{name: "EmptySource",
			args: args{
				source: Empty[int](),
				size:   2,
			},
			want: SliceAll([][]int{}),
		},
		{name: "1",
			args: args{
				source: VarAll(1, 2),
				size:   2,
			},
			want: VarAll([]int{1, 2}),
		},
		{name: "2",
			args: args{
				source: VarAll(1, 2, 3),
				size:   2,
			},
			want: VarAll([]int{1, 2}, []int{3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Chunk(tt.args.source, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chunk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Chunk() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Chunk() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/partitioning-data#example
func ExampleChunk() {
	chunkNumber := 0
	rng, _ := Range(0, 8)
	chunk, _ := Chunk(rng, 3)
	for ii := range chunk {
		chunkNumber++
		fmt.Printf("Chunk %d:%v\n", chunkNumber, ii)
	}
	// Output:
	// Chunk 1:[0 1 2]
	// Chunk 2:[3 4 5]
	// Chunk 3:[6 7]
}
