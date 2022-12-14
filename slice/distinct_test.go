package slice

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/solsw/go2linq/v2"
)

var (
	testString1 = "test"
	testString2 = "test"
)

func TestDistinct_int(t *testing.T) {
	type args struct {
		source  []int
		equaler go2linq.Equaler[int]
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source: nil,
			},
			want: nil,
		},
		{name: "EmptySource",
			args: args{
				source: []int{},
			},
			want: []int{},
		},
		{name: "Distinct",
			args: args{
				source:  []int{1, 2, 3, 4, 5, 6, 7, 8},
				equaler: go2linq.EqualerFunc[int](func(i1, i2 int) bool { return i1%2 == i2%2 }),
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Distinct(tt.args.source, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("Distinct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinct_string(t *testing.T) {
	type args struct {
		source  []string
		equaler go2linq.Equaler[string]
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source:  nil,
				equaler: go2linq.CaseInsensitiveEqualer,
			},
			want: nil,
		},
		{name: "EmptySource",
			args: args{
				source: []string{},
			},
			want: []string{},
		},
		{name: "NilEqualerUsesDeepEqualer",
			args: args{
				source:  []string{"xyz", testString1, "XYZ", testString2, "def"},
				equaler: nil,
			},
			want: []string{"xyz", testString1, "XYZ", "def"},
		},
		{name: "NonNullEqualer",
			args: args{
				source:  []string{"xyz", testString1, "XYZ", testString2, "def"},
				equaler: go2linq.CaseInsensitiveEqualer,
			},
			want: []string{"xyz", testString1, "def"},
		},
		{name: "Distinct",
			args: args{
				source: []string{"A", "a", "b", "c", "b"},
			},
			want: []string{"A", "a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Distinct(tt.args.source, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistinctEq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctCmp_int(t *testing.T) {
	type args struct {
		source   []int
		comparer go2linq.Comparer[int]
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "NilSource",
			args: args{
				source:   nil,
				comparer: go2linq.Order[int]{},
			},
			want: nil,
		},
		{name: "EmptySource",
			args: args{
				source:   []int{},
				comparer: go2linq.Order[int]{},
			},
			want: []int{},
		},
		{name: "DistinctCmp",
			args: args{
				source:   []int{1, 2, 3, 4},
				comparer: go2linq.Order[int]{},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctCmp(tt.args.source, tt.args.comparer)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistinctCmp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctCmp_string(t *testing.T) {
	type args struct {
		source   []string
		comparer go2linq.Comparer[string]
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "DistinctCmp1",
			args: args{
				source:   []string{"xyz", testString1, "XYZ", testString2, "def"},
				comparer: go2linq.CaseInsensitiveComparer,
			},
			want: []string{"xyz", testString1, "def"},
		},
		{name: "DistinctCmp2",
			args: args{
				source:   []string{"A", "a", "b", "c", "b"},
				comparer: go2linq.CaseInsensitiveComparer,
			},
			want: []string{"A", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistinctCmp(tt.args.source, tt.args.comparer)
			if (err != nil) != tt.wantErr {
				t.Errorf("DistinctCmp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDistinct(b *testing.B) {
	N := 10000
	ii1, _ := Range(1, N)
	ii2, _ := Range(1, N)
	rand.Shuffle(N, reflect.Swapper(ii2))
	ii3 := append(ii1, ii2...)
	var got []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, _ = Distinct(ii3, go2linq.Equaler[int](go2linq.Order[int]{}))
	}
	b.StopTimer()
	if !reflect.DeepEqual(ii1, got) {
		b.Errorf("Distinct() = %v, want %v", got, ii1)
	}
}

func BenchmarkDistinctCmp(b *testing.B) {
	N := 10000
	ii1, _ := Range(1, N)
	ii2, _ := Range(1, N)
	rand.Shuffle(N, reflect.Swapper(ii2))
	ii3 := append(ii1, ii2...)
	var got []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got, _ = DistinctCmp(ii3, go2linq.Comparer[int](go2linq.Order[int]{}))
	}
	b.StopTimer()
	if !reflect.DeepEqual(ii1, got) {
		b.Errorf("DistinctCmp() = %v, want %v", got, ii1)
	}
}
