package main

import (
	"fmt"
	"os"
)

func main() {
	//打开8/main.go 9/main.go...一直到194/main.go 并写入"package main"内容
	for i := 10; i < 195; i++ {
		f, err := os.Create(fmt.Sprintf("%d/main.go", i))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString("package main")
		//把光标移动到package main结尾 也就是文件结尾 seek(0,2) 0表示偏移量 2表示从文件结尾开始偏移 1
		f.Seek(0, 1)

	}

	//f, err := os.OpenFile("main.go", os.O_WRONLY, 0666)

	//f, err := os.Create("main.go")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//f.WriteString("package main")

}
