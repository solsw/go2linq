package go2linq

import (
	"fmt"
	"iter"
	"math"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OrderByTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OrderByDescendingTest.cs

func TestOrderBy_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[int]
		wantErr bool
	}{
		{name: "1234",
			args: args{
				source: VarAll(4, 1, 3, 2),
			},
			want: VarAll(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OrderBy(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("OrderBy() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func ExampleOrderBy() {
	fmt.Println(StringDef[string](
		errorhelper.Must(OrderBy(VarAll("zero", "one", "two", "three", "four", "five"))),
	))
	// Output:
	// [five four one three two zero]
}

// see OrderByEx1 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.orderby
func ExampleOrderByLs() {
	pets := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
	}
	orderByLs, _ := OrderByLs(SliceAll(pets), func(p1, p2 Pet) bool { return p1.Age < p2.Age })
	for pet := range orderByLs {
		fmt.Printf("%s - %d\n", pet.Name, pet.Age)
	}
	// Output:
	// Whiskers - 1
	// Boots - 4
	// Barley - 8
}

// see OrderByDescendingEx1 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending
func ExampleOrderByDescLs() {
	decimals := []float64{6.2, 8.3, 0.5, 1.3, 6.3, 9.7}
	less := func(f1, f2 float64) bool {
		_, fr1 := math.Modf(f1)
		_, fr2 := math.Modf(f2)
		if math.Abs(fr1-fr2) < 0.001 {
			return f1 < f2
		}
		return fr1 < fr2
	}
	orderByDescLs, _ := OrderByDescLs(SliceAll(decimals), less)
	for num := range orderByDescLs {
		fmt.Println(num)
	}
	// Output:
	// 9.7
	// 0.5
	// 8.3
	// 6.3
	// 1.3
	// 6.2
}
