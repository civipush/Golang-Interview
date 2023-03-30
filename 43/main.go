package main

import "fmt"

//x 已声明，y 没有声明，判断每条语句的对错。
//1. x, _ := f()
//2. x, _ = f()
//3. x, y := f()
//4. x, y = f()

func main() {
	var x int
	//x, _ := f()
	x, _ = f()
	x, y := f()
	x, y = f()
	fmt.Println(x, y)
	//1.错，x 已经声明，不能使用 :=；
	//2.对；
	//3.对，当多值赋值时，:= 左边的变量无论声明与否都可以；
	//4.错，y 没有声明。
}

func f() (int, interface{}) {

	return 1, 2
}