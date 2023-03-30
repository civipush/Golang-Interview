package main

import (
	"fmt"
	"reflect"
)

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "11"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "2"}}

	//if sm1 == sm2 {// 结构体中的 map 不能比较 如果要比较，需要使用 reflect.DeepEqual() 函数
	//
	//	fmt.Println("sm1 == sm2")
	//}
	//结构体只能比较是否相等，但是不能比较大小；
	//想同类型的结构体才能进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关；
	//如果struct的所有成员都可以比较，则该struct就可以通过==或!=进行比较是否相同，
	//比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等；
	//那有什么是可以比较的呢？答案是可以比较的类型有：bool、数值型、字符、指针、数组等，
	//但是slice、map、函数是不能比较的。

	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	}else{
		fmt.Println("sm1 != sm2")
	}
}