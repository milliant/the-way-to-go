package GoroutinesAndChannels

import (
	"fmt"
)

type A int

func Sum(a, b A, ch chan A) {
	ch <- (a + b)
}

func (_ A) Main() {
	var a, b A
	a = A(3)
	b = A(5)
	ch := make(chan A)
	go Sum(a, b, ch)
	var sum A
	sum = <-ch
	fmt.Printf("sum of 2 nums is %d", sum)
}
