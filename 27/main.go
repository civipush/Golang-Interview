package main

import "fmt"

func main() {
	//s := make(map[string]int)
	s:=make(map[string]*int)
	delete(s, "h")
	fmt.Println(s["h"])// 输出 0 如果不存在 返回0
	// 删除 map 不存在的键值对时，不会报错，相当于没有任何作用；
	//获取不存在的减值对时，返回值类型对应的零值，所以返回 0。
	// 但是如果是map[string]*int 会输出nil
}