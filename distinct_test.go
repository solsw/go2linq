//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/DistinctTest.cs

var (
	testString1 = "test"
	testString2 = "test"
)

func Test_Distinct_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceNoComparer",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Distinct(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Distinct() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Distinct() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Distinct() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}

func Test_Distinct2_string(t *testing.T) {
	type args struct {
		source Enumerable[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "1",
			args: args{
				source: NewEnSlice("A", "a", "b", "c", "b"),
			},
			want: NewEnSlice("A", "a", "b", "c"),
		},
		{name: "2",
			args: args{
				source: NewEnSlice("b", "a", "d", "a"),
			},
			want: NewEnSlice("b", "a", "d"),
		},
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#distinct-and-distinctby
		{name: "Distinct",
			args: args{
				source: NewEnSlice("Mercury", "Venus", "Venus", "Earth", "Mars", "Earth"),
			},
			want: NewEnSlice("Mercury", "Venus", "Earth", "Mars"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Distinct(tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", EnToString(got), EnToString(tt.want))
			}
		})
	}
}

func Test_DistinctEq2_string(t *testing.T) {
	type args struct {
		source  Enumerable[string]
		equaler Equaler[string]
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[string]
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceWithComparer",
			args: args{
				equaler: CaseInsensitiveEqualer,
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullComparerUsesDefault",
			args: args{
				source: NewEnSlice("xyz", testString1, "XYZ", testString2, "def"),
			},
			want: NewEnSlice("xyz", testString1, "XYZ", "def"),
		},
		{name: "NonNullEqualer",
			args: args{
				source:  NewEnSlice("xyz", testString1, "XYZ", testString2, "def"),
				equaler: CaseInsensitiveEqualer,
			},
			want: NewEnSlice("xyz", testString1, "def"),
		},
		{name: "DistinctStringsWithCaseInsensitiveComparer",
			args: args{
				source:  NewEnSlice("xyz", testString1, "XYZ", testString2, "def"),
				equaler: CaseInsensitiveEqualer,
			},
			want: NewEnSlice("xyz", testString1, "def"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctEq(tt.args.source, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistinctEq() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("DistinctEq() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DistinctEq() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}

func Test_DistinctCmp2_string(t *testing.T) {
	type args struct {
		source Enumerable[string]
		cmp    Comparer[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "DistinctStringsWithCaseInsensitiveComparer",
			args: args{
				source: NewEnSlice("xyz", testString1, "XYZ", testString2, "def"),
				cmp:    CaseInsensitiveComparer,
			},
			want: NewEnSlice("xyz", testString1, "def"),
		},
		{name: "3",
			args: args{
				source: NewEnSlice("A", "a", "b", "c", "b"),
				cmp:    CaseInsensitiveComparer,
			},
			want: NewEnSlice("A", "b", "c"),
		},
		{name: "4",
			args: args{
				source: NewEnSlice("b", "a", "d", "a"),
				cmp:    CaseInsensitiveComparer,
			},
			want: NewEnSlice("b", "a", "d"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DistinctCmp(tt.args.source, tt.args.cmp)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DistinctCmp() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}

func Test_DistinctCmp2_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
		cmp    Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "EmptyEnumerable",
			args: args{
				source: Empty[int](),
				cmp:    Order[int]{},
			},
			want: Empty[int](),
		},
		{name: "1",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
				cmp:    Order[int]{},
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
		{name: "2",
			args: args{
				source: ConcatMust(NewEnSlice(1, 2, 3, 4), NewEnSlice(1, 2, 3, 4)),
				cmp:    Order[int]{},
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DistinctCmp(tt.args.source, tt.args.cmp)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DistinctCmp() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}
