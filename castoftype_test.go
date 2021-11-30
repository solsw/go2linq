//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/CastTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OfTypeTest.cs

func Test_Cast_interface_int(t *testing.T) {
	type args struct {
		source Enumerator[interface{}]
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerator[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSource",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "UnboxToInt",
			args: args{
				source: NewOnSlice[interface{}](10, 30, 50),
			},
			want: NewOnSlice[int](10, 30, 50),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Cast[interface{}, int](tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cast() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Cast() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Cast() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Cast_interface_string(t *testing.T) {
	type args struct {
		source Enumerator[interface{}]
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerator[string]
		wantErr     bool
		expectedErr error
	}{
		{name: "SequenceWithAllValidValues",
			args: args{
				source: NewOnSlice[interface{}]("first", "second", "third"),
			},
			want: NewOnSlice[string]("first", "second", "third"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Cast[interface{}, string](tt.args.source); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Cast() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_OfType_interface_int(t *testing.T) {
	type args struct {
		source Enumerator[interface{}]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "UnboxToInt",
			args: args{
				source: NewOnSlice[interface{}](10, 30, 50),
			},
			want: NewOnSlice[int](10, 30, 50),
		},
		{name: "OfType",
			args: args{
				source: NewOnSlice[interface{}](1, 2, "two", 3, 3.14, 4, nil),
			},
			want: NewOnSlice[int](1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := OfType[interface{}, int](tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("OfType() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_OfType_interface_string(t *testing.T) {
	type args struct {
		source Enumerator[interface{}]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "SequenceWithAllValidValues",
			args: args{
				source: NewOnSlice[interface{}]("first", "second", "third"),
			},
			want: NewOnSlice[string]("first", "second", "third"),
		},
		{name: "NullsAreExcluded",
			args: args{
				source: NewOnSlice[interface{}]("first", nil, "third"),
			},
			want: NewOnSlice[string]("first", "third"),
		},
		{name: "WrongElementTypesAreIgnored",
			args: args{
				source: NewOnSlice[interface{}]("first", interface{}(1), "third"),
			},
			want: NewOnSlice[string]("first", "third"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := OfType[interface{}, string](tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("OfType() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_OfType_interface_int64(t *testing.T) {
	type args struct {
		source Enumerator[interface{}]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int64]
	}{
		{name: "UnboxingWithWrongElementTypes",
			args: args{
				source: NewOnSlice[interface{}](int64(100), 100, int64(300)),
			},
			want: NewOnSlice[int64](int64(100), int64(300)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := OfType[interface{}, int64](tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("OfType() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
