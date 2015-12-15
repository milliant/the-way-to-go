package 14GoroutinesAndChannels
import(
"fmt"
)
type A int
func Sum(a,b int,ch chan int){
	ch <-(a+b)
}

func(a A)Main(){
	a,b := 3,5
	chan:=make(chan int)
	go Sum(a,b,chan)
	sum<-ch
	fmt.Printf("sum of 2 nums is %d",sum)
}

