package main

func GetValue() int {
	return 1
}

func main() {
	//i := GetValue()
	//switch i.(type) {//可以改成switch i.(type) {case interface{}:}
	//// 不能使用switch i.(type) {case int:} 这种写法 会报错 不能使用类型断言
	//	// 但是可以使用switch i.(type) {case interface{}:}
	////	接口类型才能使用类型选择 类型选择的语法形如：i.(type)，
	//	//	其中i是接口，type是固定关键字，需要注意的是，只有接口类型才可以使用类型选择。
	//case int:
	//	fmt.Println("int")
	//case string:
	//	fmt.Println("string")
	//case interface{}:
	//	fmt.Println("interface")
	//default:
	//	fmt.Println("unknown")
	//}
}