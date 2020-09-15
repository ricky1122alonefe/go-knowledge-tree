package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var golang_array [5]string //定义一个长度为5的字符串数组。我们可以用一个定义数组的公式：“var golang_array [n]type ”, n表示数组长度，type表示数组存储类型。
	line := "yinzhengjie is a good boy"
	list := strings.Split(line, " ") //将字符串切割成相应的一个切片。
	for k, v := range list {
		golang_array[k] = v //将切片的每一个值赋值给我们刚刚定义的长度为5的数组。
	}
	fmt.Println(reflect.TypeOf(list)) //查看list的数据类型
	fmt.Println(golang_array)
	fmt.Println(reflect.TypeOf(golang_array))
}
