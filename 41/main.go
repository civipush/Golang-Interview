package main

import "fmt"

//A B 两处应该怎么修改才能顺利编译？
func main() {
	//var m map[string]int        //A
	m := make(map[string]int)
	m["a"] = 1
	if v ,k:= m["b"]; v!= 0 {    //B
		fmt.Println(v,k)
	}
}