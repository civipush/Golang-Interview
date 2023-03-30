package main

import "fmt"

const (
	x = iota
	_
	y // 2
	z = "zz"
	k
	p = iota// 5 从0开始  5-0=5
)

func main() {
	fmt.Println(x, y, z, k, p)// 0 2 zz zz 5
}