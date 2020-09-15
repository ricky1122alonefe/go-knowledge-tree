package main

func main() {

	a := []int{0, 1, 2, 3, 4}

	//删除第i个元素
	i := 2
	a = append(a[:i], a[i+1:]...)

}

// 通过索引删除
func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
