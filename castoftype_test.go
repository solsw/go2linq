//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/CastTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OfTypeTest.cs

func Test_Cast_interface_int(t *testing.T) {
	type args struct {
		source Enumerable[any]
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSource",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "UnboxToInt",
			args: args{
				source: NewEnSlice[any](10, 30, 50),
			},
			want: NewEnSlice[int](10, 30, 50),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Cast[any, int](tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Cast() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Cast() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_CastMust_interface_string(t *testing.T) {
	type args struct {
		source Enumerable[any]
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[string]
		wantErr     bool
		expectedErr error
	}{
		{name: "SequenceWithAllValidValues",
			args: args{
				source: NewEnSlice[any]("first", "second", "third"),
			},
			want: NewEnSlice[string]("first", "second", "third"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CastMust[any, string](tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("CastMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_OfTypeMust_interface_int(t *testing.T) {
	type args struct {
		source Enumerable[any]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "UnboxToInt",
			args: args{
				source: NewEnSlice[any](10, 30, 50),
			},
			want: NewEnSlice[int](10, 30, 50),
		},
		{name: "OfType",
			args: args{
				source: NewEnSlice[any](1, 2, "two", 3, 3.14, 4, nil),
			},
			want: NewEnSlice[int](1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OfTypeMust[any, int](tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("OfTypeMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_OfTypeMust_interface_string(t *testing.T) {
	type args struct {
		source Enumerable[any]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "SequenceWithAllValidValues",
			args: args{
				source: NewEnSlice[any]("first", "second", "third"),
			},
			want: NewEnSlice[string]("first", "second", "third"),
		},
		{name: "NullsAreExcluded",
			args: args{
				source: NewEnSlice[any]("first", nil, "third"),
			},
			want: NewEnSlice[string]("first", "third"),
		},
		{name: "WrongElementTypesAreIgnored",
			args: args{
				source: NewEnSlice[any]("first", any(1), "third"),
			},
			want: NewEnSlice[string]("first", "third"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OfTypeMust[any, string](tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("OfTypeMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_OfTypeMust_interface_int64(t *testing.T) {
	type args struct {
		source Enumerable[any]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int64]
	}{
		{name: "UnboxingWithWrongElementTypes",
			args: args{
				source: NewEnSlice[any](int64(100), 100, int64(300)),
			},
			want: NewEnSlice[int64](int64(100), int64(300)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OfTypeMust[any, int64](tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("OfTypeMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}
