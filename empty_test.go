//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/EmptyTest.cs

func Test_Empty_int(t *testing.T) {
	tests := []struct {
		name string
		want Enumerator[int]
	}{
		{name: "EmptyContainsNoElements",
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Empty[int](); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Empty() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Empty_string(t *testing.T) {
	tests := []struct {
		name string
		want Enumerator[string]
	}{
		{name: "EmptyContainsNoElements",
			want: Empty[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Empty[string](); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Empty() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
