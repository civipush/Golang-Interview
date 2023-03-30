package main

import "C"
import "fmt"

func add(args ...int) int {

	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}
func main() {
	//A. add(1, 2)
	//B. add(1, 3, 7)
	//C. add([]int{1, 2})// 编译错误 传递的是切片
	//D. add([]int{1, 3, 7}…)// 可变参数 传递切片时需要加上...
	fmt.Println(add(1, 2, 3, 4, 5))

}