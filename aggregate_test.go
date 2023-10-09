package go2linq

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AggregateTest.cs

func TestAggregate_int(t *testing.T) {
	type args struct {
		source      Enumerable[int]
		accumulator func(int, int) int
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceUnseeded",
			args: args{
				source:      nil,
				accumulator: func(ag, el int) int { return ag + el },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullFuncUnseeded",
			args: args{
				source:      NewEnSlice(1, 3),
				accumulator: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilAccumulator,
		},
		{name: "UnseededAggregation",
			args: args{
				source:      NewEnSlice(1, 4, 5),
				accumulator: func(ag, el int) int { return ag*2 + el },
			},
			want: 17,
		},
		{name: "EmptySequenceUnseeded",
			args: args{
				source:      Empty[int](),
				accumulator: func(ag, el int) int { return ag + el },
			},
			wantErr:     true,
			expectedErr: ErrEmptySource,
		},
		{name: "UnseededSingleElementAggregation",
			args: args{
				source:      NewEnSlice(1),
				accumulator: func(ag, el int) int { return ag*2 + el },
			},
			want: 1,
		},
		{name: "FirstElementOfInputIsUsedAsSeedForUnseededOverload",
			args: args{
				source:      NewEnSlice(5, 3, 2),
				accumulator: func(ag, el int) int { return ag * el },
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Aggregate(tt.args.source, tt.args.accumulator)
			if (err != nil) != tt.wantErr {
				t.Errorf("Aggregate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Aggregate() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Aggregate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregateSeed_int_int(t *testing.T) {
	type args struct {
		source      Enumerable[int]
		seed        int
		accumulator func(int, int) int
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceSeeded",
			args: args{
				source:      nil,
				seed:        5,
				accumulator: func(ac, el int) int { return ac + el },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullFuncSeeded",
			args: args{
				source:      NewEnSlice(1, 3),
				seed:        5,
				accumulator: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilAccumulator,
		},
		{name: "SeededAggregation",
			args: args{
				source:      NewEnSlice(1, 4, 5),
				seed:        5,
				accumulator: func(ac, el int) int { return ac*2 + el },
			},
			want: 57,
		},
		{name: "EmptySequenceSeeded",
			args: args{
				source:      Empty[int](),
				seed:        5,
				accumulator: func(ac, el int) int { return ac + el },
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateSeed(tt.args.source, tt.args.seed, tt.args.accumulator)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateSeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("AggregateSeed() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregateSeedMust_int32_int64(t *testing.T) {
	type args struct {
		source      Enumerable[int32]
		seed        int64
		accumulator func(int64, int32) int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "DifferentSourceAndAccumulatorTypes",
			args: args{
				source:      NewEnSlice(int32(2000000000), int32(2000000000), int32(2000000000)),
				seed:        int64(0),
				accumulator: func(ac int64, el int32) int64 { return ac + int64(el) },
			},
			want: int64(6000000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AggregateSeedMust(tt.args.source, tt.args.seed, tt.args.accumulator)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateSeedMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregateSeedSel_int_int_string(t *testing.T) {
	type args struct {
		source         Enumerable[int]
		seed           int
		accumulator    func(int, int) int
		resultSelector func(int) string
	}
	tests := []struct {
		name        string
		args        args
		want        string
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceSeededWithResultSelector",
			args: args{
				source:         nil,
				seed:           5,
				accumulator:    func(ac, el int) int { return ac + el },
				resultSelector: func(r int) string { return fmt.Sprint(r) },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullFuncSeededWithResultSelector",
			args: args{
				source:         NewEnSlice(1, 3),
				seed:           5,
				accumulator:    nil,
				resultSelector: func(r int) string { return fmt.Sprint(r) },
			},
			wantErr:     true,
			expectedErr: ErrNilAccumulator,
		},
		{name: "NullProjectionSeededWithResultSelector",
			args: args{
				source:         NewEnSlice(1, 3),
				seed:           5,
				accumulator:    func(ac, el int) int { return ac + el },
				resultSelector: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "SeededAggregationWithResultSelector",
			args: args{
				source:         NewEnSlice(1, 4, 5),
				seed:           5,
				accumulator:    func(ac, el int) int { return ac*2 + el },
				resultSelector: func(r int) string { return fmt.Sprint(r) },
			},
			want: "57",
		},
		{name: "EmptySequenceSeededWithResultSelector",
			args: args{
				source:         Empty[int](),
				seed:           5,
				accumulator:    func(ac, el int) int { return ac + el },
				resultSelector: func(r int) string { return fmt.Sprint(r) },
			},
			want: "5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateSeedSel(tt.args.source, tt.args.seed, tt.args.accumulator, tt.args.resultSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateSeedSel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("AggregateSeedSel() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateSeedSel() = %v, want %v", got, tt.want)
			}
		})
	}
}

// see the last example from Enumerable.Aggregate help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.aggregate
func ExampleAggregateMust() {
	sentence := "the quick brown fox jumps over the lazy dog"
	// Split the string into individual words.
	words := strings.Fields(sentence)
	// Prepend each word to the beginning of the new sentence to reverse the word order.
	reversed := AggregateMust(
		NewEnSlice(words...),
		func(workingSentence, next string) string { return next + " " + workingSentence },
	)
	fmt.Println(reversed)
	// Output:
	// dog lazy the over jumps fox brown quick the
}

// see the second example from Enumerable.Aggregate help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.aggregate
func ExampleAggregateSeedMust() {
	ints := []int{4, 8, 8, 3, 9, 0, 7, 8, 2}
	// Count the even numbers in the array, using a seed value of 0.
	numEven := AggregateSeedMust(
		NewEnSlice(ints...),
		0,
		func(total, next int) int {
			if next%2 == 0 {
				return total + 1
			}
			return total
		},
	)
	fmt.Printf("The number of even integers is: %d\n", numEven)
	// Output:
	// The number of even integers is: 6
}

// see the first example from Enumerable.Aggregate help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.aggregate
func ExampleAggregateSeedSelMust() {
	fruits := []string{"apple", "mango", "orange", "passionfruit", "grape"}
	// Determine whether any string in the array is longer than "banana".
	longestName := AggregateSeedSelMust(
		NewEnSlice(fruits...),
		"banana",
		func(longest, next string) string {
			if len(next) > len(longest) {
				return next
			}
			return longest
		},
		// Return the final result as an upper case string.
		func(fruit string) string { return strings.ToUpper(fruit) },
	)
	fmt.Printf("The fruit with the longest name is %s.\n", longestName)
	// Output:
	// The fruit with the longest name is PASSIONFRUIT.
}
