//go:build go1.18

package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/RepeatTest.cs

func Test_Repeat_string(t *testing.T) {
	type args struct {
		element string
		count   int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[string]
		wantErr     bool
		expectedErr error
	}{
		{name: "SimpleRepeat",
			args: args{
				element: "foo",
				count:   3,
			},
			want: NewEnSlice("foo", "foo", "foo"),
		},
		{name: "EmptyRepeat",
			args: args{
				element: "foo",
				count:   0,
			},
			want: Empty[string](),
		},
		{name: "NegativeCount",
			args: args{
				element: "foo",
				count:   -1,
			},
			wantErr:     true,
			expectedErr: ErrNegativeCount,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Repeat(tt.args.element, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repeat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Repeat() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Repeat() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_RepeatMust_int(t *testing.T) {
	type args struct {
		element int
		count   int
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "1",
			args: args{
				element: 0,
				count:   0,
			},
			want: Empty[int](),
		},
		{name: "2",
			args: args{
				element: 2,
				count:   2,
			},
			want: NewEnSlice(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RepeatMust(tt.args.element, tt.args.count)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("RepeatMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

// see the example from Enumerable.Repeat help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.repeat#examples
func ExampleRepeatMust() {
	strs := RepeatMust("I like programming.", 4)
	enr := strs.GetEnumerator()
	for enr.MoveNext() {
		str := enr.Current()
		fmt.Println(str)
	}
	// Output:
	// I like programming.
	// I like programming.
	// I like programming.
	// I like programming.
}
