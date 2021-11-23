//go:build go1.18

package go2linq

import (
	"reflect"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/FirstTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/FirstOrDefaultTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SingleTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SingleOrDefaultTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/LastTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/LastOrDefaultTest.cs

func Test_FirstErr_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
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
				source: NewOnSlice(5),
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithoutPredicate",
			args: args{
				source: NewOnSlice(5, 10),
			},
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithoutPredicate",
			args: args{
				source: Select(
					NewOnSlice(15, 1, 0, 3),
					func(x int) int { return 10 / x },
				),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstErr(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("FirstErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstErr() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_FirstPredErr_int(t *testing.T) {
	type args struct {
		source    Enumerator[int]
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
				source: NewOnSlice(1, 2, 3, 4),
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
				source:    NewOnSlice(5),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "SingleElementSequenceWithNonMatchingPredicate",
			args: args{
				source:    NewOnSlice(2),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNoMatch,
		},
		{name: "MultipleElementSequenceWithNoPredicateMatches",
			args: args{
				source:    NewOnSlice(1, 2, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNoMatch,
		},
		{name: "MultipleElementSequenceWithSinglePredicateMatch",
			args: args{
				source:    NewOnSlice(1, 2, 5, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			args: args{
				source:    NewOnSlice(1, 2, 5, 10, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithPredicate",
			args: args{
				source: Select(
					NewOnSlice(15, 1, 0, 3),
					func(x int) int { return 10 / x },
				),
				predicate: func(y int) bool { return y > 5 },
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstPredErr(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstPredErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("FirstPredErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstPredErr() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_FirstOrDefault_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
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
				source: NewOnSlice(5),
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithoutPredicate",
			args: args{
				source: NewOnSlice(5, 10),
			},
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithoutPredicate",
			args: args{
				source: Select(
					NewOnSlice(15, 1, 0, 3),
					func(x int) int { return 10 / x },
				),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstOrDefault(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_FirstOrDefaultPredErr_int(t *testing.T) {
	type args struct {
		source    Enumerator[int]
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
				source: NewOnSlice(1, 2, 3, 4),
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
				source:    NewOnSlice(5),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "SingleElementSequenceWithNonMatchingPredicate",
			args: args{
				source:    NewOnSlice(2),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "MultipleElementSequenceWithNoPredicateMatches",
			args: args{
				source:    NewOnSlice(1, 2, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "MultipleElementSequenceWithSinglePredicateMatch",
			args: args{
				source:    NewOnSlice(1, 2, 5, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			args: args{
				source:    NewOnSlice(1, 2, 5, 10, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithPredicate",
			args: args{
				source: Select(
					NewOnSlice(15, 1, 0, 3),
					func(x int) int { return 10 / x },
				),
				predicate: func(y int) bool { return y > 5 },
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstOrDefaultPredErr(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstOrDefaultPredErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("FirstOrDefaultPredErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOrDefaultPredErr() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_SingleErr_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
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
				source: NewOnSlice(5),
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithoutPredicate",
			args: args{
				source: NewOnSlice(5, 10),
			},
			wantErr:     true,
			expectedErr: ErrMultipleElements,
		},
		{name: "EarlyOutWithoutPredicate",
			args: args{
				source: Select(
					NewOnSlice(1, 2, 0),
					func(x int) int { return 10 / x },
				),
			},
			wantErr:     true,
			expectedErr: ErrMultipleElements,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SingleErr(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("SingleErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("SingleErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleErr() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_SinglePredErr_int(t *testing.T) {
	type args struct {
		source    Enumerator[int]
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
				source: NewOnSlice(1, 2, 3, 4),
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
				source:    NewOnSlice(5),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "SingleElementSequenceWithNonMatchingPredicate",
			args: args{
				source:    NewOnSlice(2),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNoMatch,
		},
		{name: "MultipleElementSequenceWithNoPredicateMatches",
			args: args{
				source:    NewOnSlice(1, 2, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNoMatch,
		},
		{name: "MultipleElementSequenceWithSinglePredicateMatch",
			args: args{
				source:    NewOnSlice(1, 3, 5, 4, 2),
				predicate: func(x int) bool { return x > 4 },
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			args: args{
				source:    NewOnSlice(1, 2, 5, 10, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrMultipleMatch,
		},
		{name: "EarlyOutWithPredicate",
			args: args{
				source: Select(
					NewOnSlice(1, 2, 0),
					func(x int) int { return 10 / x },
				),
				predicate: func(int) bool { return true },
			},
			wantErr:     true,
			expectedErr: ErrMultipleMatch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SinglePredErr(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("SinglePredErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("SinglePredErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SinglePredErr() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_SingleOrDefaultErr_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
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
			want: 0,
		},
		{name: "SingleElementSequenceWithoutPredicate",
			args: args{
				source: NewOnSlice(5),
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithoutPredicate",
			args: args{
				source: NewOnSlice(5, 10),
			},
			wantErr:     true,
			expectedErr: ErrMultipleElements,
		},
		{name: "EarlyOutWithoutPredicate",
			args: args{
				source: Select(
					NewOnSlice(1, 2, 0),
					func(x int) int { return 10 / x },
				),
			},
			wantErr:     true,
			expectedErr: ErrMultipleElements,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SingleOrDefaultErr(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("SingleOrDefaultErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("SingleOrDefaultErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleOrDefaultErr() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_SingleOrDefaultPredErr_int(t *testing.T) {
	type args struct {
		source    Enumerator[int]
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
				source: NewOnSlice(1, 2, 3, 4),
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
				source:    NewOnSlice(5),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "SingleElementSequenceWithNonMatchingPredicate",
			args: args{
				source:    NewOnSlice(2),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "MultipleElementSequenceWithNoPredicateMatches",
			args: args{
				source:    NewOnSlice(1, 2, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "MultipleElementSequenceWithSinglePredicateMatch",
			args: args{
				source:    NewOnSlice(1, 2, 5, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			args: args{
				source:    NewOnSlice(1, 2, 5, 10, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrMultipleMatch,
		},
		{name: "EarlyOutWithPredicate",
			args: args{
				source: Select(
					NewOnSlice(1, 2, 0),
					func(x int) int { return 10 / x },
				),
				predicate: func(int) bool { return true },
			},
			wantErr:     true,
			expectedErr: ErrMultipleMatch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SingleOrDefaultPredErr(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("SingleOrDefaultPredErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("SingleOrDefaultPredErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleOrDefaultPredErr() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_LastErr_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
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
				source: NewOnSlice(5),
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithoutPredicate",
			args: args{
				source: NewOnSlice(5, 10),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LastErr(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("LastErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastErr() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_LastPredErr_int(t *testing.T) {
	type args struct {
		source    Enumerator[int]
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
				source: NewOnSlice(1, 2, 3, 4),
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
				source:    NewOnSlice(5),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "SingleElementSequenceWithNonMatchingPredicate",
			args: args{
				source:    NewOnSlice(2),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNoMatch,
		},
		{name: "MultipleElementSequenceWithNoPredicateMatches",
			args: args{
				source:    NewOnSlice(1, 2, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNoMatch,
		},
		{name: "MultipleElementSequenceWithSinglePredicateMatch",
			args: args{
				source:    NewOnSlice(1, 2, 5, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			args: args{
				source:    NewOnSlice(1, 2, 5, 10, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LastPredErr(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastPredErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("LastPredErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastPredErr() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_LastOrDefault_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
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
				source: NewOnSlice(5),
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithoutPredicate",
			args: args{
				source: NewOnSlice(5, 10),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastOrDefault(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastOrDefault() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_LastOrDefaultPredErr_int(t *testing.T) {
	type args struct {
		source    Enumerator[int]
		predicate func(int) bool
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NullPredicate",
			args: args{
				source: NewOnSlice(1, 2, 3, 4),
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
				source:    NewOnSlice(5),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "SingleElementSequenceWithNonMatchingPredicate",
			args: args{
				source:    NewOnSlice(2),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "MultipleElementSequenceWithNoPredicateMatches",
			args: args{
				source:    NewOnSlice(1, 2, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "MultipleElementSequenceWithSinglePredicateMatch",
			args: args{
				source:    NewOnSlice(1, 2, 5, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			args: args{
				source:    NewOnSlice(1, 2, 5, 10, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LastOrDefaultPredErr(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastOrDefaultPredErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("LastOrDefaultPredErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastOrDefaultPredErr() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
