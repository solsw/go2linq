package go2linq

import (
	"iter"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/CastTest.cs

func TestCast_any_int(t *testing.T) {
	type args struct {
		source iter.Seq[any]
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSource",
			args: args{
				source: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "UnboxToInt",
			args: args{
				source: VarAll[any](10, 30, 50),
			},
			want: VarAll(10, 30, 50),
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
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Cast() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestCast_any_string(t *testing.T) {
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
				source: VarAll[any]("first", "second", "third"),
			},
			want: VarAll("first", "second", "third"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Cast[any, string](tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Cast() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}
