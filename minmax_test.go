package go2linq

import (
	"fmt"
	"iter"
	"math"
	"reflect"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/MinTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/MaxTest.cs

func TestMin_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "SimpleSequenceNoSelector",
			args: args{
				source: VarAll(5, 10, 6, 2, 13, 8),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Min(tt.args.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin_float64_Inf(t *testing.T) {
	type args struct {
		source iter.Seq[float64]
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "SequenceContainingBothInfinities",
			args: args{
				source: VarAll(1., math.Inf(+1), math.Inf(-1)),
			},
			want: math.Inf(-1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Min(tt.args.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin_float64_NaN(t *testing.T) {
	type args struct {
		source iter.Seq[float64]
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "SequenceContainingNaN",
			args: args{
				source: VarAll(1., math.Inf(+1), math.NaN(), math.Inf(-1)),
			},
			want: math.NaN(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Min(tt.args.source)
			if !math.IsNaN(got) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinSel_string_int(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
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
				source:   VarAll("xyz", "ab", "abcde", "0"),
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
				t.Errorf("MinSel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinSelLs_string_rune(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) rune
		less     func(rune, rune) bool
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
				source:   VarAll("xyz", "ab", "abcde", "0"),
				selector: func(s string) rune { return []rune(s)[0] },
				less:     func(r1, r2 rune) bool { return r1 < r2 },
			},
			want: '0',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinSelLs(tt.args.source, tt.args.selector, tt.args.less)
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

func TestMinBySel_string_int(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "MinElement",
			args: args{
				source:   VarAll("xyz", "ab", "abcde", "0"),
				selector: func(s string) int { return len(s) },
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := MinBySel(tt.args.source, tt.args.selector)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinBySel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinBySelLs_string_rune(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) rune
		less     func(rune, rune) bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "MinElement2",
			args: args{
				source:   VarAll("xyz", "ab", "abcde", "0"),
				selector: func(s string) rune { return []rune(s)[0] },
				less:     func(r1, r2 rune) bool { return r1 < r2 },
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := MinBySelLs(tt.args.source, tt.args.selector, tt.args.less)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinBySelLs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "SimpleSequenceWithSelector",
			args: args{
				source: VarAll(5, 10, 6, 2, 13, 8),
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Max(tt.args.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax_float64(t *testing.T) {
	type args struct {
		source iter.Seq[float64]
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "SimpleSequenceFloat64",
			args: args{
				source: VarAll(-2.5, 2.5, 0.),
			},
			want: 2.5,
		},
		{name: "SequenceContainingBothInfinities",
			args: args{
				source: VarAll(1., math.Inf(+1), math.Inf(-1)),
			},
			want: math.Inf(+1),
		},
		{name: "SequenceContainingNaN",
			args: args{
				source: VarAll(1., math.Inf(+1), math.NaN(), math.Inf(-1)),
			},
			want: math.Inf(+1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Max(tt.args.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSel_string_int(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "SimpleSequenceWithSelector",
			args: args{
				source:   VarAll("xyz", "ab", "abcde", "0"),
				selector: func(s string) int { return len(s) },
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := MaxSel(tt.args.source, tt.args.selector)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSel_string_rune(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{name: "SimpleSequenceWithSelector",
			args: args{
				source:   VarAll("zyx", "ab", "abcde", "0"),
				selector: func(s string) rune { return []rune(s)[0] },
			},
			want: 'z',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := MaxSel(tt.args.source, tt.args.selector)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxBySel_string_int(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "MaxElement",
			args: args{
				source:   VarAll("xyz", "ab", "abcde", "0"),
				selector: func(s string) int { return len(s) },
			},
			want: "abcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := MaxBySel(tt.args.source, tt.args.selector)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxBySel() = %v, want %v", got, tt.want)
			}
		})
	}
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.min
func ExampleMin() {
	doubles := []float64{1.5e+104, 9e+103, -2e+103}
	min, _ := Min(SliceAll(doubles))
	fmt.Printf("The smallest number is %G.\n", min)
	// Output:
	// The smallest number is -2E+103.
}

// see MinEx3 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.min
func ExampleMinLs() {
	pets := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
	}
	minLs, _ := MinLs(
		SliceAll(pets),
		// Compares Pet's ages.
		func(p1, p2 Pet) bool { return p1.Age < p2.Age },
	)
	fmt.Printf("The 'minimum' animal is %s.\n", minLs.Name)
	// Output:
	// The 'minimum' animal is Whiskers.
}

// see MinEx4 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.min
func ExampleMinSel() {
	pets := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
	}
	minSel, _ := MinSel(SliceAll(pets), func(pet Pet) int { return pet.Age })
	fmt.Printf("The youngest animal is age %d.\n", minSel)
	// Output:
	// The youngest animal is age 1.
}

func ExampleMinBySel() {
	minBySel1, _ := MinBySel(
		errorhelper.Must(Range(1, 10)),
		func(i int) int { return i * i % 10 },
	)
	fmt.Println(minBySel1)
	minBySel2, _ := MinBySel(
		VarAll("one", "two", "three", "four", "five"),
		func(s string) int { return len(s) },
	)
	fmt.Println(minBySel2)
	// Output:
	// 10
	// one
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.max
func ExampleMax() {
	longs := []int{4294967296, 466855135, 81125}
	max, _ := Max(SliceAll(longs))
	fmt.Printf("The largest number is %d.\n", max)
	// Output:
	// The largest number is 4294967296.
}

// see MaxEx3 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.max
func ExampleMaxLs() {
	pets := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
	}
	maxLs, _ := MaxLs(
		SliceAll(pets),
		// Compares Pets by summing each Pet's age and name length.
		func(p1, p2 Pet) bool { return p1.Age+len(p1.Name) < p2.Age+len(p2.Name) },
	)
	fmt.Printf("The 'maximum' animal is %s.\n", maxLs.Name)
	// Output:
	// The 'maximum' animal is Barley.
}

// see MaxEx4 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.max
func ExampleMaxSel() {
	pets := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
	}
	maxSel, _ := MaxSel(SliceAll(pets), func(pet Pet) int { return pet.Age + len(pet.Name) })
	fmt.Printf("The maximum pet age plus name length is %d.\n", maxSel)
	// Output:
	// The maximum pet age plus name length is 14.
}

func ExampleMaxBySel() {
	maxBySel1, _ := MaxBySel(
		errorhelper.Must(Range(1, 10)),
		func(i int) int { return i * i % 10 },
	)
	fmt.Println(maxBySel1)
	maxBySel2, _ := MaxBySel(
		VarAll("one", "two", "three", "four", "five"),
		func(s string) int { return len(s) },
	)
	fmt.Println(maxBySel2)
	// Output:
	// 3
	// three
}
