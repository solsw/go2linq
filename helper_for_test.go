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
	return enrFunc[T]{
		mvNxt: func() bool {
			panic("test panic")
		},
		crrnt: func() T {
			return ZeroValue[T]()
		},
		rst: func() {},
	}
}
