//go:build go1.18

package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Range generates a slice of ints within a specified range.
func Range(start, count int) ([]int, error) {
	en, err := go2linq.Range(start, count)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// RangeMust is like Range but panics in case of error.
func RangeMust(start, count int) []int {
	r, err := Range(start, count)
	if err != nil {
		panic(err)
	}
	return r
}
