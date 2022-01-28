//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/DefaultIfEmptyTest.cs

func Test_DefaultIfEmptyMust_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "EmptySequenceNoDefaultValue",
			args: args{
				source: Empty[int](),
			},
			want: NewEnSlice(0),
		},
		{name: "NonEmptySequenceNoDefaultValue",
			args: args{
				source: NewEnSlice(3, 1, 4),
			},
			want: NewEnSlice(3, 1, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DefaultIfEmptyMust(tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DefaultIfEmptyMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_DefaultIfEmptyDefMust_int(t *testing.T) {
	type args struct {
		source       Enumerable[int]
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "EmptySequenceWithDefaultValue",
			args: args{
				source:       Empty[int](),
				defaultValue: 5,
			},
			want: NewEnSlice(5),
		},
		{name: "NonEmptySequenceWithDefaultValue",
			args: args{
				source:       NewEnSlice(3, 1, 4),
				defaultValue: 5,
			},
			want: NewEnSlice(3, 1, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DefaultIfEmptyDefMust(tt.args.source, tt.args.defaultValue)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DefaultIfEmptyDefMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}
