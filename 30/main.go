package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

//func (t *Teacher) ShowB() {
//	fmt.Println("teacher showB")
//}
//注释了上面的方法，就会输出 showB
func main() {
	t := Teacher{}
	t.ShowB()
	// 输出 teacher showB
	//结构体嵌套。嵌套结构体时，内部类型的属性、方法，可以为外部类型所有，就好像是外部类型自己的一样。
	//在嵌套结构体中，People 称为内部类型，Teacher 称为外部类型；
	//此外，外部类型还可以定义自己的属性和方法，甚至可以定义与内部相同的方法，这样内部类型的方法就会被“屏蔽”。
	//这个例子中的 ShowB() 就是同名方法。
}