package slice

import (
	"reflect"
	"testing"
)

func TestChunk_int(t *testing.T) {
	type args struct {
		source []int
		size   int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{name: "NilSource1",
			want: nil,
		},
		{name: "NilSource2",
			args: args{
				size: 2,
			},
			want: nil,
		},
		{name: "EmptySource",
			args: args{
				source: []int{},
				size:   2,
			},
			want: [][]int{},
		},
		{name: "1",
			args: args{
				source: []int{1, 2},
				size:   2,
			},
			want: [][]int{{1, 2}},
		},
		{name: "2",
			args: args{
				source: []int{1, 2, 3},
				size:   2,
			},
			want: [][]int{{1, 2}, {3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Chunk(tt.args.source, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chunk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
