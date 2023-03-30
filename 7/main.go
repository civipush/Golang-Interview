package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	//s1 = append(s1, s2)
	// 不能使用 new() 生成的指针来调用 append() 函数 。
	// 应当使用 make() 生成的切片来调用 append() 函数。
	s1 =append(s1,s2...)//...表示将s2中的元素一个一个的添加到s1中
	fmt.Println(s1)
}