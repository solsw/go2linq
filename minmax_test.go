//go:build go1.18

package go2linq

import (
	"math"
	"reflect"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/MinTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/MaxTest.cs

func Test_MinErr_string_int(t *testing.T) {
	type args struct {
		source   Enumerator[string]
		selector func(string) int
		lesser   Lesser[int]
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NilSelector",
			args: args{
				source: Empty[string](),
				lesser: IntLesser,
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "EmptySequenceWithSelector",
			args: args{
				source:   Empty[string](),
				selector: func(s string) int { return len(s) },
				lesser:   IntLesser,
			},
			wantErr:     true,
			expectedErr: ErrEmptySource,
		},
		{name: "SimpleSequenceWithSelector",
			args: args{
				source:   NewOnSlice("xyz", "ab", "abcde", "0"),
				selector: func(s string) int { return len(s) },
				lesser:   IntLesser,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinErr(tt.args.source, tt.args.selector, tt.args.lesser)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("MinErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MinErr_string_rune(t *testing.T) {
	type args struct {
		source   Enumerator[string]
		selector func(string) rune
		lesser   Lesser[rune]
	}
	tests := []struct {
		name        string
		args        args
		want        rune
		wantErr     bool
		expectedErr error
	}{
		{name: "SimpleSequenceWithSelector2",
			args: args{
				source:   NewOnSlice("xyz", "ab", "abcde", "0"),
				selector: func(s string) rune { return []rune(s)[0] },
				lesser:   LesserFunc[rune](func(r1, r2 rune) bool { return r1 < r2 }),
			},
			want: '0',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinErr(tt.args.source, tt.args.selector, tt.args.lesser)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("MinErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MinEl_string_int(t *testing.T) {
	type args struct {
		source   Enumerator[string]
		selector func(string) int
		lesser   Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "MinElement",
			args: args{
				source:   NewOnSlice("xyz", "ab", "abcde", "0"),
				selector: func(s string) int { return len(s) },
				lesser:   IntLesser,
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinEl(tt.args.source, tt.args.selector, tt.args.lesser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinEl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MinEl_string_rune(t *testing.T) {
	type args struct {
		source   Enumerator[string]
		selector func(string) rune
		lesser   Lesser[rune]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "MinElement2",
			args: args{
				source:   NewOnSlice("xyz", "ab", "abcde", "0"),
				selector: func(s string) rune { return []rune(s)[0] },
				lesser:   LesserFunc[rune](func(r1, r2 rune) bool { return r1 < r2 }),
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinEl(tt.args.source, tt.args.selector, tt.args.lesser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinEl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Min_int(t *testing.T) {
	type args struct {
		source   Enumerator[int]
		selector func(int) int
		lesser   Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "SimpleSequenceNoSelector",
			args: args{
				source:   NewOnSlice(5, 10, 6, 2, 13, 8),
				selector: Identity[int],
				lesser:   IntLesser,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.source, tt.args.selector, tt.args.lesser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Min_float64(t *testing.T) {
	type args struct {
		source   Enumerator[float64]
		selector func(float64) float64
		lesser   Lesser[float64]
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "SequenceContainingBothInfinities",
			args: args{
				source:   NewOnSlice(1., math.Inf(+1), math.Inf(-1)),
				selector: Identity[float64],
				lesser:   Float64Lesser,
			},
			want: math.Inf(-1),
		},
		{name: "SequenceContainingNaN",
			args: args{
				source:   NewOnSlice(1., math.Inf(+1), math.NaN(), math.Inf(-1)),
				selector: Identity[float64],
				lesser:   Float64Lesser,
			},
			want: math.Inf(-1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.source, tt.args.selector, tt.args.lesser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Max_string_int(t *testing.T) {
	type args struct {
		source   Enumerator[string]
		selector func(string) int
		lesser   Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "SimpleSequenceWithSelector",
			args: args{
				source:   NewOnSlice("xyz", "ab", "abcde", "0"),
				selector: func(s string) int { return len(s) },
				lesser:   IntLesser,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.source, tt.args.selector, tt.args.lesser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Max_string_rune(t *testing.T) {
	type args struct {
		source   Enumerator[string]
		selector func(string) rune
		lesser   Lesser[rune]
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{name: "SimpleSequenceWithSelector",
			args: args{
				source:   NewOnSlice("zyx", "ab", "abcde", "0"),
				selector: func(s string) rune { return []rune(s)[0] },
				lesser:   LesserFunc[rune](func(r1, r2 rune) bool { return r1 < r2 }),
			},
			want: 'z',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.source, tt.args.selector, tt.args.lesser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Max_int(t *testing.T) {
	type args struct {
		source   Enumerator[int]
		selector func(int) int
		lesser   Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "SimpleSequenceWithSelector",
			args: args{
				source:   NewOnSlice(5, 10, 6, 2, 13, 8),
				selector: Identity[int],
				lesser:   IntLesser,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.source, tt.args.selector, tt.args.lesser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Max_float64(t *testing.T) {
	type args struct {
		source   Enumerator[float64]
		selector func(float64) float64
		lesser   Lesser[float64]
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "SimpleSequenceFloat64",
			args: args{
				source:   NewOnSlice(-2.5, 2.5, 0.),
				selector: Identity[float64],
				lesser:   Float64Lesser,
			},
			want: 2.5,
		},
		{name: "SequenceContainingBothInfinities",
			args: args{
				source:   NewOnSlice(1., math.Inf(+1), math.Inf(-1)),
				selector: Identity[float64],
				lesser:   Float64Lesser,
			},
			want: math.Inf(+1),
		},
		{name: "SequenceContainingNaN",
			args: args{
				source:   NewOnSlice(1., math.Inf(+1), math.NaN(), math.Inf(-1)),
				selector: Identity[float64],
				lesser:   Float64Lesser,
			},
			want: math.Inf(+1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.source, tt.args.selector, tt.args.lesser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MaxEl_string_int(t *testing.T) {
	type args struct {
		source   Enumerator[string]
		selector func(string) int
		lesser   Lesser[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "MaxElement",
			args: args{
				source:   NewOnSlice("xyz", "ab", "abcde", "0"),
				selector: func(s string) int { return len(s) },
				lesser:   IntLesser,
			},
			want: "abcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxEl(tt.args.source, tt.args.selector, tt.args.lesser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxEl() = %v, want %v", got, tt.want)
			}
		})
	}
}
