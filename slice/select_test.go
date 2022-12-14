package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSelect_int_int(t *testing.T) {
	var count int
	type args struct {
		source   []int
		selector func(int) int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source:   nil,
				selector: func(x int) int { return x + 1 },
			},
			want: nil,
		},
		{name: "EmptySource",
			args: args{
				source:   []int{},
				selector: func(x int) int { return x + 1 },
			},
			want: []int{},
		},
		{name: "NilSelector",
			args: args{
				source:   []int{1, 3, 7, 9, 10},
				selector: nil,
			},
			wantErr: true,
		},
		{name: "SimpleProjection",
			args: args{
				source:   []int{1, 5, 2},
				selector: func(x int) int { return x * 2 },
			},
			want: []int{2, 10, 4},
		},
		{name: "SideEffectsInProjection",
			args: args{
				source:   []int{-3, -2, -1}, // Actual values won't be relevant
				selector: func(int) int { count++; return count },
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Select(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelect_int_string(t *testing.T) {
	type args struct {
		source   []int
		selector func(int) string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "SimpleProjectionToDifferentType",
			args: args{
				source:   []int{1, 5, 2},
				selector: func(x int) string { return fmt.Sprint(x) },
			},
			want: []string{"1", "5", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Select(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelect_string_string(t *testing.T) {
	type args struct {
		source   []string
		selector func(string) string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "Select",
			args: args{
				source:   []string{"an", "apple", "a", "day"},
				selector: func(s string) string { return string([]rune(s)[0]) },
			},
			want: []string{"a", "a", "a", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Select(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectIdx_int_int(t *testing.T) {
	type args struct {
		source   []int
		selector func(int, int) int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source:   nil,
				selector: func(x, idx int) int { return x + idx },
			},
			want: nil,
		},
		{name: "EmptySource",
			args: args{
				source:   []int{},
				selector: func(x, idx int) int { return x + idx },
			},
			want: []int{},
		},
		{name: "NilSelector",
			args: args{
				source:   []int{1, 3, 7, 9, 10},
				selector: nil,
			},
			wantErr: true,
		},
		{name: "SimpleProjection",
			args: args{
				source:   []int{1, 5, 2},
				selector: func(x, idx int) int { return x + idx*10 },
			},
			want: []int{1, 15, 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectIdx(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectIdx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectIdx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectIdx_string_string(t *testing.T) {
	type args struct {
		source   []string
		selector func(string, int) string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "SimpleProjection",
			args: args{
				source: []string{"one", "two", "three", "four"},
				selector: func(s string, idx int) string {
					r := s
					for i := 0; i < idx; i++ {
						r += s
					}
					return r
				},
			},
			want: []string{"one", "twotwo", "threethreethree", "fourfourfourfour"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectIdx(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectIdx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectIdx() = %v, want %v", got, tt.want)
			}
		})
	}
}
