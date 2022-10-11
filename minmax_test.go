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

func TestMinSel_string_int(t *testing.T) {
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

func TestMinSelLs_string_rune(t *testing.T) {
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

func TestMinBySelMust_string_int(t *testing.T) {
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

func TestMinBySelLsMust_string_rune(t *testing.T) {
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

func TestMinMust_int(t *testing.T) {
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

func TestMinMust_float64(t *testing.T) {
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

func TestMaxSelMust_string_int(t *testing.T) {
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

func TestMaxSelMust_string_rune(t *testing.T) {
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

func TestMaxMust_int(t *testing.T) {
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

func TestMaxMust_float64(t *testing.T) {
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

func TestMaxBySelMust_string_int(t *testing.T) {
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

// see the first example from Enumerable.Min help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min
func ExampleMinMust() {
	doubles := NewEnSlice(1.5e+104, 9e+103, -2e+103)
	min := MinMust(doubles)
	fmt.Printf("The smallest number is %G.\n", min)
	// Output:
	// The smallest number is -2E+103.
}

// see MinEx3 example from Enumerable.Min help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min
func ExampleMinLsMust() {
	pets := NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	minLs := MinLsMust(pets,
		// Compares Pet's ages.
		Lesser[Pet](LesserFunc[Pet](
			func(p1, p2 Pet) bool { return p1.Age < p2.Age },
		)),
	)
	fmt.Printf("The 'minimum' animal is %s.\n", minLs.Name)
	// Output:
	// The 'minimum' animal is Whiskers.
}

// see MinEx4 example from Enumerable.Min help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min
func ExampleMinSelMust() {
	pets := NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	minSel := MinSelMust(pets, func(pet Pet) int { return pet.Age })
	fmt.Printf("The youngest animal is age %d.\n", minSel)
	// Output:
	// The youngest animal is age 1.
}

func ExampleMinBySelMust() {
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

// see the first example from Enumerable.Max help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max
func ExampleMaxMust() {
	longs := NewEnSlice(4294967296, 466855135, 81125)
	max := MaxMust(longs)
	fmt.Printf("The largest number is %d.\n", max)
	// Output:
	// The largest number is 4294967296.
}

// see MaxEx3 example from Enumerable.Max help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max
func ExampleMaxLsMust() {
	pets := NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	maxLs := MaxLsMust(pets,
		// Compares Pets by summing each Pet's age and name length.
		Lesser[Pet](LesserFunc[Pet](
			func(p1, p2 Pet) bool { return p1.Age+len(p1.Name) < p2.Age+len(p2.Name) },
		)),
	)
	fmt.Printf("The 'maximum' animal is %s.\n", maxLs.Name)
	// Output:
	// The 'maximum' animal is Barley.
}

// see MaxEx4 example from Enumerable.Max help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max
func ExampleMaxSelMust() {
	pets := NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	maxSel := MaxSelMust(pets, func(pet Pet) int { return pet.Age + len(pet.Name) })
	fmt.Printf("The maximum pet age plus name length is %d.\n", maxSel)
	// Output:
	// The maximum pet age plus name length is 14.
}

func ExampleMaxBySelMust() {
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
