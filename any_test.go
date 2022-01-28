//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AnyTest.cs

func Test_AnyMust_int(t *testing.T) {
	type args struct {
		source    Enumerable[int]
		predicate func(int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "EmptySequenceWithoutPredicate",
			args: args{
				source: Empty[int](),
			},
			want: false,
		},
		{name: "NonEmptySequenceWithoutPredicate",
			args: args{
				source: NewEnSlice(0),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnyMust(tt.args.source)
			if got != tt.want {
				t.Errorf("AnyMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AnyPred_int(t *testing.T) {
	type args struct {
		source    Enumerable[int]
		predicate func(int) bool
	}
	tests := []struct {
		name        string
		args        args
		want        bool
		wantErr     bool
		expectedErr error
	}{
		{name: "NullPredicate",
			args: args{
				source: NewEnSlice(1, 3, 5),
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "EmptySequenceWithPredicate",
			args: args{
				source:    Empty[int](),
				predicate: func(x int) bool { return x > 10 },
			},
			want: false,
		},
		{name: "NonEmptySequenceWithPredicateMatchingElement",
			args: args{
				source:    NewEnSlice(1, 5, 20, 30),
				predicate: func(x int) bool { return x > 10 },
			},
			want: true,
		},
		{name: "NonEmptySequenceWithPredicateNotMatchingElement",
			args: args{
				source:    NewEnSlice(1, 5, 8, 9),
				predicate: func(x int) bool { return x > 10 },
			},
			want: false,
		},
		{name: "SequenceIsNotEvaluatedAfterFirstMatch",
			args: args{
				source:    SelectMust(NewEnSlice(10, 2, 0, 3), func(x int) int { return 10 / x }),
				predicate: func(y int) bool { return y > 2 },
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
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("AnyPred() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("AnyPred() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AnyPredMust_interface(t *testing.T) {
	type args struct {
		source    Enumerable[any]
		predicate func(any) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1",
			args: args{
				source:    NewEnSlice[any](1, 2, 3, 4),
				predicate: func(e any) bool { return e.(int) == 4 },
			},
			want: true,
		},
		{name: "2",
			args: args{
				source:    NewEnSlice[any]("one", "two", "three", "four"),
				predicate: func(e any) bool { return len(e.(string)) == 4 },
			},
			want: true,
		},
		{name: "3",
			args: args{
				source:    NewEnSlice[any](1, 2, "three", "four"),
				predicate: func(e any) bool { _, ok := e.(int); return ok },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnyPredMust(tt.args.source, tt.args.predicate)
			if got != tt.want {
				t.Errorf("AnyPredMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
