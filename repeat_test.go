package go2linq

import (
	"fmt"
	"iter"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/RepeatTest.cs

func TestRepeat_string(t *testing.T) {
	type args struct {
		element string
		count   int
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[string]
		wantErr     bool
		expectedErr error
	}{
		{name: "NegativeCount",
			args: args{
				element: "foo",
				count:   -1,
			},
			wantErr:     true,
			expectedErr: ErrNegativeCount,
		},
		{name: "SimpleRepeat",
			args: args{
				element: "foo",
				count:   3,
			},
			want: VarAll("foo", "foo", "foo"),
		},
		{name: "EmptyRepeat",
			args: args{
				element: "foo",
				count:   0,
			},
			want: Empty[string](),
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
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Repeat() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestRepeat_int(t *testing.T) {
	type args struct {
		element int
		count   int
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[int]
		wantErr bool
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
			want: VarAll(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Repeat(tt.args.element, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repeat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Repeat() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

// example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.repeat#examples
func ExampleRepeat() {
	ss, _ := Repeat("I like programming.", 4)
	for s := range ss {
		fmt.Println(s)
	}
	// Output:
	// I like programming.
	// I like programming.
	// I like programming.
	// I like programming.
}
