package go2linq

import (
	"errors"
)

// type Error struct {
// 	error
// }

var (
	ErrDuplicateKeys    = errors.New("duplicate keys")
	ErrEmptySource      = errors.New("empty source")
	ErrIndexOutOfRange  = errors.New("index out of range")
	ErrMultipleElements = errors.New("multiple elements")
	ErrMultipleMatch    = errors.New("multiple match")
	ErrNegativeCount    = errors.New("negative count")
	ErrNilAccumulator   = errors.New("nil accumulator")
	ErrNilAction        = errors.New("nil action")
	ErrNilComparer      = errors.New("nil comparer")
	ErrNilLesser        = errors.New("nil lesser")
	ErrNilPredicate     = errors.New("nil predicate")
	ErrNilSelector      = errors.New("nil selector")
	ErrNilSource        = errors.New("nil source")
	ErrNoMatch          = errors.New("no match")
	ErrSizeOutOfRange   = errors.New("size out of range")
)
