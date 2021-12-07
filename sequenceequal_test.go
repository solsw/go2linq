//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SequenceEqualTest.cs

func Test_SequenceEqual_int(t *testing.T) {
	type args struct {
		first  Enumerator[int]
		second Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "0",
			args: args{
				first:  Empty[int](),
				second: Empty[int](),
			},
			want: true,
		},
		{name: "01",
			args: args{
				first:  NewOnSlice(1),
				second: Empty[int](),
			},
			want: false,
		},
		{name: "02",
			args: args{
				first:  Empty[int](),
				second: NewOnSlice(2),
			},
			want: false,
		},
		{name: "1",
			args: args{
				first:  NewOnSlice(1),
				second: NewOnSlice(1),
			},
			want: true,
		},
		{name: "UnequalLengthsBothArrays",
			args: args{
				first:  NewOnSlice(1, 5, 3),
				second: NewOnSlice(1, 5, 3, 10),
			},
			want: false,
		},
		{name: "UnequalLengthsBothRangesFirstLonger",
			args: args{
				first:  RangeMust(0, 11),
				second: RangeMust(0, 10),
			},
			want: false,
		},
		{name: "UnequalLengthsBothRangesSecondLonger",
			args: args{
				first:  RangeMust(0, 10),
				second: RangeMust(0, 11),
			},
			want: false,
		},
		{name: "UnequalData",
			args: args{
				first:  NewOnSlice(1, 5, 3, 9),
				second: NewOnSlice(1, 5, 3, 10),
			},
			want: false,
		},
		{name: "EqualDataBothArrays",
			args: args{
				first:  NewOnSlice(1, 5, 3, 10),
				second: NewOnSlice(1, 5, 3, 10),
			},
			want: true,
		},
		{name: "EqualDataBothRanges",
			args: args{
				first:  RangeMust(0, 10),
				second: RangeMust(0, 10),
			},
			want: true,
		},
		{name: "OrderMatters",
			args: args{
				first:  NewOnSlice(1, 2),
				second: NewOnSlice(2, 1),
			},
			want: false,
		},
		{name: "ReturnAtFirstDifference",
			args: args{
				first:  SelectMust(NewOnSliceEn(1, 5, 10, 2, 0), func(i int) int { return 10 / i }),
				second: SelectMust(NewOnSliceEn(1, 5, 10, 1, 0), func(i int) int { return 10 / i }),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := SequenceEqual(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("SequenceEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SequenceEqual_string(t *testing.T) {
	type args struct {
		first  Enumerator[string]
		second Enumerator[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "2",
			args: args{
				first:  NewOnSlice("one", "two", "three", "four"),
				second: NewOnSlice("one", "two", "three", "four"),
			},
			want: true,
		},
		{name: "4",
			args: args{
				first:  NewOnSlice("a", "b"),
				second: NewOnSlice("a"),
			},
			want: false,
		},
		{name: "5",
			args: args{
				first:  NewOnSlice("a"),
				second: NewOnSlice("a", "b"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := SequenceEqual(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("SequenceEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SequenceEqualEq_string(t *testing.T) {
	type args struct {
		first  Enumerator[string]
		second Enumerator[string]
		eq     Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1",
			args: args{
				first:  NewOnSlice("a", "b"),
				second: NewOnSlice("a", "B"),
				eq:     CaseInsensitiveEqualer,
			},
			want: true,
		},
		{name: "CustomEqualityComparer",
			args: args{
				first:  NewOnSlice("foo", "BAR", "baz"),
				second: NewOnSlice("FOO", "bar", "Baz"),
				eq:     CaseInsensitiveEqualer,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := SequenceEqualEq(tt.args.first, tt.args.second, tt.args.eq); got != tt.want {
				t.Errorf("SequenceEqualEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SequenceEqualSelf_int(t *testing.T) {
	r0 := RangeMust(0, 0)
	r1 := RangeMust(0, 1)
	r2 := RangeMust(0, 2)
	r3 := RepeatMust(1, 4)
	type args struct {
		first  Enumerator[int]
		second Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Same0",
			args: args{
				first:  r0,
				second: r0,
			},
			want: true,
		},
		{name: "Same1",
			args: args{
				first:  r1,
				second: r1,
			},
			want: true,
		},
		{name: "Same2",
			args: args{
				first:  r2,
				second: r2,
			},
			want: true,
		},
		{name: "Same3",
			args: args{
				first:  TakeMust(r3, 2),
				second: SkipMust(r3, 2),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := SequenceEqualSelf(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("SequenceEqualSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
