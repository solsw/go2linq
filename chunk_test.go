//go:build go1.18

package go2linq

import (
	"testing"
)

func TestChunk_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
		size   int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[[]int]
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
			want: NewEnSlice([][]int{}...),
		},
		{name: "1",
			args: args{
				source: NewEnSlice(1, 2),
				size:   2,
			},
			want: NewEnSlice([]int{1, 2}),
		},
		{name: "2",
			args: args{
				source: NewEnSlice(1, 2, 3),
				size:   2,
			},
			want: NewEnSlice([]int{1, 2}, []int{3}),
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
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}
