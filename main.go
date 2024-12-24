package main

import "fmt"

func main() {
	//字面量创建切片，执行go build -gcflags -S main.go，查看汇编代码
	s := []int{1, 2, 3}
	fmt.Println(s)

	append(s, 4, 5)

}
