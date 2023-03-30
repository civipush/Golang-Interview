package main

import "fmt"

func main() {
	a := []int{7, 8, 9}
	fmt.Printf("%+v\n", a)//[7 8 9]
	ap(a)
	fmt.Printf("%+v\n", a)//[7 8 9]
	app(a)
	fmt.Printf("%+v\n", a)//[1 8 9] 说明a是引用传递
}

func ap(a []int) {
	a = append(a, 10) // 这里的 a 是一个新的切片，不会影响到原来的切片
	//append导致底层数组重新分配内存 会导致原来的切片指向的底层数组被垃圾回收
}

func app(a []int) {
	a[0] = 1
}