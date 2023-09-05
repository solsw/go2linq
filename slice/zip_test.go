package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestZip_string_int_string(t *testing.T) {
	r1, _ := Range(5, 10)
	r23, _ := Range(5, 3)
	type args struct {
		first          []string
		second         []int
		resultSelector func(string, int) string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "ShortFirstSlice",
			args: args{
				first:          []string{"a", "b", "c"},
				second:         r1,
				resultSelector: func(s string, i int) string { return fmt.Sprintf("%s:%d", s, i) },
			},
			want: []string{"a:5", "b:6", "c:7"},
		},
		{name: "ShortSecondSlice",
			args: args{
				first:          []string{"a", "b", "c", "d", "e"},
				second:         r23,
				resultSelector: func(s string, i int) string { return fmt.Sprintf("%s:%d", s, i) },
			},
			want: []string{"a:5", "b:6", "c:7"},
		},
		{name: "EqualLengthSlices",
			args: args{
				first:          []string{"a", "b", "c"},
				second:         r23,
				resultSelector: func(s string, i int) string { return fmt.Sprintf("%s:%d", s, i) },
			},
			want: []string{"a:5", "b:6", "c:7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Zip(tt.args.first, tt.args.second, tt.args.resultSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZip_string_string_string(t *testing.T) {
	ss1 := []string{"a", "b", "c"}
	ss2 := []string{"a", "b", "c", "d", "e"}
	type args struct {
		first          []string
		second         []string
		resultSelector func(string, string) string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "1",
			args: args{
				first:          []string{"one", "two", "three", "four"},
				second:         []string{"four", "three", "two", "one"},
				resultSelector: func(s1, s2 string) string { return s1 + s2 },
			},
			want: []string{"onefour", "twothree", "threetwo", "fourone"},
		},
		{name: "SameEnumerableString1",
			args: args{
				first:          ss1,
				second:         ss1,
				resultSelector: func(s1, s2 string) string { return s1 + s2 },
			},
			want: []string{"aa", "bb", "cc"},
		},
		{name: "AdjacentElements",
			args: args{
				first:          ss2,
				second:         ss2[1:],
				resultSelector: func(s1, s2 string) string { return s1 + s2 },
			},
			want: []string{"ab", "bc", "cd", "de"},
		},
		{name: "AdjacentElements2",
			args: args{
				first:          ss2[1:],
				second:         ss2,
				resultSelector: func(s1, s2 string) string { return s1 + s2 },
			},
			want: []string{"ba", "cb", "dc", "ed"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Zip(tt.args.first, tt.args.second, tt.args.resultSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZip_int_int_string(t *testing.T) {
	en0, _ := Range(1, 4)
	en1 := en0[:2]
	en2 := en0[2:]
	type args struct {
		first          []int
		second         []int
		resultSelector func(int, int) string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "SameEnumerableInt00",
			args: args{
				first:          en0,
				second:         en0,
				resultSelector: func(i1, i2 int) string { return fmt.Sprintf("%d%d", i1, i2) },
			},
			want: []string{"11", "22", "33", "44"},
		},
		{name: "SameEnumerableInt01",
			args: args{
				first:          en0[2:],
				second:         en0,
				resultSelector: func(i1, i2 int) string { return fmt.Sprintf("%d%d", i1, i2) },
			},
			want: []string{"31", "42"},
		},
		{name: "SameEnumerableInt1",
			args: args{
				first:          en1,
				second:         en1,
				resultSelector: func(i1, i2 int) string { return fmt.Sprintf("%d%d", i1, i2) },
			},
			want: []string{"11", "22"},
		},
		{name: "SameEnumerableInt2",
			args: args{
				first:          en2,
				second:         en2,
				resultSelector: func(i1, i2 int) string { return fmt.Sprintf("%d%d", i1, i2) },
			},
			want: []string{"33", "44"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Zip(tt.args.first, tt.args.second, tt.args.resultSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZip_string_string_int(t *testing.T) {
	type args struct {
		first          []string
		second         []string
		resultSelector func(string, string) int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "1",
			args: args{
				first:          []string{"a", "b", "c"},
				second:         []string{"one", "two", "three", "four"},
				resultSelector: func(s1, s2 string) int { return len(s1 + s2) },
			},
			want: []int{4, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Zip(tt.args.first, tt.args.second, tt.args.resultSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZip_int_rune_string(t *testing.T) {
	type args struct {
		first          []int
		second         []rune
		resultSelector func(int, rune) string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/projection-operations#zip
		{name: "Zip",
			args: args{
				first:          []int{1, 2, 3, 4, 5, 6, 7},
				second:         []rune{'A', 'B', 'C', 'D', 'E', 'F'},
				resultSelector: func(number int, letter rune) string { return fmt.Sprintf("%d = %c (%[2]d)", number, letter) },
			},
			want: []string{"1 = A (65)", "2 = B (66)", "3 = C (67)", "4 = D (68)", "5 = E (69)", "6 = F (70)"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Zip(tt.args.first, tt.args.second, tt.args.resultSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}
