package go2linq

import (
	"cmp"
	"fmt"
	"iter"
	"math/rand"
	"reflect"
	"strings"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/DistinctTest.cs

var (
	testString1 = "test"
	testString2 = "test"
)

func TestDistinct_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "NilSource",
			args: args{
				source: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Distinct(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Distinct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Distinct() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Distinct() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestDistinct_string(t *testing.T) {
	type args struct {
		source iter.Seq[string]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "1",
			args: args{
				source: VarAll("A", "a", "b", "c", "b"),
			},
			want: VarAll("A", "a", "b", "c"),
		},
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#distinct-and-distinctby
		{name: "Distinct",
			args: args{
				source: VarAll("Mercury", "Venus", "Venus", "Earth", "Mars", "Earth"),
			},
			want: VarAll("Mercury", "Venus", "Earth", "Mars"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Distinct(tt.args.source)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Distinct() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestDistinctEq_string(t *testing.T) {
	type args struct {
		source iter.Seq[string]
		equal  func(string, string) bool
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[string]
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceWithComparer",
			args: args{
				source: nil,
				equal:  CaseInsensitiveEqual,
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NilEqual",
			args: args{
				source: VarAll("xyz", testString1, "XYZ", testString2, "def"),
				equal:  nil,
			},
			wantErr:     true,
			expectedErr: ErrNilEqual,
		},
		{name: "SimpleDistinctEq",
			args: args{
				source: VarAll("xyz", testString1, "XYZ", testString2, "def"),
				equal:  generichelper.DeepEqual[string],
			},
			want: VarAll("xyz", testString1, "XYZ", "def"),
		},
		{name: "1",
			args: args{
				source: VarAll("xyz", testString1, "XYZ", testString2, "def"),
				equal:  CaseInsensitiveEqual,
			},
			want: VarAll("xyz", testString1, "def"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctEq(tt.args.source, tt.args.equal)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistinctEq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("DistinctEq() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("DistinctEq() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestDistinctCmp_string(t *testing.T) {
	type args struct {
		source  iter.Seq[string]
		compare func(string, string) int
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "DistinctStringsWithCaseInsensitiveComparer",
			args: args{
				source:  VarAll("xyz", testString1, "XYZ", testString2, "def"),
				compare: CaseInsensitiveCompare,
			},
			want: VarAll("xyz", testString1, "def"),
		},
		{name: "3",
			args: args{
				source:  VarAll("A", "a", "b", "c", "b"),
				compare: CaseInsensitiveCompare,
			},
			want: VarAll("A", "b", "c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DistinctCmp(tt.args.source, tt.args.compare)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("DistinctCmp() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestDistinctCmp_int(t *testing.T) {
	type args struct {
		source  iter.Seq[int]
		compare func(int, int) int
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "EmptyEnumerable",
			args: args{
				source:  Empty[int](),
				compare: cmp.Compare[int],
			},
			want: Empty[int](),
		},
		{name: "1",
			args: args{
				source:  VarAll(1, 2, 3, 4),
				compare: cmp.Compare[int],
			},
			want: VarAll(1, 2, 3, 4),
		},
		{name: "2",
			args: args{
				source:  errorhelper.Must(Concat(VarAll(1, 2, 3, 4), VarAll(1, 2, 3, 4))),
				compare: cmp.Compare[int],
			},
			want: VarAll(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DistinctCmp(tt.args.source, tt.args.compare)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("DistinctCmp() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func BenchmarkDistinctEq(b *testing.B) {
	N := 10000
	rng := errorhelper.Must(Range(1, N))
	slc, _ := ToSlice(errorhelper.Must(Range(1, N)))
	rand.Shuffle(N, reflect.Swapper(slc))
	concat, _ := Concat(rng, SliceAll(slc))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, _ := DistinctEq(concat, generichelper.DeepEqual[int])
		// SequenceEqual is measured because the sequence must be enumerated to obtain the results
		equal, _ := SequenceEqual(got, rng)
		if !equal {
			b.Errorf("DistinctEq() = %v, want %v", StringDef(got), StringDef(rng))
		}
	}
}

func BenchmarkDistinctCmp(b *testing.B) {
	N := 10000
	rng := errorhelper.Must(Range(1, N))
	slc, _ := ToSlice(errorhelper.Must(Range(1, N)))
	rand.Shuffle(N, reflect.Swapper(slc))
	concat, _ := Concat(rng, SliceAll(slc))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, _ := DistinctCmp(concat, cmp.Compare[int])
		// SequenceEqual is measured because the sequence must be enumerated to obtain the results
		equal, _ := SequenceEqual(got, rng)
		if !equal {
			b.Errorf("DistinctCmp() = %v, want %v", StringDef(got), StringDef(rng))
		}
	}
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinct
func ExampleDistinct() {
	ages := []int{21, 46, 46, 55, 17, 21, 55, 55}
	distinct, _ := Distinct(SliceAll(ages))
	fmt.Println("Distinct ages:")
	for age := range distinct {
		fmt.Println(age)
	}
	// Output:
	// Distinct ages:
	// 21
	// 46
	// 55
	// 17
}

// last two examples from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinct
func ExampleDistinctEq() {
	products := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
		{Name: "Apple", Code: 9},
		{Name: "lemon", Code: 12},
	}
	//Exclude duplicates.
	distinctEq, _ := DistinctEq(SliceAll(products), func(p1, p2 Product) bool {
		return p1.Code == p2.Code && strings.EqualFold(p1.Name, p2.Name)
	})
	for product := range distinctEq {
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}
	// Output:
	// apple 9
	// orange 4
	// lemon 12
}
