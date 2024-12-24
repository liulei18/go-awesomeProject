// 告诉编译器程序要编译成可执行文件，而不是共享库
package main

import (
	"fmt"
	"math/cmplx"
)
import "reflect"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	/*
		//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
		// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
		s := "gopher"
		fmt.Println("Hello and welcome, %s!", s)

		for i := 1; i <= 5; i++ {
			//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
			// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
			fmt.Println("i =", 100/i)
		}*/

	fmt.Println("Hello World")

	//1.变量 Go是一种静态类型的语言，编码时确定变量类型
	var a int
	fmt.Println(a)

	message := "hello world"

	var b, c int = 1, 2

	fmt.Println(b, c, message)

	//2.基本数据类型
	var a1 bool = true
	var b1 int = 1
	var c1 string = "hello world"
	d := 1.22
	var x complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Println(a1, b1, c1, reflect.TypeOf(d))
	fmt.Printf("%T", a1)
	fmt.Println(x)

	/**
	3.数组在生命中定义擦好高难度，不能扩展
	*/
	var a2 [5]int
	fmt.Println(a2)
	var mutiD [2][3]int
	fmt.Println(mutiD)

	/**
	数组的值在运行时不能更改。
	数组也不能直接获取子数组，所有有了切片
	4.切片，可以随时扩展。切片是数组的抽象。切片使用数组作为底层的结构。
	切片包含：长度、容量和指向底层数组的指针

	*/
	var b3 []int
	fmt.Println(b3)

	numbers := make([]int, 5, 10)
	fmt.Println(numbers)

	//append、copy函数可以增加切片容量
	numbers = append(numbers, 1, 2, 3, 4)
	fmt.Println(numbers)

	numbers2 := make([]int, 15)
	//将原切片复制到新切片
	copy(numbers2, numbers)
	fmt.Println(numbers2)

	//初始化长度为4，以及赋值
	numbers3 := []int{1, 2, 3, 4}
	fmt.Println(numbers3)
	//切片创建子切片
	slice1 := numbers3[2:]
	fmt.Println(slice1)
	slice2 := numbers3[:3]
	fmt.Println(slice2)
	slice3 := numbers3[1:4]
	fmt.Println(slice3)

	/**
	5.map

	*/
	//key是string，value时int的map类型变量
	m := make(map[string]int)
	m["clearity"] = 2
	m["simplicity"] = 3
	//printing the values
	fmt.Println(m["clearity"])
	fmt.Println(m["simplicity"])

	/**
	6.类型转化
	*/
	a5 := 1.9
	b5 := int(a5)
	fmt.Println(b5)

	/**
	7.流程控制
	*/
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("none")
	}

	/**
	8.循环
	*/
	i1 := 0
	sum := 0
	for i1 < 10 {
		sum += 1
		i1++
	}
	fmt.Println(sum)

	sum2 := 0
	for i2 := 0; i2 < 10; i2++ {
		sum2 += 1
	}
	fmt.Println(sum2)

	/**
	9.指针
	指针是一个地址，地址指向一个值。
	*/
	var ap *int
	a6 := 12
	//&可以获取变量的地址
	ap = &a6
	//*访问指针指向的值
	fmt.Println(*ap)
	//将结构体作为参数传递或者为已定义类型声明方法时，通常首选指针
	//值传递，赋值意味着更多的内存
	//传递指针，函数更改的值将反应在方法调用者中

	i3 := 10
	increment(&i3)
	fmt.Println(i3)

	/**
	10.函数
	*/
	fmt.Println(add(2, 1))
	fmt.Println(add1(2, 2))

	//Go 函数可以返回多个返回值，将多个返回值用逗号分隔开即可:
	sum3, message3 := add2(2, 12)
	fmt.Println(message3)
	fmt.Println(sum3)

	/**
	11.结构体
	*/
	//创建person实例
	//方式1:指定属性和值
	p1 := person{name: "Bob", age: 42, gender: "Male"}
	fmt.Println(p1.name)
	//方式2:指定值
	p2 := person{"Bob", 42, "Male"}
	fmt.Println(p2.name)
	fmt.Println(p2.age)
	fmt.Println(p2.gender)

	//指针访问结构体里面的属性
	pp := &person{"zhangsan", 1, "Female"}
	fmt.Println(pp.name)

	/**
	12.方法
	*/
	pp1 := &person{"lisi", 42, "Male"}
	pp1.describe()
	pp1.setAge(45)
	pp1.setName("Hari")
	fmt.Println(pp1.name)

	/**
	13.接口
	是一系列方法的集合
	*/
	var a8 animal
	a8 = snake{Poisonous: true}
	fmt.Println(a8.description())
	a8 = cat{Sound: "Meow!!!"}
	fmt.Println(a8.description())

}

func (p person) wake() {

}

// 定义接口
type animal interface {
	description() string
}

// 结构体
type cat struct {
	Type  string
	Sound string
}
type snake struct {
	Type      string
	Poisonous bool
}

func (s snake) description() string {
	return fmt.Sprintf("Poisonous: %v", s.Poisonous)
}

func (c cat) description() string {
	return fmt.Sprintf("Sound: %v", c.Sound)
}

// 定义方法
func (p *person) describe() {
	fmt.Printf("%v is %v years old", p.name, p.age)
	fmt.Println("")
}
func (p *person) setAge(age int) {
	p.age = age
}
func (p person) setName(name string) {
	p.name = name
}

// 结构体是不同字段类型集合
type person struct {
	name   string
	age    int
	gender string
}

func increment(i *int) {
	*i++
}

func add(a int, b int) int {
	c := a + b
	return c
}

// Go 函数可以返回多个返回值，将多个返回值用逗号分隔开即可:
func add1(a int, b int) (c int) {
	c = a + b
	return
}

func add2(a int, b int) (int, string) {
	c := a + b
	return c, "successfully added"
}
