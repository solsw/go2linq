//go:build go1.18

package go2linq

import (
	"testing"
)

func Test_NewOnChan_int(t *testing.T) {
	in1 := make(chan int)
	go func() {
		for i := 1; i <= 4; i++ {
			in1 <- i
		}
		close(in1)
	}()
	type args struct {
		chn <-chan int
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "1",
			args: args{
				chn: in1,
			},
			want: NewOnSlice(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOnChan[int](tt.args.chn); !SequenceEqualMust(got, tt.want) {
				t.Errorf("NewOnChan() failed")
			}
		})
	}
}
