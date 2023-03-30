package main

import "fmt"

func main() {
	a := 5
	b := 8.1
	//fmt.Println(a + b)
	// 报错 不能将float64类型的值赋值给int类型的变量
	fmt.Println(a + int(b))
}