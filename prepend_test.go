package go2linq

import (
	"errors"
	"iter"
	"testing"
)

func TestPrepend_string(t *testing.T) {
	type args struct {
		source  iter.Seq[string]
		element string
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[string]
		wantErr     bool
		expectedErr error
	}{
		{name: "01",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "02",
			args: args{
				element: "two",
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "EmptySource",
			args: args{
				source:  Empty[string](),
				element: "two",
			},
			want: VarToSeq("two"),
		},
		{name: "1",
			args: args{
				source:  VarToSeq("one", "two"),
				element: "zero",
			},
			want: VarToSeq("zero", "one", "two"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Prepend(tt.args.source, tt.args.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("Prepend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("Prepend() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Prepend() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}
