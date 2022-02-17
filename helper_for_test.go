//go:build go1.18

package go2linq

type (
	Pet struct {
		Name string
		Age  int
	}

	Person struct {
		LastName string
		Pets     []Pet
	}

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
