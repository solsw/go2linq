package slice

import (
	"math"
	"reflect"
	"testing"

	"github.com/solsw/collate"
	"github.com/solsw/go2linq/v3"
)

func TestOrderByKey_string_int(t *testing.T) {
	type args struct {
		source      []string
		keySelector func(string) int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "AscendingSort",
			args: args{
				source:      []string{"the", "quick", "brown", "fox", "jumps"},
				keySelector: func(s string) int { return len(s) },
			},
			want: []string{"the", "fox", "quick", "brown", "jumps"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OrderByKey(tt.args.source, tt.args.keySelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderByKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderByDescKey_string_int(t *testing.T) {
	type args struct {
		source      []string
		keySelector func(string) int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "DescendingSort",
			args: args{
				source:      []string{"the", "quick", "brown", "fox", "jumps"},
				keySelector: func(s string) int { return len(s) },
			},
			want: []string{"quick", "brown", "jumps", "the", "fox"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OrderByDescKey(tt.args.source, tt.args.keySelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderByDescKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderByDescKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderByKey_int_int(t *testing.T) {
	type args struct {
		source      []int
		keySelector func(int) int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "1234",
			args: args{
				source:      []int{4, 1, 3, 2},
				keySelector: go2linq.Identity[int],
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OrderByKey(tt.args.source, tt.args.keySelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderByKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderByDescKey_string_rune(t *testing.T) {
	type args struct {
		source      []string
		keySelector func(string) rune
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "DescendingSort",
			args: args{
				source:      []string{"the", "quick", "brown", "fox", "jumps"},
				keySelector: func(s string) rune { return []rune(s)[0] },
			},
			want: []string{"the", "quick", "jumps", "fox", "brown"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OrderByDescKey(tt.args.source, tt.args.keySelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderByDescKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderByDescKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderByKeyLs_intint(t *testing.T) {
	type args struct {
		source      []elel[int]
		keySelector func(elel[int]) int
		lesser      collate.Lesser[int]
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "SimpleUniqueKeys",
			args: args{
				source:      []elel[int]{{1, 10}, {2, 12}, {3, 11}},
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser:      collate.Order[int]{},
			},
			want: []int{1, 3, 2},
		},
		{name: "OrderingIsStable",
			args: args{
				source:      []elel[int]{{1, 10}, {2, 11}, {3, 11}, {4, 10}},
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser:      collate.Order[int]{},
			},
			want: []int{1, 4, 2, 3},
		},
		{name: "CustomLess",
			args: args{
				source:      []elel[int]{{1, 15}, {2, -13}, {3, 11}},
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser: collate.LesserFunc[int](func(i1, i2 int) bool {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					return f1 < f2
				}),
			},
			want: []int{3, 2, 1},
		},
		{name: "CustomComparer",
			args: args{
				source:      []elel[int]{{1, 15}, {2, -13}, {3, 11}},
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser: collate.ComparerFunc[int](func(i1, i2 int) int {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					switch {
					case f1 < f2:
						return -1
					case f1 > f2:
						return 1
					}
					return 0
				}),
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0, err := OrderByKeyLs(tt.args.source, tt.args.keySelector, tt.args.lesser)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderByKeyLs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, _ := Select(got0, func(e elel[int]) int { return e.e1 })
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderByKeyLs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderByDescKeyLs_intint(t *testing.T) {
	type args struct {
		source      []elel[int]
		keySelector func(elel[int]) int
		lesser      collate.Lesser[int]
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "SimpleUniqueKeys",
			args: args{
				source:      []elel[int]{{1, 10}, {2, 12}, {3, 11}},
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser:      collate.Order[int]{},
			},
			want: []int{2, 3, 1},
		},
		{name: "OrderingIsStable",
			args: args{
				source:      []elel[int]{{1, 10}, {2, 11}, {3, 11}, {4, 10}},
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser:      collate.Order[int]{},
			},
			want: []int{2, 3, 1, 4},
		},
		{name: "CustomLess",
			args: args{
				source:      []elel[int]{{1, 15}, {2, -13}, {3, 11}},
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser: collate.LesserFunc[int](func(i1, i2 int) bool {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					return f1 < f2
				}),
			},
			want: []int{1, 2, 3},
		},
		{name: "CustomComparer",
			args: args{
				source:      []elel[int]{{1, 15}, {2, -13}, {3, 11}},
				keySelector: func(e elel[int]) int { return e.e2 },
				lesser: collate.ComparerFunc[int](func(i1, i2 int) int {
					f1 := math.Abs(float64(i1))
					f2 := math.Abs(float64(i2))
					switch {
					case f1 < f2:
						return -1
					case f1 > f2:
						return 1
					}
					return 0
				}),
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0, err := OrderByDescKeyLs(tt.args.source, tt.args.keySelector, tt.args.lesser)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderByDescKeyLs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, _ := Select(got0, func(e elel[int]) int { return e.e1 })
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderByDescKeyLs() = %v, want %v", got, tt.want)
			}
		})
	}
}
