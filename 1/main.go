package main

import "fmt"

func main() {
	defer_call()
}

func defer_call()  {
	//defer 语句会将函数推迟到外层函数返回之后执行。
	//推迟调用的参数会立刻生成，但是在外层函数返回前函数都不会被调用。
	//调用顺序为后进先出
	defer func() {fmt.Println("打印前")}()
	defer func() {fmt.Println("打印中")}()
	defer func() {fmt.Println("打印后")}()

	panic("触发异常")
}