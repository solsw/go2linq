package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/CountTest.cs

func TestCount_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NonCollectionCount",
			args: args{
				source: RangeMust(2, 5),
			},
			want: 5,
		},
		{name: "0",
			args: args{
				source: Empty[int](),
			},
			want: 0,
		},
		{name: "NullSourceThrowsArgumentNullException",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Count(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Count() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountMust_string(t *testing.T) {
	type args struct {
		source Enumerable[string]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1",
			args: args{
				source: NewEnSlice("zero", "one", "two", "three", "four", "five"),
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountMust(tt.args.source)
			if got != tt.want {
				t.Errorf("CountMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountPred_int(t *testing.T) {
	type args struct {
		source    Enumerable[int]
		predicate func(int) bool
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "PredicatedNullSourceThrowsArgumentNullException",
			args: args{
				predicate: func(x int) bool { return x == 1 },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "PredicatedNullPredicateThrowsArgumentNullException",
			args: args{
				source: NewEnSlice(3, 5, 20, 15),
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "PredicatedCount",
			args: args{
				source:    RangeMust(2, 5),
				predicate: func(x int) bool { return x%2 == 0 },
			},
			want: 3,
		},
		{name: "11",
			args: args{
				source:    NewEnSlice(1, 2, 3, 4),
				predicate: func(int) bool { return false },
			},
			want: 0,
		},
		{name: "12",
			args: args{
				source:    NewEnSlice(1, 2, 3, 4),
				predicate: func(int) bool { return true },
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountPred(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountPred() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("CountPred() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("CountPred() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountPredMust_string(t *testing.T) {
	type args struct {
		source    Enumerable[string]
		predicate func(string) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "21",
			args: args{
				source:    NewEnSlice("one", "two", "three", "four"),
				predicate: func(s string) bool { return len(s) == 3 },
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountPredMust(tt.args.source, tt.args.predicate)
			if got != tt.want {
				t.Errorf("CountPredMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

// see the first example from Enumerable.Count help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.count
func ExampleCountMust() {
	fruits := NewEnSlice("apple", "banana", "mango", "orange", "passionfruit", "grape")
	numberOfFruits := CountMust(fruits)
	fmt.Printf("There are %d fruits in the collection.\n", numberOfFruits)
	// Output:
	// There are 6 fruits in the collection.
}

// see CountEx2 example from Enumerable.Count help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.count
func ExampleCountPredMust_ex1() {
	pets := NewEnSlice(
		Pet{Name: "Barley", Vaccinated: true},
		Pet{Name: "Boots", Vaccinated: false},
		Pet{Name: "Whiskers", Vaccinated: false},
	)
	numberUnvaccinated := CountPredMust(pets, func(p Pet) bool { return p.Vaccinated == false })
	fmt.Printf("There are %d unvaccinated animals.\n", numberUnvaccinated)
	// Output:
	// There are 2 unvaccinated animals.
}

// see LongCountEx2 example from Enumerable.LongCount help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.longcount
func ExampleCountPredMust_ex2() {
	pets := NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	const Age = 3
	count := CountPredMust(pets, func(pet Pet) bool { return pet.Age > Age })
	fmt.Printf("There are %d animals over age %d.\n", count, Age)
	// Output:
	// There are 2 animals over age 3.
}
