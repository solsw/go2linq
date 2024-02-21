package go2linq

import (
	"iter"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/EmptyTest.cs

func TestEmpty_int(t *testing.T) {
	tests := []struct {
		name string
		want iter.Seq[int]
	}{
		{name: "EmptyContainsNoElements",
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Empty[int]()
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Empty() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestEmpty_string(t *testing.T) {
	tests := []struct {
		name string
		want iter.Seq[string]
	}{
		{name: "EmptyContainsNoElements",
			want: Empty[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Empty[string]()
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Empty() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}
