//go:build go1.18

package go2linq

import (
	"testing"
)

func Test_enrSlice_moveNext(t *testing.T) {
	type args struct {
		enr *enrSlice[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "0",
			args: args{enr: newEnrSlice[int]()},
			want: false,
		},
		{name: "1",
			args: args{enr: newEnrSlice(1)},
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

func Test_enrSlice_moveNext_2(t *testing.T) {
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

func Test_enrSlice_current_0(t *testing.T) {
	type args struct {
		enr *enrSlice[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0",
			args: args{enr: newEnrSlice[int]()},
			want: 0,
		},
		{name: "1",
			args: args{enr: newEnrSlice(1)},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := enrSlice_current(tt.args.enr); got != tt.want {
				t.Errorf("enrSlice_current() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_enrSlice_current_2_0(t *testing.T) {
	t.Run("", func(t *testing.T) {
		enr := newEnrSlice(1, 2)
		got := enrSlice_current(enr)
		want := 0
		if got != want {
			t.Errorf("enrSlice_current() = %v, want %v", got, want)
		}
	})
}

func Test_enrSlice_current_2_1(t *testing.T) {
	t.Run("", func(t *testing.T) {
		enr := newEnrSlice(1, 2)
		enr.MoveNext()
		got := enrSlice_current(enr)
		want := 1
		if got != want {
			t.Errorf("enrSlice_current() = %v, want %v", got, want)
		}
	})
}

func Test_enrSlice_current_2_2(t *testing.T) {
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

func Test_enrSlice_current_2_3(t *testing.T) {
	t.Run("", func(t *testing.T) {
		enr := newEnrSlice(1, 2)
		enr.MoveNext()
		enr.MoveNext()
		enr.MoveNext()
		got := enrSlice_current(enr)
		want := 2
		if got != want {
			t.Errorf("enrSlice_current() = %v, want %v", got, want)
		}
	})
}
