package slice

import (
	"testing"
)

func TestAll_int(t *testing.T) {
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
				predicate: func(x int) bool { return x > 10 },
			},
			want: true,
		},
		{name: "EmptySource",
			args: args{
				source:    []int{},
				predicate: func(x int) bool { return x > 10 },
			},
			want: true,
		},
		{name: "NilPredicate",
			args: args{
				source:    []int{1, 3, 5},
				predicate: nil,
			},
			wantErr: true,
		},
		{name: "PredicateMatchingNoElements",
			args: args{
				source:    []int{1, 5, 20, 30},
				predicate: func(x int) bool { return x < 0 },
			},
			want: false,
		},
		{name: "PredicateMatchingSomeElements",
			args: args{
				source:    []int{1, 5, 8, 9},
				predicate: func(x int) bool { return x > 3 },
			},
			want: false,
		},
		{name: "PredicateMatchingAllElements",
			args: args{
				source:    []int{1, 5, 8, 9},
				predicate: func(x int) bool { return x > 0 },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := All(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll_any(t *testing.T) {
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
				source:    []any{"one", "two", "three", "four"},
				predicate: func(e any) bool { return len(e.(string)) >= 3 },
			},
			want: true,
		},
		{name: "2",
			args: args{
				source:    []any{"one", "two", "three", "four"},
				predicate: func(e any) bool { return len(e.(string)) > 3 },
			},
			want: false,
		},
		{name: "3",
			args: args{
				source:    []any{1, 2, "three", "four"},
				predicate: func(e any) bool { _, ok := e.(int); return ok },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := All(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
