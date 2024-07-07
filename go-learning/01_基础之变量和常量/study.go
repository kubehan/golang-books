package main

import "fmt"

func main() {
	// 变量
	var age int = 20
	var name string = "Alice"
	fmt.Println("年龄:", age)
	fmt.Println("名字:", name)

	// 变量的类型推断
	height := 1.78
	fmt.Println("身高:", height)

	// 常量
	const pi = 3.14159
	const message = "Hello, World!"
	fmt.Println("圆周率:", pi)
	fmt.Println("消息:", message)

	// 常量的类型
	const (
		trueValue  = true
		falseValue = false
	)
	fmt.Println("true:", trueValue)
	fmt.Println("false:", falseValue)

	// 变量的赋值
	age = 21
	fmt.Println("更新后的年龄:", age)

	// 常量的使用
	fmt.Println("常量pi的平方:", pi*pi)

	// 变量的作用域
	if true {
		var innerAge int = 30
		fmt.Println("内层年龄:", innerAge)
	}
	fmt.Println("外层年龄:", age)

	// 常量的作用域
	fmt.Println("常量message在if语句外部仍然可用:", message)
}
