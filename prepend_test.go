//go:build go1.18

package go2linq

import (
	"testing"
)

func TestPrepend_string(t *testing.T) {
	type args struct {
		source  Enumerable[string]
		element string
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[string]
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
			want: NewEnSlice("two"),
		},
		{name: "1",
			args: args{
				source:  NewEnSlice("one", "two"),
				element: "zero",
			},
			want: NewEnSlice("zero", "one", "two"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Prepend(tt.args.source, tt.args.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("Prepend() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Prepend() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Prepend() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}
