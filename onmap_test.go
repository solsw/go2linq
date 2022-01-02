//go:build go1.18

package go2linq

import (
	"testing"
)

func Test_NewOnMap_int_string(t *testing.T) {
	m1 := make(map[int]string)
	m1[1] = "one"
	m1[2] = "two"
	// m1[3] = "three"
	type args struct {
		m map[int]string
	}
	tests := []struct {
		name string
		args args
		want Enumerator[KeyElement[int, string]]
	}{
		{name: "1",
			args: args{
				m: m1,
			},
			want: NewOnSlice(
				KeyElement[int, string]{key: 1, element: "one"},
				KeyElement[int, string]{key: 2, element: "two"},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewOnMap[int, string](tt.args.m)
			if !SequenceEqualMust[KeyElement[int, string]](got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("NewOnMap() = '%v', want '%v'", String[KeyElement[int, string]](got), String(tt.want))
			}
		})
	}
}
