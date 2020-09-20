package main

import "log"

func main() {
	a := make(chan int)
	log.Println(a == nil)
	close(a)
}
