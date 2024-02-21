package go2linq

import (
	"fmt"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SelectTest.cs

func TestSelect_int_int(t *testing.T) {
	var count int
	type args struct {
		source   iter.Seq[int]
		selector func(int) int
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceThrowsNullArgumentException",
			args: args{
				source:   nil,
				selector: func(x int) int { return x + 1 },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullProjectionThrowsNullArgumentException",
			args: args{
				source:   VarAll(1, 3, 7, 9, 10),
				selector: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "SimpleProjection",
			args: args{
				source:   VarAll(1, 5, 2),
				selector: func(x int) int { return x * 2 },
			},
			want: VarAll(2, 10, 4),
		},
		{name: "EmptySource",
			args: args{
				source:   Empty[int](),
				selector: func(x int) int { return x * 2 },
			},
			want: Empty[int](),
		},
		{name: "SideEffectsInProjection1",
			args: args{
				source:   VarAll(3, 2, 1), // Actual values won't be relevant
				selector: func(int) int { count++; return count },
			},
			want: VarAll(1, 2, 3),
		},
		{name: "SideEffectsInProjection2",
			args: args{
				source:   VarAll(1, 2, 3), // Actual values won't be relevant
				selector: func(int) int { count++; return count },
			},
			want: VarAll(4, 5, 6),
		},
		{name: "SideEffectsInProjection3",
			args: args{
				source:   VarAll(1, 2, 3), // Actual values won't be relevant
				selector: func(int) int { count++; return count },
			},
			want: VarAll(11, 12, 13),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Select(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Select() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Select() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
		if tt.name == "SideEffectsInProjection2" {
			// will be used in "SideEffectsInProjection3"
			count = 10
		}
	}
}

func TestSelect_int_string(t *testing.T) {
	type args struct {
		source   iter.Seq[int]
		selector func(int) string
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "SimpleProjectionToDifferentType",
			args: args{
				source:   VarAll(1, 5, 2),
				selector: func(x int) string { return fmt.Sprint(x) },
			},
			want: VarAll("1", "5", "2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Select(tt.args.source, tt.args.selector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Select() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestSelect_string_string(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) string
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/projection-operations#select
		{name: "Select",
			args: args{
				source:   VarAll("an", "apple", "a", "day"),
				selector: func(s string) string { return string([]rune(s)[0]) },
			},
			want: VarAll("a", "a", "a", "d"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Select(tt.args.source, tt.args.selector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Select() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestSelectIdx_int_int(t *testing.T) {
	type args struct {
		source   iter.Seq[int]
		selector func(int, int) int
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "WithIndexNullSourceThrowsNullArgumentException",
			args: args{
				source:   nil,
				selector: func(x, idx int) int { return x + idx },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "WithIndexNullSelectorThrowsNullArgumentException",
			args: args{
				source:   VarAll(1, 3, 7, 9, 10),
				selector: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "WithIndexSimpleProjection",
			args: args{
				source:   VarAll(1, 5, 2),
				selector: func(x, idx int) int { return x + idx*10 },
			},
			want: VarAll(1, 15, 22),
		},
		{name: "WithIndexEmptySource",
			args: args{
				source:   Empty[int](),
				selector: func(x, idx int) int { return x + idx },
			},
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectIdx(tt.args.source, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectIdx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("SelectIdx() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SelectIdx() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.select
func ExampleSelect_ex1() {
	rnge, _ := Range(1, 10)
	squares, _ := Select(rnge, func(x int) int { return x * x })
	next, stop := iter.Pull(squares)
	defer stop()
	for {
		num, ok := next()
		if !ok {
			break
		}
		fmt.Println(num)
	}
	// Output:
	// 1
	// 4
	// 9
	// 16
	// 25
	// 36
	// 49
	// 64
	// 81
	// 100
}

func ExampleSelect_ex2() {
	numbers := []string{"one", "two", "three", "four", "five"}
	select1, _ := Select(
		SliceAll(numbers),
		func(s string) string {
			return string(s[0]) + string(s[len(s)-1])
		},
	)
	fmt.Println(StringDef(select1))
	select2, _ := Select(
		SliceAll(numbers),
		func(s string) string {
			runes := []rune(s)
			reversedRunes, _ := ToSlice(errorhelper.Must(Reverse(SliceAll(runes))))
			return string(reversedRunes)
		},
	)
	fmt.Println(StringDef(select2))
	// Output:
	// [oe to te fr fe]
	// [eno owt eerht ruof evif]
}

// last example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.select

type indexStr struct {
	index int
	str   string
}

func ExampleSelectIdx() {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}
	selectIdx, _ := SelectIdx(
		SliceAll(fruits),
		func(fruit string, index int) indexStr {
			return indexStr{index: index, str: fruit[:index]}
		},
	)
	for is := range selectIdx {
		fmt.Printf("%+v\n", is)
	}
	// Output:
	// {index:0 str:}
	// {index:1 str:b}
	// {index:2 str:ma}
	// {index:3 str:ora}
	// {index:4 str:pass}
	// {index:5 str:grape}
}
