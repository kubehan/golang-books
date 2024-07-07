package main

import "fmt"

var (
	a = 1
	b = 2
	c = 3
)

// 定义一个常量
const (
	//定义一个常量
	pi   = 3.1415
	e    = 2.7182
	note = "声明了pi和e这两个常量之后，在整个程序运行期间它们的值都不能再发生变化了"
)

func main() {
	const (
		//定义一个常量
		myname = "张三99"
		mynote = "作用域：常量的作用域取决于它所在的代码块，只能在定义它的代码块内部使用"
	)
	fmt.Println("常量", myname, mynote)
	d := b
	//打印名字
	var firstname = "张"
	var lastname = "三"
	var allname = firstname + lastname
	age := 18
	fmt.Println("大家好\n我叫", allname, "我今年", age, "岁")
	fmt.Printf("大家好，我叫%s今年%d岁\n", allname, age)
	fmt.Println("我是变量a,b=d:", a, b, d)
	getconst()
}

func Print_Info() {
	fmt.Println("Go语言基础之变量和常量的笔记")
	fmt.Println("变量是用来存储数据的容器，可以在程序运行时改变其值。")
	fmt.Println("在Go语言中，声明变量需要使用关键字`var`，后面跟上变量名和可选的类型声明。")
	fmt.Println("例如：")
	fmt.Println("```go")
	fmt.Println("var age int = 20")
	fmt.Println("```")
	fmt.Println("也可以不指定类型，让Go根据初始化值自动推断类型：")
	fmt.Println("```go")
	fmt.Println("name := \"Alice\"")
	fmt.Println("```")
	fmt.Println("常量是在程序运行时值不会改变的量。")
	fmt.Println("在Go语言中，声明常量需要使用关键字`const`，后面跟上常量名和对应的值。")
	fmt.Println("例如：")
	fmt.Println("```go")
	fmt.Println("const pi = 3.14159")
	fmt.Println("```")
	fmt.Println("变量的类型在声明时就必须确定，并且在程序的整个生命周期中都不能改变。")
	fmt.Println("Go语言支持多种基本数据类型，如整型（int）、浮点型（float64）、布尔型（bool）、字符串型（string）等。")
	fmt.Println("此外，Go还支持结构体、数组、切片、映射等多种复杂数据类型。")
	fmt.Println("常量的类型在声明时就已经确定，并且在程序的整个生命周期中都不能改变。")
	fmt.Println("常量的类型可以是任何Go语言支持的数据类型，包括基本数据类型和复杂数据类型。")
	fmt.Println("在Go语言中，可以使用等号（=）对变量进行赋值。")
	fmt.Println("例如：")
	fmt.Println("```go")
	fmt.Println("age = 21")
	fmt.Println("```")
	fmt.Println("需要注意的是，赋值操作会改变变量的值，但是不会改变变量的类型。")
	fmt.Println("常量可以在程序的任何地方使用，但是需要注意的是，由于常量的值在程序运行时不会改变，因此在使用常量时应该避免对其进行修改。")
	fmt.Println("如果需要修改某个值，应该将其定义为变量而不是常量。")
	fmt.Println("变量的作用域是指变量在程序中可以被访问的范围。")
	fmt.Println("变量的作用域取决于它所在的代码块，一般来说，函数内部的变量只能在该函数内部访问，而函数外部的变量可以在整个程序中访问。")
	fmt.Println("常量的作用域与变量类似，也取决于它所在的代码块。")
	fmt.Println("常量可以在声明它的代码块以及该代码块内的所有嵌套代码块中使用。")
}
func getconst() {
	//pi = 99999
	fmt.Println("常量", pi, e, note)
	Study_anonymous()
}

// 匿名变量的使用
func Study_anonymous() {
	a, b := 1, 2
	fmt.Println("a,b", a, b)

	//匿名变量不占用命名空间，不会分配内存
	//匿名变量与匿名变量之间也不会因为多次声明而无法使用
	//匿名变量与有名变量名称可以相同
	//匿名变量不占用命名空间，不会分配内存
	//匿名变量与匿名变量之间也不会因为多次声明而无法使用
	//匿名变量与有名变量名称可以相同"
	_, c := 3, 4
	fmt.Println("c", c)
	Print_Info()
}
