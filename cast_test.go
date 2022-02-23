//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/CastTest.cs

func Test_Cast_any_int(t *testing.T) {
	type args struct {
		source Enumerable[any]
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[int]
		wantErr     bool
		expectedErr error
		wantEnErr   bool
	}{
		{name: "NullSource",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		// {name: "WrongElementType",
		// 	args: args{
		// 		source: NewEnSlice[any](10, 30, "five"),
		// 	},
		// 	want:      NewEnSlice[int](10, 30, 50),
		// 	wantEnErr: true,
		// },
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
				t.Errorf("Cast() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
			// equal, err := SequenceEqualErr(got, tt.want)
			// if (err != nil) != tt.wantEnErr {
			// 	t.Errorf("Cast() error = %v, wantEnErr %v", err, tt.wantEnErr)
			// 	return
			// }
			// if tt.wantEnErr {
			// 	return
			// }
			// if !equal {
			// 	t.Errorf("Cast() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			// }
		})
	}
}

func Test_CastMust_any_string(t *testing.T) {
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
				t.Errorf("CastMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}
