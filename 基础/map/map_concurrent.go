package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func main() {
	m := make(map[int]int)
	go func() { //开一个协程写map
		for i := 0; i < 10000; i++ {
			lock.Lock() //加锁
			m[i] = i
			lock.Unlock() //解锁
		}
	}()
	go func() { //开一个协程读map
		for i := 0; i < 10000; i++ {
			lock.Lock() //加锁
			fmt.Println(m[i])
			lock.Unlock() //解锁
		}
	}()
	time.Sleep(time.Second * 20)
}
