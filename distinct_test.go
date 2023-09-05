package go2linq

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"

	"github.com/solsw/collate"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/DistinctTest.cs

var (
	testString1 = "test"
	testString2 = "test"
)

func TestDistinct_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[int]
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
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestDistinctMust_string(t *testing.T) {
	type args struct {
		source Enumerable[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "1",
			args: args{
				source: NewEnSlice("A", "a", "b", "c", "b"),
			},
			want: NewEnSlice("A", "a", "b", "c"),
		},
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#distinct-and-distinctby
		{name: "Distinct",
			args: args{
				source: NewEnSlice("Mercury", "Venus", "Venus", "Earth", "Mars", "Earth"),
			},
			want: NewEnSlice("Mercury", "Venus", "Earth", "Mars"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DistinctMust(tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DistinctMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestDistinctEq_string(t *testing.T) {
	type args struct {
		source  Enumerable[string]
		equaler collate.Equaler[string]
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[string]
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceWithComparer",
			args: args{
				source:  nil,
				equaler: collate.CaseInsensitiveEqualer,
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullComparerUsesDefault",
			args: args{
				source:  NewEnSlice("xyz", testString1, "XYZ", testString2, "def"),
				equaler: nil,
			},
			want: NewEnSlice("xyz", testString1, "XYZ", "def"),
		},
		{name: "1",
			args: args{
				source:  NewEnSlice("xyz", testString1, "XYZ", testString2, "def"),
				equaler: collate.CaseInsensitiveEqualer,
			},
			want: NewEnSlice("xyz", testString1, "def"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctEq(tt.args.source, tt.args.equaler)
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
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DistinctEq() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestDistinctCmpMust_string(t *testing.T) {
	type args struct {
		source   Enumerable[string]
		comparer collate.Comparer[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "DistinctStringsWithCaseInsensitiveComparer",
			args: args{
				source:   NewEnSlice("xyz", testString1, "XYZ", testString2, "def"),
				comparer: collate.CaseInsensitiveComparer,
			},
			want: NewEnSlice("xyz", testString1, "def"),
		},
		{name: "3",
			args: args{
				source:   NewEnSlice("A", "a", "b", "c", "b"),
				comparer: collate.CaseInsensitiveComparer,
			},
			want: NewEnSlice("A", "b", "c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DistinctCmpMust(tt.args.source, tt.args.comparer)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DistinctCmpMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestDistinctCmpMust_int(t *testing.T) {
	type args struct {
		source   Enumerable[int]
		comparer collate.Comparer[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "EmptyEnumerable",
			args: args{
				source:   Empty[int](),
				comparer: collate.Order[int]{},
			},
			want: Empty[int](),
		},
		{name: "1",
			args: args{
				source:   NewEnSlice(1, 2, 3, 4),
				comparer: collate.Order[int]{},
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
		{name: "2",
			args: args{
				source:   ConcatMust(NewEnSliceEn(1, 2, 3, 4), NewEnSliceEn(1, 2, 3, 4)),
				comparer: collate.Order[int]{},
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DistinctCmpMust(tt.args.source, tt.args.comparer)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("DistinctCmpMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func BenchmarkDistinctEqMust(b *testing.B) {
	N := 10000
	ii1 := RangeMust(1, N)
	ii2 := ToSliceMust(RangeMust(1, N))
	rand.Shuffle(N, reflect.Swapper(ii2))
	ii3 := ConcatMust(ii1, NewEnSliceEn(ii2...))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got := DistinctEqMust(ii3, collate.Order[int]{})
		// SequenceEqual is measured since Enumerable must be enumerated to obtain the results
		if !SequenceEqualMust(ii1, got) {
			b.Errorf("DistinctEqMust() = %v, want %v", got, ii1)
		}
	}
}

func BenchmarkDistinctCmpMust(b *testing.B) {
	N := 10000
	ii1 := RangeMust(1, N)
	ii2 := ToSliceMust(RangeMust(1, N))
	rand.Shuffle(N, reflect.Swapper(ii2))
	ii3 := ConcatMust(ii1, NewEnSliceEn(ii2...))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got := DistinctCmpMust(ii3, collate.Order[int]{})
		// SequenceEqual is measured since Enumerable must be enumerated to obtain the results
		if !SequenceEqualMust(ii1, got) {
			b.Errorf("DistinctCmpMust() = %v, want %v", got, ii1)
		}
	}
}

// see the first example from Enumerable.Distinct help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinct
func ExampleDistinctMust() {
	ages := []int{21, 46, 46, 55, 17, 21, 55, 55}
	distinct := DistinctMust(NewEnSliceEn(ages...))
	fmt.Println("Distinct ages:")
	enr := distinct.GetEnumerator()
	for enr.MoveNext() {
		age := enr.Current()
		fmt.Println(age)
	}
	// Output:
	// Distinct ages:
	// 21
	// 46
	// 55
	// 17
}

// see the last two examples from Enumerable.Distinct help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinct
func ExampleDistinctEqMust() {
	products := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
		{Name: "Apple", Code: 9},
		{Name: "lemon", Code: 12},
	}
	var eqf collate.Equaler[Product] = collate.EqualerFunc[Product](
		func(p1, p2 Product) bool {
			return p1.Code == p2.Code && strings.EqualFold(p1.Name, p2.Name)
		},
	)
	//Exclude duplicates.
	distinctEq := DistinctEqMust(NewEnSliceEn(products...), eqf)
	enr := distinctEq.GetEnumerator()
	for enr.MoveNext() {
		product := enr.Current()
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}
	// Output:
	// apple 9
	// orange 4
	// lemon 12
}
