//go:build go1.18

package go2linq

import (
	"reflect"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ElementAtTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ElementAtOrDefaultTest.cs

func Test_ElementAt_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
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
				source: NewOnSlice(1, 2, 3, 4),
				idx:    -1,
			},
			wantErr:     true,
			expectedErr: ErrIndexOutOfRange,
		},
		{name: "OvershootIndex",
			args: args{
				source: NewOnSlice(1, 2, 3, 4),
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
					t.Errorf("ElementAt() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ElementAt_string(t *testing.T) {
	type args struct {
		source Enumerator[string]
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
				source: NewOnSlice("one", "two", "three", "four"),
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
					t.Errorf("ElementAt() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ElementAtOrDefault_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
		idx    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "NegativeIndex",
			args: args{
				source: NewOnSlice(1, 2, 3, 4),
				idx:    -1,
			},
			want: 0,
		},
		{name: "OvershootIndex",
			args: args{
				source: NewOnSlice(1, 2, 3, 4),
				idx:    4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ElementAtOrDefault(tt.args.source, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementAtOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ElementAtOrDefault_string(t *testing.T) {
	type args struct {
		source Enumerator[string]
		idx    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "ValidIndex",
			args: args{
				source: NewOnSlice("one", "two", "three", "four"),
				idx:    2,
			},
			want: "three",
		},
		{name: "InvalidIndex",
			args: args{
				source: NewOnSlice("one", "two", "three", "four"),
				idx:    5,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ElementAtOrDefault(tt.args.source, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementAtOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
