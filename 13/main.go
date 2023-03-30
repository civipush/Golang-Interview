package main

import "fmt"

//字符串连接，下面语法正确的是？
//A. str := 'abc' + '123'
//B. str := "abc" + "123"
//C. str := '123' + "abc"
//D. fmt.Sprintf("abc%d", 123)

func main() {
	//str0 := 'abc' + '123'// 不能将类型“int32”赋给类型“int32” (编译器报错)
	//fmt.Println(str0)
	str := "abc" + "123"
	fmt.Println(str)
	//str2 := '123' + "abc"// 不能将类型“string”赋给类型“int32” (编译器报错)
	str3 := fmt.Sprintf("abc%d", 123)
	fmt.Println(str3)
}