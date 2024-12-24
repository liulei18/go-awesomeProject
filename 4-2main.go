package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	//16
	fmt.Println(unsafe.Sizeof("慕课网"))
	fmt.Println(unsafe.Sizeof("慕课网Moody老师"))

	s := "慕课网"
	//先转成pointer万能指针，在转成sh指针
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	//9
	fmt.Println(sh.Len)

	s1 := "慕课网moody"
	//先转成pointer万能指针，在转成sh指针
	sh1 := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	//14
	fmt.Println(sh1.Len)

	//字符串是变长编码，角标方式访问底层字节数组，是无法打印字
	s2 := "慕课网moody"
	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i])
	}

	//utf8.go decode、encode方法，多个字节解码成一个字符，1个字符编码成多个字节，rune在go中是utf8字符的意思，底层是1个int32
	for _, char := range s2 {
		fmt.Println(char)
		fmt.Printf("%c\n", char)
	}
}
