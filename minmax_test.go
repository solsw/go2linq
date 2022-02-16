//go:build go1.18

package go2linq

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/MinTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/MaxTest.cs

func Test_MinSel_string_int(t *testing.T) {
	type args struct {
		source   Enumerable[string]
		selector func(string) int
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NilSelector",
			args: args{
				source: Empty[string](),
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "EmptySequenceWithSelector",
			args: args{
				source:   Empty[string](),
				selector: func(s string) int { return len(s) },
			},
			wantErr:     true,
			expectedErr: ErrEmptySource,
		},
		{name: "SimpleSequenceWithSelector",
			args: args{
				source:   NewEnSlice("xyz", "ab", "abcde", "0"),
				selector: func(s string) int { return len(s) },
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinSel(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinSel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("MinSel() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinSelLs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MinSelLs_string_rune(t *testing.T) {
	type args struct {
		source   Enumerable[string]
		selector func(string) rune
		lesser   Lesser[rune]
	}
	tests := []struct {
		name        string
		args        args
		want        rune
		wantErr     bool
		expectedErr error
	}{
		{name: "SimpleSequenceWithSelector2",
			args: args{
				source:   NewEnSlice("xyz", "ab", "abcde", "0"),
				selector: func(s string) rune { return []rune(s)[0] },
				lesser:   LesserFunc[rune](func(r1, r2 rune) bool { return r1 < r2 }),
			},
			want: '0',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinSelLs(tt.args.source, tt.args.selector, tt.args.lesser)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinSelLs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("MinSelLs() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinSelLs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MinBySelMust_string_int(t *testing.T) {
	type args struct {
		source   Enumerable[string]
		selector func(string) int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "MinElement",
			args: args{
				source:   NewEnSlice("xyz", "ab", "abcde", "0"),
				selector: func(s string) int { return len(s) },
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinBySelMust(tt.args.source, tt.args.selector)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinBySelMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MinBySelLsMust_string_rune(t *testing.T) {
	type args struct {
		source   Enumerable[string]
		selector func(string) rune
		lesser   Lesser[rune]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "MinElement2",
			args: args{
				source:   NewEnSlice("xyz", "ab", "abcde", "0"),
				selector: func(s string) rune { return []rune(s)[0] },
				lesser:   LesserFunc[rune](func(r1, r2 rune) bool { return r1 < r2 }),
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinBySelLsMust(tt.args.source, tt.args.selector, tt.args.lesser)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinBySelLsMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MinMust_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "SimpleSequenceNoSelector",
			args: args{
				source: NewEnSlice(5, 10, 6, 2, 13, 8),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinMust(tt.args.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MinMust_float64(t *testing.T) {
	type args struct {
		source Enumerable[float64]
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "SequenceContainingBothInfinities",
			args: args{
				source: NewEnSlice(1., math.Inf(+1), math.Inf(-1)),
			},
			want: math.Inf(-1),
		},
		{name: "SequenceContainingNaN",
			args: args{
				source: NewEnSlice(1., math.Inf(+1), math.NaN(), math.Inf(-1)),
			},
			want: math.Inf(-1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinMust(tt.args.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MaxSelMust_string_int(t *testing.T) {
	type args struct {
		source   Enumerable[string]
		selector func(string) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "SimpleSequenceWithSelector",
			args: args{
				source:   NewEnSlice("xyz", "ab", "abcde", "0"),
				selector: func(s string) int { return len(s) },
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSelMust(tt.args.source, tt.args.selector)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSelMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MaxSelMust_string_rune(t *testing.T) {
	type args struct {
		source   Enumerable[string]
		selector func(string) rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{name: "SimpleSequenceWithSelector",
			args: args{
				source:   NewEnSlice("zyx", "ab", "abcde", "0"),
				selector: func(s string) rune { return []rune(s)[0] },
			},
			want: 'z',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSelMust(tt.args.source, tt.args.selector)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSelMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MaxMust_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "SimpleSequenceWithSelector",
			args: args{
				source: NewEnSlice(5, 10, 6, 2, 13, 8),
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxMust(tt.args.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MaxMust_float64(t *testing.T) {
	type args struct {
		source Enumerable[float64]
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "SimpleSequenceFloat64",
			args: args{
				source: NewEnSlice(-2.5, 2.5, 0.),
			},
			want: 2.5,
		},
		{name: "SequenceContainingBothInfinities",
			args: args{
				source: NewEnSlice(1., math.Inf(+1), math.Inf(-1)),
			},
			want: math.Inf(+1),
		},
		{name: "SequenceContainingNaN",
			args: args{
				source: NewEnSlice(1., math.Inf(+1), math.NaN(), math.Inf(-1)),
			},
			want: math.Inf(+1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxMust(tt.args.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MaxBySelMust_string_int(t *testing.T) {
	type args struct {
		source   Enumerable[string]
		selector func(string) int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "MaxElement",
			args: args{
				source:   NewEnSlice("xyz", "ab", "abcde", "0"),
				selector: func(s string) int { return len(s) },
			},
			want: "abcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxBySelMust(tt.args.source, tt.args.selector)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxBySelMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Example_MinBySelMust() {
	fmt.Println(
		MinBySelMust(
			RangeMust(1, 10),
			func(i int) int { return i * i % 10 },
		),
	)
	fmt.Println(
		MinBySelMust(
			NewEnSlice("one", "two", "three", "four", "five"),
			func(s string) int { return len(s) },
		),
	)
	// Output:
	// 10
	// one
}

func Example_MaxBySelMust() {
	fmt.Println(
		MaxBySelMust(
			RangeMust(1, 10),
			func(i int) int { return i * i % 10 },
		),
	)
	fmt.Println(
		MaxBySelMust(
			NewEnSlice("one", "two", "three", "four", "five"),
			func(s string) int { return len(s) },
		),
	)
	// Output:
	// 3
	// three
}
