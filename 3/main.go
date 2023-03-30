package main

import "fmt"

// 1.
func main1() {
	s := make([]int, 5)// 生成一个长度为 5 的切片
	s = append(s, 1, 2, 3)
	fmt.Println(s) // [0 0 0 0 0 1 2 3]
}

// 2.
func main() {
	s := make([]int, 0)// 生成一个空的切片的 长度和容量都是 0
	// nil slice 的长度和容量都是 0 性质是 nil
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s) // [1 2 3 4] // 为什么不是 [0 0 0 0 1 2 3 4] ？
	// 因为 make([]int, 0) 生成的是一个空的切片，而不是一个长度为 0 的切片。长度为0的切片是 nil slice。
}