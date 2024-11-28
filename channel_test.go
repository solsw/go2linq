package go2linq

import (
	"fmt"
	"iter"
	"testing"

	"github.com/solsw/iterhelper"
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

func chn3() chan int {
	ch := make(chan int)
	go func() {
		ch <- 4
		ch <- 3
		ch <- 2
		ch <- 1
		close(ch)
	}()
	return ch
}

func TestChanAll_int(t *testing.T) {
	type args struct {
		c <-chan int
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "nil channel",
			args: args{c: nil},
			want: Empty[int](),
		},
		{name: "closed channel",
			args: args{c: closedCh()},
			want: Empty[int](),
		},
		{name: "2",
			args: args{c: chn2()},
			want: iterhelper.VarSeq(1),
		},
		{name: "3",
			args: args{c: chn3()},
			want: iterhelper.VarSeq(4, 3, 2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChanAll(tt.args.c)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("ChanAll() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

func TestChanAll_int_2(t *testing.T) {
	t.Run("", func(t *testing.T) {
		next, stop := iter.Pull(ChanAll(chn2()))
		defer stop()
		_, _ = next()
		got, _ := next()
		want := 0
		if got != want {
			t.Errorf("ChanAll() = %v, want %v", got, want)
		}
	})
}

func ExampleChanAll() {
	seq1 := ChanAll(chn3())
	seq2, _ := Select(seq1, func(i int) int { return 12 / i })
	first1, _ := First(seq2)
	fmt.Println(first1)
	skip, _ := Skip(seq2, 2)
	first2, _ := First(skip)
	fmt.Println(first2)
	// Output:
	// 3
	// 12
}
