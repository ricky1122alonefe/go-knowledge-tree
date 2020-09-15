package main

import (
	"fmt"
	"sync"
	"time"
)

var s1 []int
var lock sync.Mutex //互斥锁
func appendValues(i int) {
	lock.Lock() //加锁
	s1 = append(s1, i)
	lock.Unlock() //解锁
}

func main() {
	for i := 0; i < 10000; i++ {
		go appendValues(i)
	}
	//sort.Ints(s) //给切片排序,先排完序再打印,和下面一句效果相同
	time.Sleep(time.Second) //间隔1s再打印,防止一边插入数据一边打印时数据乱序
	for i, v := range s1 {
		fmt.Println(i, ":", v)
	}
}
