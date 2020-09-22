package main

import (
	"fmt"
	"time"
)

func main() {
	in := producer(1, 2, 3, 4, 5, 6, 7, 8)
	var process int
	var t time.Ticker
	go func() {
		for {
			select {
			case x, ok := <-in:
				if !ok {
					return
				}
				fmt.Println(x)
				process++
			case <-t.C:
				fmt.Printf("Working, processedCnt = %d\n", process)
			}
		}
	}()

	time.Sleep(time.Second * 10)
}
func producer(num ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range num {
			out <- v
		}
	}()

	return out
}
