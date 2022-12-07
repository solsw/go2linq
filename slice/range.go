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

// RangeMust is like Range but panics in case of error.
func RangeMust(start, count int) []int {
	r, err := Range(start, count)
	if err != nil {
		panic(err)
	}
	return r
}
