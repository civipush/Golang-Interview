package main

import "fmt"

func hello() []string {
	return nil
}

func main() {
	//h := hello
	h:=hello()
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	// Output: not nil // 为什么不是nil呢？
	// 将 hello() 赋值给变量h，而不是函数的返回值，所以输出 not nil
	// 因为函数是一等公民，函数也是一种类型，函数类型的零值是nil
}