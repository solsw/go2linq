package slice

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/solsw/go2linq/v2"
)

func TestJoinMust_string_rune(t *testing.T) {
	en := []string{"fs", "sf", "ff", "ss"}
	type args struct {
		outer            []string
		inner            []string
		outerKeySelector func(string) rune
		innerKeySelector func(string) rune
		resultSelector   func(string, string) string
		equaler          go2linq.Equaler[rune]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "SimpleJoin",
			args: args{
				outer:            []string{"first", "second", "third"},
				inner:            []string{"essence", "offer", "eating", "psalm"},
				outerKeySelector: func(oel string) rune { return ([]rune(oel))[0] },
				innerKeySelector: func(iel string) rune { return ([]rune(iel))[1] },
				resultSelector:   func(oel, iel string) string { return oel + ":" + iel },
			},
			want: []string{"first:offer", "second:essence", "second:psalm"},
		},
		{name: "SameSlice",
			args: args{
				outer:            en,
				inner:            en,
				outerKeySelector: func(oel string) rune { return ([]rune(oel))[0] },
				innerKeySelector: func(iel string) rune { return ([]rune(iel))[1] },
				resultSelector:   func(oel, iel string) string { return oel + ":" + iel },
			},
			want: []string{"fs:sf", "fs:ff", "sf:fs", "sf:ss", "ff:sf", "ff:ff", "ss:fs", "ss:ss"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinMust(tt.args.outer, tt.args.inner, tt.args.outerKeySelector, tt.args.innerKeySelector, tt.args.resultSelector, tt.args.equaler)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JoinMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinMust_string(t *testing.T) {
	type args struct {
		outer            []string
		inner            []string
		outerKeySelector func(string) string
		innerKeySelector func(string) string
		resultSelector   func(string, string) string
		equaler          go2linq.Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "CustomComparer",
			args: args{
				outer: []string{"ABCxxx", "abcyyy", "defzzz", "ghizzz"},
				inner: []string{"000abc", "111gHi", "222333"},
				outerKeySelector: func(oel string) string {
					return strings.ToLower(oel[:3])
				},
				innerKeySelector: func(iel string) string {
					return strings.ToLower(iel[3:])
				},
				resultSelector: func(oel, iel string) string { return oel + ":" + iel },
			},
			want: []string{"ABCxxx:000abc", "abcyyy:000abc", "ghizzz:111gHi"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinMust(tt.args.outer, tt.args.inner, tt.args.outerKeySelector, tt.args.innerKeySelector, tt.args.resultSelector, tt.args.equaler)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JoinMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinMust_CustomComparer(t *testing.T) {
	outer := []string{"ABCxxx", "abcyyy", "defzzz", "ghizzz"}
	inner := []string{"000abc", "111gHi", "222333"}
	got := JoinMust(outer, inner,
		func(oel string) string { return oel[:3] },
		func(iel string) string { return iel[3:] },
		func(oel, iel string) string { return oel + ":" + iel },
		go2linq.CaseInsensitiveEqualer,
	)
	want := []string{"ABCxxx:000abc", "abcyyy:000abc", "ghizzz:111gHi"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("JoinMust_CustomComparer = %v, want %v", got, want)
	}
}

func TestJoinMust_DifferentSourceTypes(t *testing.T) {
	outer := []int{5, 3, 7}
	inner := []string{"bee", "giraffe", "tiger", "badger", "ox", "cat", "dog"}
	got := JoinMust(outer, inner,
		go2linq.Identity[int],
		func(iel string) int { return len(iel) },
		func(oel int, iel string) string { return fmt.Sprintf("%d:%s", oel, iel) },
		nil,
	)
	want := []string{"5:tiger", "3:bee", "3:cat", "3:dog", "7:giraffe"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("JoinMust_DifferentSourceTypes = %v, want %v", got, want)
	}
}
