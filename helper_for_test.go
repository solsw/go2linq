//go:build go1.18

package go2linq

type (
	elel[T any] struct {
		e1, e2 T
	}

	elelel[T any] struct {
		e1, e2, e3 T
	}

	elelelel[T any] struct {
		e1, e2, e3, e4 T
	}
)

func panickingEnumerator[T any]() Enumerator[T] {
	return OnFunc[T]{
		mvNxt: func() bool {
			panic("test panic")
		},
		crrnt: func() T {
			var t0 T
			return t0
		},
		rst: func() {},
	}
}
