package main

import "fmt"

//赋值正确的是（）
//A. var x = nil
//B. var x interface{} = nil
//C. var x string = nil
//D. var x error = nil
func main() {
	//var x = nil// nil只能赋值给指针、chan、func、interface、map、或slice、类型的变量
	var x2 interface{} = nil
	//var x3 string = nil// 不能将类型“nil”赋给类型“string” (编译器报错)
	var x4 error = nil

	//println(x, x2, x3, x4)
	//println(x2, x4)
	fmt.Println(x2, x4)

}