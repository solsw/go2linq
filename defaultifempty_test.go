//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/DefaultIfEmptyTest.cs

func Test_DefaultIfEmpty_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "EmptySequenceNoDefaultValue",
			args: args{
				source: Empty[int](),
			},
			want: NewOnSlice(0),
		},
		{name: "NonEmptySequenceNoDefaultValue",
			args: args{
				source: NewOnSlice(3, 1, 4),
			},
			want: NewOnSlice(3, 1, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := DefaultIfEmpty(tt.args.source); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("DefaultIfEmpty() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_DefaultIfEmptyDef_int(t *testing.T) {
	type args struct {
		source       Enumerator[int]
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "EmptySequenceWithDefaultValue",
			args: args{
				source:       Empty[int](),
				defaultValue: 5,
			},
			want: NewOnSlice(5),
		},
		{name: "NonEmptySequenceWithDefaultValue",
			args: args{
				source:       NewOnSlice(3, 1, 4),
				defaultValue: 5,
			},
			want: NewOnSlice(3, 1, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := DefaultIfEmptyDef(tt.args.source, tt.args.defaultValue); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("DefaultIfEmptyDef() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
