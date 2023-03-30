package main


import "fmt"

type MyInt1 int
type MyInt2 = int

func main() {
	var i int = 0
	//var i1 MyInt1 = i // 不能将类型“int”赋给类型“main.MyInt1” (编译器报错)
	var i1 MyInt1 = MyInt1(i) // 可以这样写
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
	fmt.Println(i2)
}