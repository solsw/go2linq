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

func Test_DistinctErr_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
		wantErr bool
		expectedErr error
	}{
		{name: "NullSourceNoComparer",
			wantErr: true,
			expectedErr: ErrNilSource,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctErr(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistinctErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("DistinctErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("DistinctErr() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Distinct_string(t *testing.T) {
	type args struct {
		source Enumerator[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "1",
			args: args{
				source: NewOnSlice("A", "a", "b", "c", "b"),
			},
			want: NewOnSlice("A", "a", "b", "c"),
		},
		{name: "2",
			args: args{
				source: NewOnSlice("b", "a", "d", "a"),
			},
			want: NewOnSlice("b", "a", "d"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distinct(tt.args.source); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Distinct() = %v, want %v", String(got), String(tt.want))
			}
		})
	}
}

func Test_DistinctEqErr_string(t *testing.T) {
	type args struct {
		source Enumerator[string]
		eq Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
		wantErr bool
		expectedErr error
	}{
		{name: "NullSourceWithComparer",
			args: args{
				eq: CaseInsensitiveEqualer,
			},
			wantErr: true,
			expectedErr: ErrNilSource,
		},
		{name: "NullComparerUsesDefault",
			args: args{
				source: NewOnSlice("xyz", testString1, "XYZ", testString2, "def"),
			},
			want: NewOnSlice("xyz", testString1, "XYZ", "def"),
		},
		{name: "NonNullEqualer",
			args: args{
				source: NewOnSlice("xyz", testString1, "XYZ", testString2, "def"),
				eq: CaseInsensitiveEqualer,
			},
			want: NewOnSlice("xyz", testString1, "def"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctEqErr(tt.args.source, tt.args.eq)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistinctEqErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("DistinctEqErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("DistinctEqErr() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_DistinctEq_string(t *testing.T) {
	type args struct {
		source Enumerator[string]
		eq Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "DistinctStringsWithCaseInsensitiveComparer",
			args: args{
				source: NewOnSlice("xyz", testString1, "XYZ", testString2, "def"),
				eq: CaseInsensitiveEqualer,
			},
			want: NewOnSlice("xyz", testString1, "def"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctEq(tt.args.source, tt.args.eq); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("DistinctEq() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_DistinctCmp_string(t *testing.T) {
	type args struct {
		source Enumerator[string]
		cmp Comparer[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "DistinctStringsWithCaseInsensitiveComparer",
			args: args{
				source: NewOnSlice("xyz", testString1, "XYZ", testString2, "def"),
				cmp: CaseInsensitiveComparer,
			},
			want: NewOnSlice("xyz", testString1, "def"),
		},
		{name: "3",
			args: args{
				source: NewOnSlice("A", "a", "b", "c", "b"),
				cmp: CaseInsensitiveComparer,
			},
			want: NewOnSlice("A", "b", "c"),
		},
		{name: "4",
			args: args{
				source: NewOnSlice("b", "a", "d", "a"),
				cmp: CaseInsensitiveComparer,
			},
			want: NewOnSlice("b", "a", "d"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctCmp(tt.args.source, tt.args.cmp); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("DistinctCmp() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_DistinctCmp_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
		cmp Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "EmptyEnumerable",
			args: args{
				source: Empty[int](),
				cmp: IntComparer,
			},
			want: Empty[int](),
		},
		{name: "1",
			args: args{
				source: NewOnSlice(1, 2, 3, 4),
				cmp: IntComparer,
			},
			want: NewOnSlice(1, 2, 3, 4),
		},
		{name: "2",
			args: args{
				source: Concat(NewOnSlice(1, 2, 3, 4), NewOnSlice(1, 2, 3, 4)),
				cmp: IntComparer,
			},
			want: NewOnSlice(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctCmp(tt.args.source, tt.args.cmp); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("DistinctCmp() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_DistinctEq_Reset(t *testing.T) {
	source := DistinctEq(
		NewOnSlice("xyz", testString1, "XYZ", testString2, "def"),
		CaseInsensitiveEqualer)
	got1 := NewOnSlice(Slice(source)...)
	source.Reset()
	got2 := NewOnSlice(Slice(source)...)
	if !SequenceEqual(got1, got2) {
		got1.Reset()
		got2.Reset()
		t.Errorf("Reset error: '%v' != '%v'", String(got1), String(got2))
	}
}
