package main

import "fmt"

//填入哪个选项以输出yes nil？
func main() {
	//var s1 []int
	//nil 切片和 nil 相等，一般用来表示一个不存在的切片；
	//空切片和 nil 不相等，表示一个空的集合。
	var s2 = []int{}
	if s2 == nil {
		fmt.Println("yes nil")
	}else{
		fmt.Println("no nil")
	}
}
//A. s1 //s1是nil
//B. s2
//C. s1、s2 都可以