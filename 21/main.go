package main

import "fmt"

func hello(num ...int) {
	num[0] = 180
}

func main() {
	i := []int{5, 6, 7}
	hello(i...)
	// 可变参数args的地址跟实际外部slice的地址一样，用的同一个slice
	fmt.Println(i)// 180 可变参数传递过去，改变了第一个值。

}