//go:build go1.18

package go2linq

import (
	"reflect"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ToListTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ToArrayTest.cs

func TestToSlice_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name        string
		args        args
		want        []int
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSource",
			args: args{
				source: nil,
			},
			want:        nil,
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "SimpleSlice",
			args: args{
				source: NewEnSlice[int](1, 2, 3, 4),
			},
			want: []int{1, 2, 3, 4},
		},
		{name: "ConversionOfLazilyEvaluatedSequence",
			args: args{
				source: SelectMust(RangeMust(3, 3), func(x int) int { return x * 2 }),
			},
			want: []int{6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToSlice[int](tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("ToSlice() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestToSliceErr_interface_int(t *testing.T) {
// 	type args struct {
// 		source Enumerable[int]
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []int
// 		wantErr bool
// 	}{
// 		{name: "CastExceptionOnWrongElementType",
// 			args: args{
// 				source: CastMust[any, int](NewEnSlice[any](1.0, 2.0, 3.0, 4.0, "five")),
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := EnToSliceErr[int](tt.args.source)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("EnToSliceErr() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("EnToSliceErr() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
