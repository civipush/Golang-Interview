package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p :=1
	fmt.Println(&p)//0xc00000a0b8
	incr(&p)
	fmt.Println(p)//2 传递的是指针 会改变原来的值
	fmt.Println(&p)//0xc00000a0b8
	// 但是不会改变原来的指针 也就是说指针的值是不会改变的 但是指针指向的值是会改变的
}