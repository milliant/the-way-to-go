package GoroutinesAndChannels

import (
	"fmt"
)

type A int // A类型是（int）基础类型的别名  type alias basicType

func Sum(a, b A, ch chan A) {
	ch <- (a + b)
}

func (_ A) Main() {
	var a A
	a = A(3) //basicType和alias类型可以相互进行类型转换
	b := int(a)

	ch := make(chan A)
	go Sum(a, A(b), ch)
	var sum A
	sum = <-ch
	fmt.Printf("sum of 2 nums is %d", sum)
}
