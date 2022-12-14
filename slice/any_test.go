package slice

import (
	"testing"
)

func TestAny(t *testing.T) {
	type args struct {
		source []int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source: nil,
			},
			want: false,
		},
		{name: "EmptySource",
			args: args{
				source: []int{},
			},
			want: false,
		},
		{name: "NonEmptySource",
			args: args{
				source: []int{0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Any(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Any() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyPred_int(t *testing.T) {
	type args struct {
		source    []int
		predicate func(int) bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source:    nil,
				predicate: nil,
			},
			want: false,
		},
		{name: "EmptySource",
			args: args{
				source:    []int{},
				predicate: func(x int) bool { return x > 10 },
			},
			want: false,
		},
		{name: "NilPredicate",
			args: args{
				source:    []int{1, 3, 5},
				predicate: nil,
			},
			wantErr: true,
		},
		{name: "PredicateMatchingElement",
			args: args{
				source:    []int{1, 5, 20, 30},
				predicate: func(x int) bool { return x > 10 },
			},
			want: true,
		},
		{name: "PredicateNotMatchingElement",
			args: args{
				source:    []int{1, 5, 8, 9},
				predicate: func(x int) bool { return x > 10 },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AnyPred(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnyPred() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AnyPred() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyPred_any(t *testing.T) {
	type args struct {
		source    []any
		predicate func(any) bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "1",
			args: args{
				source:    []any{1, 2, 3, 4},
				predicate: func(e any) bool { return e.(int) == 2 },
			},
			want: true,
		},
		{name: "2",
			args: args{
				source:    []any{"one", "two", "three", "four"},
				predicate: func(e any) bool { return len(e.(string)) == 4 },
			},
			want: true,
		},
		{name: "3",
			args: args{
				source:    []any{1, 2, "three", "four"},
				predicate: func(e any) bool { _, ok := e.(string); return ok },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AnyPred(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnyPred() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AnyPred() = %v, want %v", got, tt.want)
			}
		})
	}
}
