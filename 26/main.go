package main

import "fmt"

func main() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")// 输出 nil
		//当且仅当 i 的动态类型和动态值都为 nil 时，i 才等于 nil。
		return
	}
	fmt.Println("not nil")
}