//go:build go1.18

package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SelectManyTest.cs

func Test_SelectMany_int_rune(t *testing.T) {
	type args struct {
		source Enumerator[int]
		selector func(int) Enumerator[rune]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[rune]
	}{
		{name: "SimpleFlatten",
			args: args{
				source: NewOnSlice(3, 5, 20, 15),
				selector: func(x int) Enumerator[rune] {
					return NewOnSlice([]rune(fmt.Sprint(x))...)
				},
			},
			want: NewOnSlice('3', '5', '2', '0', '1', '5'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectMany(tt.args.source, tt.args.selector); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("SelectMany() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_SelectMany_int_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
		selector func(int) Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "1",
			args: args{
				source: NewOnSlice(1, 2, 3, 4),
				selector: func(i int) Enumerator[int] {
					return NewOnSlice(i, i*i)
				},
			},
			want: NewOnSlice(1, 1, 2, 4, 3, 9, 4, 16),
		},
		{name: "2",
			args: args{
				source: NewOnSlice(1, 2, 3, 4),
				selector: func(i int) Enumerator[int] {
					if i%2 == 0 {
						return Empty[int]()
					}
					return NewOnSlice(i, i*i)
				},
			},
			want: NewOnSlice(1, 1, 3, 9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectMany(tt.args.source, tt.args.selector); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("SelectMany() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_SelectManyIdx_int_rune(t *testing.T) {
	type args struct {
		source Enumerator[int]
		selector func(int, int) Enumerator[rune]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[rune]
	}{
		{name: "SimpleFlattenWithIndex",
			args: args{
				source: NewOnSlice(3, 5, 20, 15),
				selector: func(x, index int) Enumerator[rune] {
					return NewOnSlice([]rune(fmt.Sprint(x + index))...)
				},
			},
			want: NewOnSlice('3', '6', '2', '2', '1', '8'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectManyIdx(tt.args.source, tt.args.selector); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("SelectManyIdx() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_SelectManyIdx_int_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
		selector func(int, int) Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "1",
			args: args{
				source: NewOnSlice(1, 2, 3, 4),
				selector: func(i, idx int) Enumerator[int] {
					if idx%2 == 0 {
						return Empty[int]()
					}
					return NewOnSlice(i, i*i)
				},
			},
			want: NewOnSlice(2, 4, 4, 16),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectManyIdx(tt.args.source, tt.args.selector); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("SelectManyIdx() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_SelectManyColl_int_rune_string(t *testing.T) {
	type args struct {
		source Enumerator[int]
		collectionSelector func(int) Enumerator[rune]
		resultSelector func(int, rune) string
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "FlattenWithProjection",
			args: args{
				source: NewOnSlice(3, 5, 20, 15),
				collectionSelector: func(x int) Enumerator[rune] {
					return NewOnSlice([]rune(fmt.Sprint(x))...)
				},
				resultSelector: func(x int, c rune) string {
					return fmt.Sprintf("%d: %s", x, string(c))
				},
			},
			want: NewOnSlice("3: 3", "5: 5", "20: 2", "20: 0", "15: 1", "15: 5"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectManyColl(tt.args.source, tt.args.collectionSelector, tt.args.resultSelector); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("SelectManyColl() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_SelectManyCollIdx_int_rune_string(t *testing.T) {
	type args struct {
		source Enumerator[int]
		collectionSelector func(int, int) Enumerator[rune]
		resultSelector func(int, rune) string
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "FlattenWithProjectionAndIndex",
			args: args{
				source: NewOnSlice(3, 5, 20, 15),
				collectionSelector: func(x, index int) Enumerator[rune] {
					return NewOnSlice([]rune(fmt.Sprint(x + index))...)
				},
				resultSelector: func(x int, c rune) string {
					return fmt.Sprintf("%d: %s", x, string(c))
				},
			},
			want: NewOnSlice("3: 3", "5: 6", "20: 2", "20: 2", "15: 1", "15: 8"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectManyCollIdx(tt.args.source, tt.args.collectionSelector, tt.args.resultSelector); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("SelectManyIdxCollMust() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
