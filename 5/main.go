package main

import "fmt"

//new() 与 make() 的区别
//new(T)  和 make(T, args)  是Go语言内建函数，用来
//分配内存，但适用的类型不用。
//new(T) 会为了 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T 的值。
//换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如 数组 、 结构体 等。
//make(T, args) 返回初始化之后的T类型的值，也不是指针 *T ，是经过初始化之后的T的引用。
//make() 只适用于 slice 、 map 和 channel 。

func main() {
	//使用new()函数 创建数组
	var arr1 = new([5]int)
	//使用make()函数 创建数组
	var arr2 = make([]int, 5)
	//使用new()函数 创建map
	var m1 = new(map[string]int)
	//使用make()函数 创建map
	var m2 = make(map[string]int)
	//使用new()函数 创建channel
	var c1 = new(chan int)
	//使用make()函数 创建channel
	var c2 = make(chan int)


	fmt.Println(arr1, arr2, m1, m2, c1, c2)
}