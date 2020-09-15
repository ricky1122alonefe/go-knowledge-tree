package main

import (
	"fmt"
)

func Assign1(s []int) {
	fmt.Printf("接受参数的地址为 addr %p\n", &s)
	s = []int{6, 6, 6}
}

func Reverse0(s [5]int) {
	fmt.Printf("Reverse0-接受参数的地址为 addr %p\n", &s)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse1(s []int) {
	fmt.Printf("Reverse1-接受参数的地址为 addr %p\n", &s)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse2(s []int) {
	s = append(s, 999)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse3(s []int) {
	s = append(s, 999, 1000, 1001)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("传入参数地址为 addr %p\n", &s)
	Assign1(s)
	fmt.Println(s)
	// (1) 输出[1, 2, 3, 4, 5, 6]
	// 因为是值拷贝传递，Assign1里的s和main里的s是不同的两个指针

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Reverse0-传入参数的地址为 addr %p\n", &array)
	Reverse0(array)
	fmt.Println(array)
	// (2) 输出[1, 2, 3, 4, 5]
	// 传递时对array进行了一次值拷贝，不会影响原来的array

	s = []int{1, 2, 3}
	fmt.Printf("Reverse1-传入参数的地址为 addr %p\n", &s)
	Reverse1(s)
	fmt.Println("---------------------------------")
	fmt.Println(s)
	fmt.Println("---------------------------------")
	// (3) 输出[1, 2, 3]
	// 在没有对s进行append时，len(s)=3，cap(s)=3
	// append之后超过了容量，返回了一个新的slice
	// 相当于只改变了新的slice，旧的slice没影响

	var a []int
	for i := 1; i <= 3; i++ {
		a = append(a, i)
	}
	Reverse2(a)
	fmt.Println(a)
	// (4) 输出[999, 3, 2]
	// 在没有对a进行append时，len(a)=3，cap(a)=4
	// append后没有超过容量，所以元素直接加在了数组上
	// 虽然函数Reverse2里将a的len加1了，但它只是一个值拷贝
	// 不会影响main里的a，所以main里的len(a)=3

	var b []int
	for i := 1; i <= 3; i++ {
		b = append(b, i)
	}
	Reverse3(b)
	fmt.Println(b)
	// (5) 输出[1, 2, 3]
	// 原理同(3)

	c := [3]int{1, 2, 3}
	d := c
	c[0] = 999
	fmt.Println(d)
	// (6) 输出[1, 2, 3]
	// 数组赋值是值拷贝，所以不会影响原来的数组
}
