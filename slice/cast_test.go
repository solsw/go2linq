package slice

import (
	"reflect"
	"testing"
)

func TestCast_any_int(t *testing.T) {
	type args struct {
		source []any
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source: nil,
			},
			want: nil,
		},
		{name: "EmptySource",
			args: args{
				source: []any{},
			},
			want: []int{},
		},
		{name: "UnboxToInt",
			args: args{
				source: []any{10, 30, 50},
			},
			want: []int{10, 30, 50},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Cast[any, int](tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCastMust_any_string(t *testing.T) {
	type args struct {
		source []any
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "SequenceWithAllValidValues",
			args: args{
				source: []any{"first", "second", "third"},
			},
			want: []string{"first", "second", "third"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CastMust[any, string](tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CastMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
