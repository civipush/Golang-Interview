package main

import "fmt"

type person struct {
	name string
}

func main() {
	var m map[person]int
	p := person{"make"}
	fmt.Println(m[p])// 0
	//map的key是引用类型，所以这里的p是一个新的变量，不是m中的key
	//打印一个map中不存在的值时，返回元素类型的零值。
	//这个例子中，m的类型是map[person]int，因为m中 不存在p，所以打印int类型的零值，即0。
}