package main

import "fmt"
type person struct {
	name string
	age  int
	sex string
}
func main() {
	//通过指针变量p访问其成员变量name,有哪几种方式？
	//var p *person// 定义一个指针变量
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4a2c5a]
	// 原因：p是一个指针变量，它的值是nil，不能直接使用p.name访问成员变量

	p := &person{} // 定义一个指针变量并初始化 值为 &person{} 可以使用 p.name 访问成员变量
	p.name = "Tom"

	fmt.Println(p.name)//
	//fmt.Println((&p).name )// 不能使用 (&p).name 这种方式访问成员变量
	fmt.Println((*p).name )// 可以使用 (*p).name 这种方式访问成员变量
	//fmt.Println(p->name )// 可以使用 p->name 这种方式访问成员变量
	//A. p.name
	//B. (&p).name
	//C. (*p).name
	//D. p->name
}