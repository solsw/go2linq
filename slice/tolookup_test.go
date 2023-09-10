package slice

import (
	"reflect"
	"testing"

	"github.com/solsw/collate"
	"github.com/solsw/go2linq/v3"
)

func TestToLookup_string_int(t *testing.T) {
	lk := &go2linq.Lookup[int, string]{KeyEq: collate.DeepEqualer[int]{}}
	lk.Add(3, "abc")
	lk.Add(3, "def")
	lk.Add(1, "x")
	lk.Add(1, "y")
	lk.Add(3, "ghi")
	lk.Add(1, "z")
	lk.Add(2, "00")
	type args struct {
		source      []string
		keySelector func(string) int
		equaler     collate.Equaler[int]
	}
	tests := []struct {
		name    string
		args    args
		want    *go2linq.Lookup[int, string]
		wantErr bool
	}{
		{name: "EmptySource",
			args: args{
				source:      []string{},
				keySelector: func(s string) int { return len(s) },
			},
			want: &go2linq.Lookup[int, string]{},
		},
		{name: "LookupWithNoComparerOrElementSelector",
			args: args{
				source:      []string{"abc", "def", "x", "y", "ghi", "z", "00"},
				keySelector: func(s string) int { return len(s) },
			},
			want: lk,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToLookup(tt.args.source, tt.args.keySelector, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.EqualTo(tt.want) {
				t.Errorf("ToLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLookup(t *testing.T) {
	lk1 := &go2linq.Lookup[string, string]{KeyEq: collate.DeepEqualer[string]{}}
	lk1.Add("abc", "abc")
	lk1.Add("def", "def")
	lk1.Add("ABC", "ABC")
	lk2 := &go2linq.Lookup[string, string]{KeyEq: collate.DeepEqualer[string]{}}
	lk2.Add("abc", "abc")
	lk2.Add("def", "def")
	lk2.Add("abc", "ABC")
	type args struct {
		source      []string
		keySelector func(string) string
		equaler     collate.Equaler[string]
	}
	tests := []struct {
		name    string
		args    args
		want    *go2linq.Lookup[string, string]
		wantErr bool
	}{
		{name: "LookupWithNilComparerButNoElementSelector",
			args: args{
				source:      []string{"abc", "def", "ABC"},
				keySelector: go2linq.Identity[string],
			},
			want: lk1,
		},
		{name: "LookupWithComparerButNoElementSelector",
			args: args{
				source:      []string{"abc", "def", "ABC"},
				keySelector: go2linq.Identity[string],
				equaler:     collate.CaseInsensitiveOrder,
			},
			want: lk2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToLookup(tt.args.source, tt.args.keySelector, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.EqualTo(tt.want) {
				t.Errorf("ToLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLookupSel(t *testing.T) {
	lk := &go2linq.Lookup[int, string]{KeyEq: collate.DeepEqualer[int]{}}
	lk.Add(3, "a")
	lk.Add(3, "d")
	lk.Add(1, "x")
	lk.Add(1, "y")
	lk.Add(3, "g")
	lk.Add(1, "z")
	lk.Add(2, "0")
	type args struct {
		source          []string
		keySelector     func(string) int
		elementSelector func(string) string
		equaler         collate.Equaler[int]
	}
	tests := []struct {
		name    string
		args    args
		want    *go2linq.Lookup[int, string]
		wantErr bool
	}{
		{name: "LookupWithElementSelectorButNoComparer",
			args: args{
				source:          []string{"abc", "def", "x", "y", "ghi", "z", "00"},
				keySelector:     func(s string) int { return len(s) },
				elementSelector: func(s string) string { return string(s[0]) },
			},
			want: lk,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToLookupSel(tt.args.source, tt.args.keySelector, tt.args.elementSelector, tt.args.equaler)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToLookupSel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToLookupSel() = %v, want %v", got, tt.want)
			}
		})
	}
}
