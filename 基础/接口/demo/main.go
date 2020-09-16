package main

import "fmt"

func main() {
	var a Bird
	a.Sing()
}

type IFace interface {
	Sing()
}

type Bird struct{}

func (receiver *Bird) Sing() {
	fmt.Println("hello sing")
}
