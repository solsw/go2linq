package go2linq

import (
	"fmt"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ReverseTest.cs

func TestReverse_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[int]
		wantErr bool
	}{
		{name: "EmptyInput",
			args: args{
				source: Empty[int](),
			},
			want: Empty[int](),
		},
		{name: "ReversedRange",
			args: args{
				source: errorhelper.Must(Range(5, 5)),
			},
			want: VarAll(9, 8, 7, 6, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Reverse(tt.args.source)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Reverse() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestReverse_string(t *testing.T) {
	type args struct {
		source iter.Seq[string]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "ReversedStrs",
			args: args{
				source: VarAll("one", "two", "three", "four", "five"),
			},
			want: VarAll("five", "four", "three", "two", "one"),
		},
		{name: "1",
			args: args{
				source: VarAll("1"),
			},
			want: VarAll("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Reverse(tt.args.source)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Reverse() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

// example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.reverse#examples
func ExampleReverse() {
	apple := []string{"a", "p", "p", "l", "e"}
	reverse, _ := Reverse(SliceAll(apple))
	for num := range reverse {
		fmt.Print(num)
	}
	fmt.Println()
	// Output:
	// elppa
}
