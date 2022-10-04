//go:build go1.18

package slice

import (
	"reflect"
	"testing"
)

func TestWhere_int(t *testing.T) {
	type args struct {
		source    []int
		predicate func(int) bool
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source:    nil,
				predicate: func(i int) bool { return i > 5 },
			},
			want: nil,
		},
		{name: "EmptySource",
			args: args{
				source:    []int{},
				predicate: func(i int) bool { return i > 5 },
			},
			want: []int{},
		},
		{name: "NilPredicate",
			args: args{
				source:    []int{1, 3, 4, 2, 8, 1},
				predicate: nil,
			},
			wantErr: true,
		},
		{name: "SimpleFiltering",
			args: args{
				source:    []int{1, 3, 4, 2, 8, 1},
				predicate: func(i int) bool { return i < 4 },
			},
			want: []int{1, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Where(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("Where() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Where() = %v, want %v", got, tt.want)
			}
		})
	}
}
