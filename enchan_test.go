package go2linq

import (
	"fmt"
)

func chn3() chan int {
	var ch = make(chan int)
	go func() {
		ch <- 4
		ch <- 3
		ch <- 2
		ch <- 1
		ch <- 0
		close(ch)
	}()
	return ch
}

func ExampleEnChan() {
	en1 := NewEnChanEn[int](chn3())
	en2 := SelectMust[int](en1, func(i int) int { return 12 / i })

	// panic: runtime error: integer divide by zero
	// fmt.Println(LastMust[int](en2))

	fmt.Println(
		FirstMust[int](en2),
	)
	fmt.Println(
		FirstMust[int](
			SkipMust[int](en2, 2),
		),
	)
	// Output:
	// 3
	// 12
}
