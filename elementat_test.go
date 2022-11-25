package go2linq

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ElementAtTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ElementAtOrDefaultTest.cs

func TestElementAt_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
		idx    int
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
				source: NewEnSlice(1, 2, 3, 4),
				idx:    -1,
			},
			wantErr:     true,
			expectedErr: ErrIndexOutOfRange,
		},
		{name: "OvershootIndex",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
				idx:    4,
			},
			wantErr:     true,
			expectedErr: ErrIndexOutOfRange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ElementAt(tt.args.source, tt.args.idx)
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
		source Enumerable[string]
		idx    int
	}
	tests := []struct {
		name        string
		args        args
		want        string
		wantErr     bool
		expectedErr error
	}{
		{name: "ValidIndex",
			args: args{
				source: NewEnSlice("one", "two", "three", "four"),
				idx:    2,
			},
			want: "three",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ElementAt(tt.args.source, tt.args.idx)
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

func TestElementAtOrDefaultMust_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
		idx    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "NegativeIndex",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
				idx:    -1,
			},
			want: 0,
		},
		{name: "OvershootIndex",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
				idx:    4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ElementAtOrDefaultMust(tt.args.source, tt.args.idx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementAtOrDefaultMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElementAtOrDefaultMust_string(t *testing.T) {
	type args struct {
		source Enumerable[string]
		idx    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "ValidIndex",
			args: args{
				source: NewEnSlice("one", "two", "three", "four"),
				idx:    2,
			},
			want: "three",
		},
		{name: "InvalidIndex",
			args: args{
				source: NewEnSlice("one", "two", "three", "four"),
				idx:    5,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ElementAtOrDefaultMust(tt.args.source, tt.args.idx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementAtOrDefaultMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

// see the example from Enumerable.ElementAt help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementat
func ExampleElementAtMust() {
	names := NewEnSlice("Hartono, Tommy", "Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu")
	rand.Seed(623)
	name := ElementAtMust(names, rand.Intn(CountMust(names)))
	fmt.Printf("The name chosen at random is '%s'.\n", name)
	// Output:
	// The name chosen at random is 'Hedlund, Magnus'.
}

// see the example from Enumerable.ElementAtOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementatordefault
func ExampleElementAtOrDefaultMust() {
	names := NewEnSlice("Hartono, Tommy", "Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu")
	index := 20
	name := ElementAtOrDefaultMust(names, index)
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
