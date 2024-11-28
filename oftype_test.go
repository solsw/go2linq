package go2linq

import (
	"iter"
	"testing"

	"github.com/solsw/iterhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/OfTypeTest.cs

func TestOfType_any_int(t *testing.T) {
	type args struct {
		source iter.Seq[any]
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[int]
		wantErr bool
	}{
		{name: "UnboxToInt",
			args: args{
				source: iterhelper.VarSeq[any](10, 30, 50),
			},
			want: iterhelper.VarSeq(10, 30, 50),
		},
		{name: "OfType",
			args: args{
				source: iterhelper.VarSeq[any](1, 2, "two", 3, 3.14, 4, nil),
			},
			want: iterhelper.VarSeq(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OfType[any, int](tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("OfType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("OfType() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

func TestOfType_any_string(t *testing.T) {
	type args struct {
		source iter.Seq[any]
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[string]
		wantErr bool
	}{
		{name: "SequenceWithAllValidValues",
			args: args{
				source: iterhelper.VarSeq[any]("first", "second", "third"),
			},
			want: iterhelper.VarSeq("first", "second", "third"),
		},
		{name: "NullsAreExcluded",
			args: args{
				source: iterhelper.VarSeq[any]("first", nil, "third"),
			},
			want: iterhelper.VarSeq("first", "third"),
		},
		{name: "WrongElementTypesAreIgnored",
			args: args{
				source: iterhelper.VarSeq("first", any(1), "third"),
			},
			want: iterhelper.VarSeq("first", "third"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OfType[any, string](tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("OfType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("OfType() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}

func TestOfType_any_int64(t *testing.T) {
	type args struct {
		source iter.Seq[any]
	}
	tests := []struct {
		name    string
		args    args
		want    iter.Seq[int64]
		wantErr bool
	}{
		{name: "UnboxingWithWrongElementTypes",
			args: args{
				source: iterhelper.VarSeq[any](int64(100), 100, int64(300)),
			},
			want: iterhelper.VarSeq(int64(100), int64(300)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OfType[any, int64](tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("OfType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("OfType() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
			}
		})
	}
}
