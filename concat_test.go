package go2linq

import (
	"fmt"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ConcatTest.cs

func TestConcat_int(t *testing.T) {
	i4 := VarAll(1, 2, 3, 4)
	rg, _ := Range(1, 4)
	type args struct {
		first  iter.Seq[int]
		second iter.Seq[int]
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[int]
		wantErr bool
	}{
		{name: "Empty",
			args: args{
				first:  Empty[int](),
				second: Empty[int](),
			},
			want: Empty[int](),
		},
		{name: "SemiEmpty1",
			args: args{
				first:  Empty[int](),
				second: VarAll(1, 2, 3, 4),
			},
			want: VarAll(1, 2, 3, 4),
		},
		{name: "SemiEmpty2",
			args: args{
				first:  VarAll(1, 2, 3, 4),
				second: Empty[int](),
			},
			want: VarAll(1, 2, 3, 4),
		},
		{name: "SimpleConcatenation",
			args: args{
				first:  VarAll(1, 2, 3, 4),
				second: VarAll(1, 2, 3, 4),
			},
			want: VarAll(1, 2, 3, 4, 1, 2, 3, 4),
		},
		{name: "SimpleConcatenation2",
			args: args{
				first:  errorhelper.Must(Range(1, 2)),
				second: errorhelper.Must(Repeat(3, 1)),
			},
			want: VarAll(1, 2, 3),
		},
		{name: "SameEnumerableInt",
			args: args{
				first:  i4,
				second: i4,
			},
			want: VarAll(1, 2, 3, 4, 1, 2, 3, 4),
		},
		{name: "SameEnumerableInt2",
			args: args{
				first:  errorhelper.Must(Take(rg, 2)),
				second: errorhelper.Must(Skip(rg, 2)),
			},
			want: VarAll(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Concat(tt.args.first, tt.args.second)
			if (err != nil) != tt.wantErr {
				t.Errorf("Concat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Concat() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestConcat_int2(t *testing.T) {
	type args struct {
		first  iter.Seq[int]
		second iter.Seq[int]
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[int]
		wantErr bool
	}{
		{name: "SecondSequenceIsntAccessedBeforeFirstUse",
			args: args{
				first:  VarAll(1, 2, 3, 4),
				second: errorhelper.Must(Select(VarAll(0, 1), func(x int) int { return 2 / x })),
			},
			want: VarAll(1, 2, 3, 4),
		},
		{name: "NotNeededElementsAreNotAccessed",
			args: args{
				first:  VarAll(1, 2, 3),
				second: errorhelper.Must(Select(VarAll(1, 0), func(x int) int { return 2 / x })),
			},
			want: VarAll(1, 2, 3, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			concat, _ := Concat(tt.args.first, tt.args.second)
			got, _ := Take(concat, 4)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Concat() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestConcat_string(t *testing.T) {
	repeat, _ := Repeat("q", 2)
	rs, _ := Skip(repeat, 1)
	type args struct {
		first  iter.Seq[string]
		second iter.Seq[string]
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[string]
		wantErr bool
	}{
		{name: "Empty",
			args: args{
				first:  Empty[string](),
				second: Empty[string](),
			},
			want: Empty[string](),
		},
		{name: "SemiEmpty",
			args: args{
				first:  Empty[string](),
				second: VarAll("one", "two", "three", "four"),
			},
			want: VarAll("one", "two", "three", "four"),
		},
		{name: "SimpleConcatenation",
			args: args{
				first:  VarAll("a", "b"),
				second: VarAll("c", "d"),
			},
			want: VarAll("a", "b", "c", "d"),
		},
		{name: "SameEnumerableString",
			args: args{
				first:  rs,
				second: rs,
			},
			want: VarAll("q", "q"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Concat(tt.args.first, tt.args.second)
			if (err != nil) != tt.wantErr {
				t.Errorf("Concat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Concat() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

// see ConcatEx1 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.concat#examples
func ExampleConcat() {
	cats := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
	}
	dogs := []Pet{
		{Name: "Bounder", Age: 3},
		{Name: "Snoopy", Age: 14},
		{Name: "Fido", Age: 9},
	}
	concat, _ := Concat(
		errorhelper.Must(Select(
			SliceAll(cats),
			func(cat Pet) string { return cat.Name },
		)),
		errorhelper.Must(Select(
			SliceAll(dogs),
			func(dog Pet) string { return dog.Name },
		)),
	)
	for name := range concat {
		fmt.Println(name)
	}
	// Output:
	// Barley
	// Boots
	// Whiskers
	// Bounder
	// Snoopy
	// Fido
}
