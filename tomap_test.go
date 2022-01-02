//go:build go1.18

package go2linq

import (
	"reflect"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ToDictionaryTest.cs

func Test_ToMap_string_rune(t *testing.T) {
	type args struct {
		source      Enumerator[string]
		keySelector func(string) rune
	}
	tests := []struct {
		name        string
		args        args
		want        map[rune]string
		wantErr     bool
		expectedErr error
	}{
		{name: "NilKeySelectorNoComparerNoElementSelector",
			args: args{
				source: Empty[string](),
			},
			wantErr:     true,
			expectedErr: ErrNilSelector,
		},
		{name: "JustKeySelector",
			args: args{
				source:      NewOnSlice("zero", "one", "two"),
				keySelector: func(s string) rune { return []rune(s)[0] },
			},
			want: map[rune]string{'z': "zero", 'o': "one", 't': "two"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToMap(tt.args.source, tt.args.keySelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("ToMap() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", String(NewOnMapImmediate(got)), String(NewOnMapImmediate(tt.want)))
			}
		})
	}
}

func Test_ToMap_string_string(t *testing.T) {
	type args struct {
		source      Enumerator[string]
		keySelector func(string) string
	}
	tests := []struct {
		name        string
		args        args
		want        map[string]string
		wantErr     bool
		expectedErr error
	}{
		{name: "DuplicateKey",
			args: args{
				source:      NewOnSlice("zero", "One", "Two", "three"),
				keySelector: func(s string) string { return strings.ToLower(string([]rune(s)[:1])) },
			},
			wantErr:     true,
			expectedErr: ErrDuplicateKeys,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToMap(tt.args.source, tt.args.keySelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("ToMap() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", String(NewOnMapImmediate(got)), String(NewOnMapImmediate(tt.want)))
			}
		})
	}
}

func Test_ToMapSel_string_rune_int(t *testing.T) {
	type args struct {
		source          Enumerator[string]
		keySelector     func(string) rune
		elementSelector func(string) int
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{name: "KeyAndElementSelector",
			args: args{
				source:          NewOnSlice("zero", "one", "two"),
				keySelector:     func(s string) rune { return []rune(s)[0] },
				elementSelector: func(s string) int { return len(s) },
			},
			want: map[rune]int{'z': 4, 'o': 3, 't': 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ToMapSel(tt.args.source, tt.args.keySelector, tt.args.elementSelector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapSel() = %v, want %v", String(NewOnMapImmediate(got)), String(NewOnMapImmediate(tt.want)))
			}
		})
	}
}

func Test_CustomSelector_string_string_int(t *testing.T) {
	source := NewOnSliceEn("zero", "one", "THREE")
	keySelector := func(s string) string { return strings.ToLower(string([]rune(s)[0])) }
	elementSelector := func(s string) int { return len(s) }
	got, _ := ToMapSel(source, keySelector, elementSelector)
	if len(got) != 3 {
		t.Errorf("len(ToMapSel()) = %v, want 3", len(got))
	}
	want := map[string]int{"z": 4, "o": 3, "t": 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToMapSel() = '%v', want '%v'", String(NewOnMapImmediate(got)), String(NewOnMapImmediate(want)))
	}
}
