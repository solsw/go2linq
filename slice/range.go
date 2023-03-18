package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Range generates a slice of [int]s within a specified range.
//
// [int]: https://pkg.go.dev/builtin#int
func Range(start, count int) ([]int, error) {
	if count < 0 {
		return nil, go2linq.ErrNegativeCount
	}
	r := make([]int, count)
	for i := 0; i < count; i++ {
		r[i] = start
		start++
	}
	return r, nil
}
