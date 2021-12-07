//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AnyTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AllTest.cs

func Test_Any_int(t *testing.T) {
	type args struct {
		source    Enumerator[int]
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
				source: NewOnSlice(0),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Any(tt.args.source); got != tt.want {
				t.Errorf("Any() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_AnyPred_int(t *testing.T) {
	type args struct {
		source    Enumerator[int]
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
				source: NewOnSlice(1, 3, 5),
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
				source:    NewOnSlice(1, 5, 20, 30),
				predicate: func(x int) bool { return x > 10 },
			},
			want: true,
		},
		{name: "NonEmptySequenceWithPredicateNotMatchingElement",
			args: args{
				source:    NewOnSlice(1, 5, 8, 9),
				predicate: func(x int) bool { return x > 10 },
			},
			want: false,
		},
		{name: "SequenceIsNotEvaluatedAfterFirstMatch",
			args: args{
				source:    SelectMust(NewOnSliceEn(10, 2, 0, 3), func(x int) int { return 10 / x }),
				predicate: func(y int) bool { return y > 2 },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AnyPred(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnyPred() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("AnyPred() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("AnyPred() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_AnyPred_interface(t *testing.T) {
	type args struct {
		source    Enumerator[interface{}]
		predicate func(interface{}) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1",
			args: args{
				source:    NewOnSlice[interface{}](1, 2, 3, 4),
				predicate: func(e interface{}) bool { return e.(int) == 4 },
			},
			want: true,
		},
		{name: "2",
			args: args{
				source:    NewOnSlice[interface{}]("one", "two", "three", "four"),
				predicate: func(e interface{}) bool { return len(e.(string)) == 4 },
			},
			want: true,
		},
		{name: "3",
			args: args{
				source:    NewOnSlice[interface{}](1, 2, "three", "four"),
				predicate: func(e interface{}) bool { _, ok := e.(int); return ok },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := AnyPred(tt.args.source, tt.args.predicate); got != tt.want {
				t.Errorf("AnyPred() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_All_int(t *testing.T) {
	type args struct {
		source    Enumerator[int]
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
				source: NewOnSlice(1, 3, 5),
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "EmptySequenceReturnsTrue",
			args: args{
				source:    Empty[int](),
				predicate: func(x int) bool { return x > 10 },
			},
			want: true,
		},
		{name: "PredicateMatchingNoElements",
			args: args{
				source:    NewOnSlice(1, 5, 20, 30),
				predicate: func(x int) bool { return x < 0 },
			},
			want: false,
		},
		{name: "PredicateMatchingSomeElements",
			args: args{
				source:    NewOnSlice(1, 5, 8, 9),
				predicate: func(x int) bool { return x > 3 },
			},
			want: false,
		},
		{name: "PredicateMatchingAllElements",
			args: args{
				source:    NewOnSlice(1, 5, 8, 9),
				predicate: func(x int) bool { return x > 0 },
			},
			want: true,
		},
		{name: "SequenceIsNotEvaluatedAfterFirstNonMatch",
			args: args{
				source:    SelectMust(NewOnSliceEn(10, 2, 0, 3), func(x int) int { return 10 / x }),
				predicate: func(y int) bool { return y > 2 },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := All(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("All() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("All() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("All() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_All_interface(t *testing.T) {
	type args struct {
		source    Enumerator[interface{}]
		predicate func(interface{}) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1",
			args: args{
				source:    NewOnSlice[interface{}]("one", "two", "three", "four"),
				predicate: func(e interface{}) bool { return len(e.(string)) >= 3 },
			},
			want: true,
		},
		{name: "2",
			args: args{
				source:    NewOnSlice[interface{}]("one", "two", "three", "four"),
				predicate: func(e interface{}) bool { return len(e.(string)) > 3 },
			},
			want: false,
		},
		{name: "3",
			args: args{
				source:    NewOnSlice[interface{}](1, 2, "three", "four"),
				predicate: func(e interface{}) bool { _, ok := e.(int); return ok },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := All(tt.args.source, tt.args.predicate); got != tt.want {
				t.Errorf("All() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
