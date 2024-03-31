package go2linq

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/solsw/runtimehelper"
)

var (
	ErrDuplicateKeys    = errors.New("duplicate keys")
	ErrEmptySource      = errors.New("empty source")
	ErrIndexOutOfRange  = errors.New("index out of range")
	ErrMultipleElements = errors.New("multiple elements")
	ErrMultipleMatch    = errors.New("multiple match")
	ErrNegativeCount    = errors.New("negative count")
	ErrNilAccumulator   = errors.New("nil accumulator")
	ErrNilAction        = errors.New("nil action")
	ErrNilCompare       = errors.New("nil compare")
	ErrNilEqual         = errors.New("nil equal")
	ErrNilLess          = errors.New("nil less")
	ErrNilPredicate     = errors.New("nil predicate")
	ErrNilSelector      = errors.New("nil selector")
	ErrNilSource        = errors.New("nil source")
	ErrNoMatch          = errors.New("no match")
	ErrSizeOutOfRange   = errors.New("size out of range")
)

func callerError(err error) error {
	s1 := path.Base(runtimehelper.NthCallerName(2))
	if s1 == "" {
		return err
	}
	s2, _, _ := strings.Cut(s1, "[")
	ss3 := strings.Split(s2, ".")
	return fmt.Errorf("%s: %w", ss3[len(ss3)-1], err)
}
