package main

import "fmt"

var (
	//size := 1024 // 不能在这里使用 := 赋值 。:= 只能用在函数内部。
	size   = 1024
	max_size = size * 2
)

func main() {
	fmt.Println(size, max_size)
}