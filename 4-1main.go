package main

import (
	"fmt"
	"unsafe"
)

type K struct {
}

type F struct {
	num1 K
	num2 int32
}

func main() {

	//变量的长度
	i := int(1)
	//指针
	p := &i
	fmt.Println(i)
	fmt.Println(*p)
	//打印变量的长度
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(unsafe.Sizeof(int(0)))

	//结构体长度
	a := K{}
	b := int(0)
	c := K{}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)

	//
	f := F{}
	k := K{}
	fmt.Printf("%p\n", &f.num1)
	fmt.Printf("%p\n", &k)

	//map的key是string value是空结构体struct{}
	//想要hashset，go没有的时候就可以用空结构体
	m := map[string]struct{}{}
	m["a"] = struct{}{}
	m["b"] = struct{}{}

	//channel用来传输数据的，得告诉channel传什么数据，channel只是负责传递一个结构体，空结构体不占用任何内存
	//a:make(chan struct{})

}
