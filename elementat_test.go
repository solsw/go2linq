package go2linq

import (
	"fmt"
	"iter"
	"math/rand"
	"reflect"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ElementAtTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ElementAtOrDefaultTest.cs

func TestElementAt_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
		index  int
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NegativeIndex",
			args: args{
				source: VarAll(1, 2, 3, 4),
				index:  -1,
			},
			wantErr:     true,
			expectedErr: ErrIndexOutOfRange,
		},
		{name: "OvershootIndex",
			args: args{
				source: VarAll(1, 2, 3, 4),
				index:  4,
			},
			wantErr:     true,
			expectedErr: ErrIndexOutOfRange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ElementAt(tt.args.source, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("ElementAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("ElementAt() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElementAt_string(t *testing.T) {
	type args struct {
		source iter.Seq[string]
		index  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "ValidIndex",
			args: args{
				source: VarAll("one", "two", "three", "four"),
				index:  2,
			},
			want: "three",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ElementAt(tt.args.source, tt.args.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElementAtOrDefault_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
		index  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "NegativeIndex",
			args: args{
				source: VarAll(1, 2, 3, 4),
				index:  -1,
			},
			want: 0,
		},
		{name: "OvershootIndex",
			args: args{
				source: VarAll(1, 2, 3, 4),
				index:  4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ElementAtOrDefault(tt.args.source, tt.args.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementAtOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElementAtOrDefault_string(t *testing.T) {
	type args struct {
		source iter.Seq[string]
		index  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "ValidIndex",
			args: args{
				source: VarAll("one", "two", "three", "four"),
				index:  2,
			},
			want: "three",
		},
		{name: "InvalidIndex",
			args: args{
				source: VarAll("one", "two", "three", "four"),
				index:  5,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ElementAtOrDefault(tt.args.source, tt.args.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementAtOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

// example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.elementat
func ExampleElementAt() {
	names := []string{"Hartono, Tommy", "Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu"}
	r := rand.New(rand.NewSource(623))
	name, _ := ElementAt(SliceAll(names), r.Intn(len(names)))
	fmt.Printf("The name chosen at random is '%s'.\n", name)
	// Output:
	// The name chosen at random is 'Hedlund, Magnus'.
}

// example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.elementatordefault
func ExampleElementAtOrDefault() {
	names := []string{"Hartono, Tommy", "Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu"}
	index := 20
	name, _ := ElementAtOrDefault(SliceAll(names), index)
	var what string
	if name == "" {
		what = "<no name at this index>"
	} else {
		what = name
	}
	fmt.Printf("The name chosen at index %d is '%s'.\n", index, what)
	// Output:
	// The name chosen at index 20 is '<no name at this index>'.
}
