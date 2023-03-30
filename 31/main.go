package main

import "fmt"

//定义一个包内全局字符串变量，下面语法正确的是（）
//A. var str string
//B. str := ""
//C. str = ""
//D. var str = ""
//答：A、D

var str0 string
//str1 := ""
//str2 = ""//编译错误
var str3 = ""
func main() {

	//fmt.Println(str0, str1, str2, str3)
	fmt.Println(str0, str3)

}