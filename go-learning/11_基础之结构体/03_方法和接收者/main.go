package main

import "fmt"

// 方法与接收者
//方法是特定类型的方法，接收者可以指向该类型的值或指针

type Vertex struct {
	Name string
	Age  int
}

// NewVertex 创建一个Vertex
//
//	func 函数名字(传入参数变量 传入参数类型) 返回值 {
//			函数本体
//			return 返回值
//		}
func NewVertex(name string, age int) Vertex {
	Vertex := Vertex{
		Name: name,
		Age:  age,
	}
	return Vertex
}
func NewVertex2(name string, age int) *Vertex {
	Vertex := Vertex{
		Name: name,
		Age:  age,
	}
	return &Vertex
}

// 定义一个方法用于打印Vertex

//	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//			函数本体
//	}
func (v Vertex) SetVertex() {
	fmt.Println("打印Vertex", v)
	fmt.Println(v.Name)
	fmt.Println(v.Age)
}

func (v *Vertex) SetVertex2(name string, age int) {
	fmt.Println("SetVertex2", v)
	v.Name = name
	v.Age = age
	fmt.Println(v.Name)
	fmt.Println(v.Age)

}

func main() {
	p := NewVertex(string("小明"), int(18))
	p.SetVertex()

	p2 := NewVertex2(string("小明"), int(18))
	p2.SetVertex()

	p3 := NewVertex2(string("小明"), int(18))
	p3.SetVertex2(string("小韩"), int(28))

}
