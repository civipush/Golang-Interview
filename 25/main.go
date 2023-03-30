package main
//cap() 函数的适用类型，下面说法正确的是()
//A. array
//B. slice
//C. map
//D. channel
func main() {
	//map的长度是不固定的，因此不能用cap()求长度
	var a [10]int
	var b []int
	//var c map[string]int
	var d chan int
	//println(cap(a), cap(b), cap(c), cap(d))// 10 0 0 0
	println(cap(a), cap(b), cap(d))// 10 0 0
}