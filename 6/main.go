package main

import "fmt"

func main() {
	//list := new([]int)
	list := make([]int, 0)
	list = append(list, 1)// 不能使用 new() 生成的指针来调用 append() 函数 。
	// 应当使用 make() 生成的切片来调用 append() 函数。
	// map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new 来初始化。

	fmt.Println(list)
}