package main

import "fmt"

func main() {
	arrayA := [2]int{100, 200}
	var arrayB [2]int

	arrayB = arrayA

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)

	testArray(arrayA)
}

/**
三个内存地址都不同，这也就验证了 Go 中数组赋值和函数传参都是值复制的
*/
func testArray(x [2]int) {
	fmt.Printf("func Array : %p , %v\n", &x, x)
}

//1.传数组指针
func testArrayPoint(x *[]int) {
	fmt.Printf("func Array : %p , %v\n", x, *x)
	(*x)[1] += 100
}
