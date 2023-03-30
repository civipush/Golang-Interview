package main

import "fmt"

func main() {
	str := "hello"
	//str[0] = 'x'//编译错误 不能修改字符串 不能修改字符串的某个字符
	//如果想修改字符串的某个字符 可以先将字符串转换为[]byte类型
	strByte := []byte(str)
	strByte[0] = 'x'
	str = string(strByte)
	fmt.Println(str)
}