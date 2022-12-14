package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Range generates a slice of ints within a specified range.
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
