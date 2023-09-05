package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAggregate(t *testing.T) {
	type args struct {
		source      []int
		accumulator func(int, int) int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source:      nil,
				accumulator: func(ag, el int) int { return ag + el },
			},
			want: 0,
		},
		{name: "EmptySource",
			args: args{
				source:      []int{},
				accumulator: func(ag, el int) int { return ag + el },
			},
			want: 0,
		},
		{name: "NilAccumulator",
			args: args{
				source:      []int{1, 3},
				accumulator: nil,
			},
			wantErr: true,
		},
		{name: "SingleElementSource",
			args: args{
				source:      []int{1},
				accumulator: func(ag, el int) int { return ag*2 + el },
			},
			want: 1,
		},
		{name: "1",
			args: args{
				source:      []int{1, 4, 5},
				accumulator: func(ag, el int) int { return ag*2 + el },
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Aggregate(tt.args.source, tt.args.accumulator)
			if (err != nil) != tt.wantErr {
				t.Errorf("Aggregate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Aggregate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregateSeed_int_int(t *testing.T) {
	type args struct {
		source      []int
		seed        int
		accumulator func(int, int) int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source:      nil,
				seed:        5,
				accumulator: func(ac, el int) int { return ac + el },
			},
			want: 5,
		},
		{name: "EmptySource",
			args: args{
				source:      []int{},
				seed:        5,
				accumulator: func(ac, el int) int { return ac + el },
			},
			want: 5,
		},
		{name: "NilAccumulator",
			args: args{
				source:      []int{1, 3},
				seed:        5,
				accumulator: nil,
			},
			wantErr: true,
		},
		{name: "1",
			args: args{
				source:      []int{1, 4, 5},
				seed:        5,
				accumulator: func(ac, el int) int { return ac*2 + el },
			},
			want: 57,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateSeed(tt.args.source, tt.args.seed, tt.args.accumulator)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateSeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregateSeed_int_string(t *testing.T) {
	type args struct {
		source      []int
		seed        string
		accumulator func(string, int) string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "DifferentSourceAndAccumulatorTypes",
			args: args{
				source:      []int{1, 2, 3, 4},
				seed:        "0",
				accumulator: func(ac string, el int) string { return ac + fmt.Sprint(el) },
			},
			want: "01234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateSeed(tt.args.source, tt.args.seed, tt.args.accumulator)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateSeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregateSeedSel_int_int_string(t *testing.T) {
	type args struct {
		source         []int
		seed           int
		accumulator    func(int, int) int
		resultSelector func(int) string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source:         nil,
				seed:           5,
				accumulator:    func(ac, el int) int { return ac + el },
				resultSelector: func(r int) string { return fmt.Sprint(r) },
			},
			want: "5",
		},
		{name: "EmptySource",
			args: args{
				source:         []int{},
				seed:           5,
				accumulator:    func(ac, el int) int { return ac + el },
				resultSelector: func(r int) string { return fmt.Sprint(r) },
			},
			want: "5",
		},
		{name: "NilAccumulator",
			args: args{
				source:         []int{1, 4, 5},
				seed:           5,
				accumulator:    nil,
				resultSelector: func(r int) string { return fmt.Sprint(r) },
			},
			wantErr: true,
		},
		{name: "NilResultSelector",
			args: args{
				source:         []int{1, 4, 5},
				seed:           5,
				accumulator:    func(ac, el int) int { return ac + el },
				resultSelector: nil,
			},
			wantErr: true,
		},
		{name: "1",
			args: args{
				source:         []int{1, 4, 5},
				seed:           5,
				accumulator:    func(ac, el int) int { return ac*2 + el },
				resultSelector: func(r int) string { return fmt.Sprint(r) },
			},
			want: "57",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateSeedSel(tt.args.source, tt.args.seed, tt.args.accumulator, tt.args.resultSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateSeedSel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateSeedSel() = %v, want %v", got, tt.want)
			}
		})
	}
}
