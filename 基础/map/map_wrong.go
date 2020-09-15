package main

import "fmt"

// 读写并不安全
func main() {
	m := make(map[int]int)
	go func() { //开一个协程写map
		for i := 0; i < 10000; i++ {
			m[i] = i
		}
	}()

	go func() { //开一个协程读map
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
		}
	}()

	//time.Sleep(time.Second * 20)
	for {

	}
}
