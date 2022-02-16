//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/EmptyTest.cs

func Test_Empty_int(t *testing.T) {
	tests := []struct {
		name string
		want Enumerable[int]
	}{
		{name: "EmptyContainsNoElements",
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Empty[int]()
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Empty() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_Empty_string(t *testing.T) {
	tests := []struct {
		name string
		want Enumerable[string]
	}{
		{name: "EmptyContainsNoElements",
			want: Empty[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Empty[string]()
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Empty() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}
