//go:build go1.18

package go2linq

import (
	"testing"
)

func TestTryGetNonEnumeratedCount_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
		count  int
	}
	tests := []struct {
		name        string
		args        args
		want        bool
		wantCount   int
		wantErr     bool
		expectedErr error
	}{
		{name: "NilSourceEmitsErrNilSource",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NonCounter",
			args: args{
				source: RangeMust(2, 5),
			},
			want: false,
		},
		{name: "CounterSlice",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
			},
			want:      true,
			wantCount: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TryGetNonEnumeratedCount(tt.args.source, &tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("TryGetNonEnumeratedCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("TryGetNonEnumeratedCount() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("TryGetNonEnumeratedCount() = %v, want %v", got, tt.want)
				return
			}
			if tt.args.count != tt.wantCount {
				t.Errorf("TryGetNonEnumeratedCount().count = %v, want %v", tt.args.count, tt.wantCount)
			}
		})
	}
}

func TestTryGetNonEnumeratedCountMust_map(t *testing.T) {
	m1 := map[int]string{1: "one", 2: "two"}
	type args struct {
		source Enumerable[KeyElement[int, string]]
		count  int
	}
	tests := []struct {
		name      string
		args      args
		want      bool
		wantCount int
	}{
		{name: "CounterMap1",
			args: args{
				source: NewEnMap(m1),
			},
			want:      true,
			wantCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TryGetNonEnumeratedCountMust(tt.args.source, &tt.args.count)
			if got != tt.want {
				t.Errorf("TryGetNonEnumeratedCount() = %v, want %v", got, tt.want)
				return
			}
			if tt.args.count != tt.wantCount {
				t.Errorf("TryGetNonEnumeratedCount().count = %v, want %v", tt.args.count, tt.wantCount)
			}
		})
	}
}
