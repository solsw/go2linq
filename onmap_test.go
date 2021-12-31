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
		want Enumerator[KeyValue[int, string]]
	}{
		{name: "1",
			args: args{
				m: m1,
			},
			want: NewOnSlice(
				KeyValue[int, string]{key: 1, value: "one"},
				KeyValue[int, string]{key: 2, value: "two"},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewOnMap[int, string](tt.args.m)
			if !SequenceEqualMust[KeyValue[int, string]](got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("NewOnMap() = '%v', want '%v'", String[KeyValue[int, string]](got), String(tt.want))
			}
		})
	}
}
