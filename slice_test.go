package go2linq

import (
	"iter"
	"testing"
)

func TestSliceAll_int(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "nil slice",
			args: args{s: nil},
			want: Empty[int](),
		},
		{name: "empty slice",
			args: args{s: []int{}},
			want: Empty[int](),
		},
		{name: "normal slice",
			args: args{s: []int{1, 2, 3, 4}},
			want: VarAll(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceAll(tt.args.s)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SliceAll() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestVarAll_int_1(t *testing.T) {
	t.Run("", func(t *testing.T) {
		next, stop := iter.Pull(VarAll(1))
		defer stop()
		_, _ = next()
		_, got := next()
		want := false
		if got != want {
			t.Errorf("VarAll_1() = %v, want %v", got, want)
		}
	})
}

func TestVarAll_int_2(t *testing.T) {
	t.Run("", func(t *testing.T) {
		next, stop := iter.Pull(VarAll(1, 2))
		defer stop()
		_, _ = next()
		got, _ := next()
		want := 2
		if got != want {
			t.Errorf("VarAll_2() = %v, want %v", got, want)
		}
	})
}
