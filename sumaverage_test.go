package go2linq

import (
	"fmt"
	"iter"
	"math"
	"reflect"
	"strconv"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SumTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AverageTest.cs

func TestSum(t *testing.T) {
	type args struct {
		source iter.Seq[float64]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "OverflowToNegInfinityFloat64",
			args: args{
				source: VarAll(-math.MaxFloat64, -math.MaxFloat64),
			},
			want: true,
		},
		{name: "OverflowToInfinityFloat64",
			args: args{
				source: VarAll(math.MaxFloat64, math.MaxFloat64),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Sum(tt.args.source)
			var want bool
			switch tt.name {
			case "OverflowToNegInfinityFloat64":
				want = math.IsInf(got, -1)
			case "OverflowToInfinityFloat64":
				want = math.IsInf(got, +1)
			}
			if want != tt.want {
				t.Errorf("IsInf(Sum()) = %v, want %v", want, tt.want)
			}
		})
	}
}

func TestSumSel_string_int(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "EmptySequenceIntWithSelector",
			args: args{
				source:   Empty[string](),
				selector: func(s string) int { return len(s) },
			},
			want: 0,
		},
		{name: "SimpleSumIntWithSelector",
			args: args{
				source:   VarAll("x", "abc", "de"),
				selector: func(s string) int { return len(s) },
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SumSel(tt.args.source, tt.args.selector)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumSel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumSel_string_float64(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "EmptySequenceFloat64WithSelector",
			args: args{
				source:   Empty[string](),
				selector: func(s string) float64 { return float64(len(s)) },
			},
			want: 0,
		},
		{name: "SimpleSumFloat64WithSelector",
			args: args{
				source:   VarAll("x", "abc", "de"),
				selector: func(s string) float64 { return float64(len(s)) },
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SumSel(tt.args.source, tt.args.selector)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumSel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumSel_string_float64IsNaN(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "SimpleSumFloat64WithSelectorWithNan",
			args: args{
				source: VarAll("x", "abc", "de"),
				selector: func(s string) float64 {
					l := len(s)
					if l == 3 {
						return math.NaN()
					}
					return float64(l)
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sumSel, _ := SumSel(tt.args.source, tt.args.selector)
			want := math.IsNaN(sumSel)
			if want != tt.want {
				t.Errorf("IsNaN(SumSel()) = %v, want %v", want, tt.want)
			}
		})
	}
}

func TestSumSel_string_float64IsInf(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "OverflowToInfinityFloat64WithSelector",
			args: args{
				source:   VarAll("x", "y"),
				selector: func(string) float64 { return math.MaxFloat64 },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SumSel(tt.args.source, tt.args.selector)
			want := math.IsInf(got, +1)
			if want != tt.want {
				t.Errorf("IsInf(SumSel()) = %v, want %v", want, tt.want)
			}
		})
	}
}

func TestAverage_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
	}
	tests := []struct {
		name        string
		args        args
		want        float64
		wantErr     bool
		expectedErr error
	}{
		{name: "EmptySequenceIntNoSelector",
			args: args{
				source: Empty[int](),
			},
			wantErr:     true,
			expectedErr: ErrEmptySource,
		},
		{name: "SimpleAverageInt",
			args: args{
				source: VarAll(5, 10, 0, 15),
			},
			want: 7.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Average(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Average() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Average() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAverage_float64IsInf(t *testing.T) {
	type args struct {
		source iter.Seq[float64]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Float64OverflowsToInfinity",
			args: args{
				source: VarAll(math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64),
			},
			want: true,
		},
		{name: "Float64OverflowsToNegInfinity",
			args: args{
				source: VarAll(-math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64, math.MaxFloat64),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Average(tt.args.source)
			var want bool
			switch tt.name {
			case "Float64OverflowsToInfinity":
				want = math.IsInf(got, +1)
			case "Float64OverflowsToNegInfinity":
				want = math.IsInf(got, -1)
			}
			if want != tt.want {
				t.Errorf("IsInf(Average()) = %v, want %v", want, tt.want)
			}
		})
	}
}

func TestAverageSel_string_int(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) int
	}
	tests := []struct {
		name        string
		args        args
		want        float64
		wantErr     bool
		expectedErr error
	}{
		{name: "SourceStrNilSelector",
			args: args{
				source: VarAll("one", "two", "three", "four"),
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "EmptySequenceIntWithSelector",
			args: args{
				source:   Empty[string](),
				selector: func(s string) int { return len(s) },
			},
			wantErr:     true,
			expectedErr: ErrEmptySource,
		},
		{name: "SimpleAverageIntWithSelector",
			args: args{
				source:   VarAll("", "abcd", "a", "b"),
				selector: func(s string) int { return len(s) },
			},
			want: 1.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AverageSel(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("AverageSel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("AverageSel() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("AverageSel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAverageSel_string_float64IsNaN(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "SequenceContainingNan",
			args: args{
				source: VarAll("x", "abc", "de"),
				selector: func(s string) float64 {
					l := len(s)
					if l == 3 {
						return math.NaN()
					}
					return float64(l)
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := AverageSel(tt.args.source, tt.args.selector)
			want := math.IsNaN(got)
			if want != tt.want {
				t.Errorf("IsNaN(AverageSel()) = %v, want %v", want, tt.want)
			}
		})
	}
}

// example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sum
func ExampleSum() {
	numbers := []float64{43.68, 1.25, 583.7, 6.5}
	sum, _ := Sum(SliceAll(numbers))
	fmt.Printf("The sum of the numbers is %g.\n", sum)
	// Output:
	// The sum of the numbers is 635.13.
}

// SumEx1 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sum
func ExampleSumSel() {
	packages := []Package{
		{Company: "Coho Vineyard", Weight: 25.2},
		{Company: "Lucerne Publishing", Weight: 18.7},
		{Company: "Wingtip Toys", Weight: 6.0},
		{Company: "Adventure Works", Weight: 33.8},
	}
	totalWeight, _ := SumSel(
		SliceAll(packages),
		func(pkg Package) float64 { return pkg.Weight },
	)
	fmt.Printf("The total weight of the packages is: %.1f\n", totalWeight)
	// Output:
	// The total weight of the packages is: 83.7
}

// example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.average
func ExampleAverage_ex1() {
	grades := []int{78, 92, 100, 37, 81}
	average, _ := Average(SliceAll(grades))
	fmt.Printf("The average grade is %g.\n", average)
	// Output:
	// The average grade is 77.6.
}

// example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.average
func ExampleAverage_ex2() {
	numbers := []string{"10007", "37", "299846234235"}
	average, _ := AverageSel(
		SliceAll(numbers),
		func(e string) int {
			r, _ := strconv.Atoi(e)
			return r
		},
	)
	fmt.Printf("The average is %.f.\n", average)
	// Output:
	// The average is 99948748093.
}

// example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.average
func ExampleAverageSel() {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}
	average, _ := AverageSel(
		SliceAll(fruits),
		func(e string) int { return len(e) },
	)
	fmt.Printf("The average string length is %g.\n", average)
	// Output:
	// The average string length is 6.5.
}
