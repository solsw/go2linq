//go:build go1.18

package go2linq

import (
	"reflect"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/FirstTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/FirstOrDefaultTest.cs

func Test_First_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceWithoutPredicate",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "EmptySequenceWithoutPredicate",
			args: args{
				source: Empty[int](),
			},
			wantErr:     true,
			expectedErr: ErrEmptySource,
		},
		{name: "SingleElementSequenceWithoutPredicate",
			args: args{
				source: NewEnSlice(5),
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithoutPredicate",
			args: args{
				source: NewEnSlice(5, 10),
			},
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithoutPredicate",
			args: args{
				source: SelectMust(NewEnSlice(15, 1, 0, 3), func(x int) int { return 10 / x }),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := First(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("First() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("First() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_FirstPred_int(t *testing.T) {
	type args struct {
		source    Enumerable[int]
		predicate func(int) bool
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceWithPredicate",
			args: args{
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullPredicate",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "EmptySequenceWithPredicate",
			args: args{
				source:    Empty[int](),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrEmptySource,
		},
		{name: "SingleElementSequenceWithMatchingPredicate",
			args: args{
				source:    NewEnSlice(5),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "SingleElementSequenceWithNonMatchingPredicate",
			args: args{
				source:    NewEnSlice(2),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNoMatch,
		},
		{name: "MultipleElementSequenceWithNoPredicateMatches",
			args: args{
				source:    NewEnSlice(1, 2, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNoMatch,
		},
		{name: "MultipleElementSequenceWithSinglePredicateMatch",
			args: args{
				source:    NewEnSlice(1, 2, 5, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			args: args{
				source:    NewEnSlice(1, 2, 5, 10, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithPredicate",
			args: args{
				source:    SelectMust(NewEnSlice(15, 1, 0, 3), func(x int) int { return 10 / x }),
				predicate: func(y int) bool { return y > 5 },
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstPred(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstPred() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("FirstPred() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstPred() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_FirstOrDefaultMust_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "EmptySequenceWithoutPredicate",
			args: args{
				source: Empty[int](),
			},
			want: 0,
		},
		{name: "SingleElementSequenceWithoutPredicate",
			args: args{
				source: NewEnSlice(5),
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithoutPredicate",
			args: args{
				source: NewEnSlice(5, 10),
			},
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithoutPredicate",
			args: args{
				source: SelectMust(NewEnSlice(15, 1, 0, 3), func(x int) int { return 10 / x }),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstOrDefaultMust(tt.args.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOrDefaultMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_FirstOrDefaultPred_int(t *testing.T) {
	type args struct {
		source    Enumerable[int]
		predicate func(int) bool
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceWithoutPredicate",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullSourceWithPredicate",
			args:        args{predicate: func(x int) bool { return x > 3 }},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullPredicate",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "EmptySequenceWithPredicate",
			args: args{
				source:    Empty[int](),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "SingleElementSequenceWithMatchingPredicate",
			args: args{
				source:    NewEnSlice(5),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "SingleElementSequenceWithNonMatchingPredicate",
			args: args{
				source:    NewEnSlice(2),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "MultipleElementSequenceWithNoPredicateMatches",
			args: args{
				source:    NewEnSlice(1, 2, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "MultipleElementSequenceWithSinglePredicateMatch",
			args: args{
				source:    NewEnSlice(1, 2, 5, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			args: args{
				source:    NewEnSlice(1, 2, 5, 10, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithPredicate",
			args: args{
				source:    SelectMust(NewEnSlice(15, 1, 0, 3), func(x int) int { return 10 / x }),
				predicate: func(y int) bool { return y > 5 },
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstOrDefaultPred(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstOrDefaultPred() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("FirstOrDefaultPred() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOrDefaultPred() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
