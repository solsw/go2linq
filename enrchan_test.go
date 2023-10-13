package go2linq

import (
	"errors"
	"testing"

	"github.com/solsw/errorhelper"
)

func closedCh() chan int {
	var ch = make(chan int)
	close(ch)
	return ch
}

func chn2() chan int {
	var ch = make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()
	return ch
}

func TestEnrChan_moveNext(t *testing.T) {
	type args struct {
		enr *enrChan[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "0",
			args: args{enr: &enrChan[int]{}},
			want: false,
		},
		{name: "1",
			args: args{enr: &enrChan[int]{chn: closedCh()}},
			want: false,
		},
		{name: "2",
			args: args{enr: &enrChan[int]{chn: chn2()}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := enrChan_moveNext(tt.args.enr); got != tt.want {
				t.Errorf("enrChan_moveNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnrChan_current(t *testing.T) {
	type args struct {
		enr *enrChan[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {name: "nil",
		// 	args: args{enr: nil},
		// 	want: 0,
		// },
		{name: "0",
			args: args{enr: &enrChan[int]{}},
			want: 0,
		},
		{name: "1",
			args: args{enr: &enrChan[int]{chn: closedCh()}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := enrChan_current(tt.args.enr); got != tt.want {
				t.Errorf("enrChan_current() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnrChan_current_2_1(t *testing.T) {
	t.Run("", func(t *testing.T) {
		enr := &enrChan[int]{chn: chn2()}
		enr.MoveNext()
		got := enrChan_current(enr)
		want := 1
		if got != want {
			t.Errorf("enrChan_current() = %v, want %v", got, want)
		}
	})
}

func TestEnrChan_current_2_2(t *testing.T) {
	t.Run("", func(t *testing.T) {
		enr := &enrChan[int]{chn: chn2()}
		enr.MoveNext()
		enr.MoveNext()
		got := enrChan_current(enr)
		want := 0
		if got != want {
			t.Errorf("enrChan_current() = %v, want %v", got, want)
		}
	})
}

func TestEnrChan_Reset(t *testing.T) {
	t.Run("", func(t *testing.T) {
		enr := &enrChan[int]{chn: chn2()}
		gotErrPanic := func() (errPanic error) {
			defer func() {
				errorhelper.PanicToError(recover(), &errPanic)
			}()
			enr.Reset()
			return nil
		}()
		want := ErrResetNotSupported
		if !errors.Is(gotErrPanic, want) {
			t.Errorf("EnrChan_Reset() panic = '%v', want '%v'", gotErrPanic, want)
		}
	})
}
