package main

import "fmt"

var s []int

func appendValue(i int) {
	s = append(s, i)
}

func main() {
	for i := 0; i < 10000; i++ { //10000个协程同时添加切片
		go appendValue(i)
	}

	for i, v := range s { //同时打印索引和值
		fmt.Println(i, ":", v)
	}
}
