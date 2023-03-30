package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5}//数组
	t := a[3:5:10]// 3代表开始位置 4代表结束位置 4代表容量
	fmt.Println(t[0])// 4 t[0] = 4
	fmt.Println(t)//[4]
	fmt.Println(len(t), cap(t))// 1 1
	//截取操作符还可以有第三个参数，形如 [i,j,k]，第三个参数 k 用来限制新切片的容量，
	//但不能超过原数组（切片）的底层数组大小。截取获得的切片的长度和容量分别是：j-i、k-i
}