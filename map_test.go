package go2linq

import (
	"fmt"
	"iter"
	"testing"
)

func TestMapAll_int_string(t *testing.T) {
	type args struct {
		m map[int]string
	}
	tests := []struct {
		name string
		args args
		want iter.Seq2[int, string]
	}{
		{name: "nil map",
			args: args{m: nil},
			want: Empty2[int, string](),
		},
		{name: "zero map",
			args: args{m: map[int]string{}},
			want: Empty2[int, string](),
		},
		{name: "empty map",
			args: args{m: make(map[int]string)},
			want: Empty2[int, string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapAll(tt.args.m)
			equal, _ := SequenceEqual2(got, tt.want)
			if !equal {
				t.Errorf("MapAll() = %v, want %v", StringDef2(got), StringDef2(tt.want))
			}
		})
	}
}

func ExampleMapAll() {
	m := map[int]string{1: "one", 2: "two", 3: "three", 4: "four"}
	for i, s := range MapAll(m) {
		fmt.Println(i, s)
	}
	// Unordered output:
	// 1 one
	// 2 two
	// 3 three
	// 4 four
}
