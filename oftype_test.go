//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OfTypeTest.cs

func Test_OfTypeMust_any_int(t *testing.T) {
	type args struct {
		source Enumerable[any]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "UnboxToInt",
			args: args{
				source: NewEnSlice[any](10, 30, 50),
			},
			want: NewEnSlice[int](10, 30, 50),
		},
		{name: "OfType",
			args: args{
				source: NewEnSlice[any](1, 2, "two", 3, 3.14, 4, nil),
			},
			want: NewEnSlice[int](1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OfTypeMust[any, int](tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("OfTypeMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_OfTypeMust_any_string(t *testing.T) {
	type args struct {
		source Enumerable[any]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "SequenceWithAllValidValues",
			args: args{
				source: NewEnSlice[any]("first", "second", "third"),
			},
			want: NewEnSlice[string]("first", "second", "third"),
		},
		{name: "NullsAreExcluded",
			args: args{
				source: NewEnSlice[any]("first", nil, "third"),
			},
			want: NewEnSlice[string]("first", "third"),
		},
		{name: "WrongElementTypesAreIgnored",
			args: args{
				source: NewEnSlice[any]("first", any(1), "third"),
			},
			want: NewEnSlice[string]("first", "third"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OfTypeMust[any, string](tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("OfTypeMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_OfTypeMust_any_int64(t *testing.T) {
	type args struct {
		source Enumerable[any]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int64]
	}{
		{name: "UnboxingWithWrongElementTypes",
			args: args{
				source: NewEnSlice[any](int64(100), 100, int64(300)),
			},
			want: NewEnSlice[int64](int64(100), int64(300)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OfTypeMust[any, int64](tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("OfTypeMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}
