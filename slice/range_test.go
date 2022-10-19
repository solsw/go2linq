//go:build go1.18

package slice

import (
	"math"
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	type args struct {
		start int
		count int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "NegativeCount",
			args: args{
				start: 10,
				count: -1,
			},
			wantErr: true,
		},
		{name: "ValidRange",
			args: args{
				start: 5,
				count: 3,
			},
			want: []int{5, 6, 7},
		},
		{name: "NegativeStart",
			args: args{
				start: -2,
				count: 5,
			},
			want: []int{-2, -1, 0, 1, 2},
		},
		{name: "SingleValueOfMaxInt32",
			args: args{
				start: math.MaxInt32,
				count: 1,
			},
			want: []int{math.MaxInt32},
		},
		{name: "EmptyRange",
			args: args{
				start: 100,
				count: 0,
			},
			want: []int{},
		},
		{name: "EmptyRangeStartingAtMinInt32",
			args: args{
				start: math.MinInt32,
				count: 0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Range(tt.args.start, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("Range() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Range() = %v, want %v", got, tt.want)
			}
		})
	}
}
