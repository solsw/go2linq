package go2linq

import (
	"iter"
	"testing"
)

func TestAppend_int(t *testing.T) {
	type args struct {
		source  iter.Seq[int]
		element int
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[int]
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
			want: VarAll(2),
		},
		{name: "1",
			args: args{
				source:  VarAll(1, 2),
				element: 3,
			},
			want: VarAll(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Append(tt.args.source, tt.args.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("Append() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Append() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Append() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}
