package main

import "fmt"

//channel，下面语法正确的是（）
//A. var ch chan int
//B. ch := make(chan int)
//C. <-ch
//D. ch<-

func main() {
	var ch chan int
	ch1 := make(chan int)
	<-ch1
	//ch1<-// 语法错误: unexpected semicolon or newline before statement

	fmt.Println(ch, ch1)
	//A、B都是申明channel；C读取channel；
	//写channel是必须带上值，所以D错误
}