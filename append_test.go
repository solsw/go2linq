//go:build go1.18

package go2linq

import (
	"testing"
)

func TestAppend_int(t *testing.T) {
	type args struct {
		source  Enumerator[int]
		element int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerator[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "01",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "02",
			args: args{
				element: 2,
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "EmptySource",
			args: args{
				source:  Empty[int](),
				element: 2,
			},
			want: NewOnSlice(2),
		},
		{name: "1",
			args: args{
				source:  NewOnSlice(1, 2),
				element: 3,
			},
			want: NewOnSlice(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Append(tt.args.source, tt.args.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("Append() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Append() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Append() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
