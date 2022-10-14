//go:build go1.18

package go2linq

import (
	"testing"
)

func TestEnrSlice_moveNext(t *testing.T) {
	type args struct {
		enr *enrSlice[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "0",
			args: args{
				enr: newEnrSlice[int](),
			},
			want: false,
		},
		{name: "1",
			args: args{
				enr: newEnrSlice(1),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := enrSlice_moveNext(tt.args.enr); got != tt.want {
				t.Errorf("enrSlice_moveNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnrSlice_moveNext_2(t *testing.T) {
	t.Run("", func(t *testing.T) {
		enr := newEnrSlice(1)
		enr.MoveNext()
		got := enrSlice_moveNext(enr)
		want := false
		if got != want {
			t.Errorf("enrSlice_moveNext() = %v, want %v", got, want)
		}
	})
}

func TestEnrSlice_current(t *testing.T) {
	type args struct {
		enr *enrSlice[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "NoElements",
			args: args{
				enr: newEnrSlice[int](),
			},
			want: 6,
		},
		{name: "EmptyElements",
			args: args{
				enr: newEnrSlice([]int{}...),
			},
			want: 23,
		},
		{name: "1",
			args: args{
				enr: newEnrSlice(1, 2),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.args.enr.MoveNext() {
				t.Skip("empty enumerator")
			}
			if got := enrSlice_current(tt.args.enr); got != tt.want {
				t.Errorf("enrSlice_current() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnrSlice_current_2(t *testing.T) {
	t.Run("", func(t *testing.T) {
		enr := newEnrSlice(1, 2)
		enr.MoveNext()
		enr.MoveNext()
		got := enrSlice_current(enr)
		want := 2
		if got != want {
			t.Errorf("enrSlice_current() = %v, want %v", got, want)
		}
	})
}
