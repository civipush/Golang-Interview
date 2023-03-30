package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, _ := range slice {
		//m[key] = &val
		m[key] = &slice[key]
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
		// 0 -> 3
		// 1 -> 3
		// 2 -> 3
		// 3 -> 3 的原因是什么？
		// 因为 val 是 slice 的一个元素，它的地址不会改变，所以最后 val 的值是 3，
		// 而 &val 取的是 val 的地址，所以最后 map 中的所有元素的值都是 &val，也就是 3 的地址。
		// 改进的方法是：
		// m[key] = &slice[key]

	}
}