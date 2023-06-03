package go2linq

import (
	"testing"
)

func Test_enSliceCount_int(t *testing.T) {
	type args struct {
		en *EnSlice[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1",
			args: args{en: NewEnSlice[int](1, 2, 3, 4)},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.en.Count(); got != tt.want {
				t.Errorf("enSliceCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
