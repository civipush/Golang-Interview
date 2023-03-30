package main

import "fmt"

func hello(i int) {
	fmt.Println(i)
}
func main() {
	i := 5
	defer hello(i)// 输出 5 defer延迟执行 会将参数复制一份 传递给函数 但是不会立即执行
	// 会在函数结束时执行 但是会在return之前执行
	i = i + 10
}