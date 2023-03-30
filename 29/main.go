package main

import "fmt"

func main() {
	i := -5
	j := +5
	fmt.Printf("%+d %+d", i, j)// 输出 -5 +5 正数前面加+号
	// %d表示输出十进制数字，
	// +表示输出数值的符号。这里不表示取反。
}