package main

import "fmt"

type dog struct {
	name string
}

type cat struct {
	name string
}

func (d *dog) say() {
	fmt.Println("dog wang wang")
}

func (c *cat) say() {
	fmt.Println("cat  miao  miao")
}

func (d *dog) move() {
	fmt.Println("dog move")
}

type person struct {
	name string
	age  int
}

// 定义一个抽象的类型，允许接受任意实现了say方法的类型
type sayer interface {
	say()
}

func data(args sayer) {
	args.say()
}

func main() {
	cat1 := cat{"cat1"}
	dog1 := dog{"dog1"}
	data(&cat1)
	data(&dog1)
	p1 := person{"p1", 18}
	//data(&p1) //person 不是sayer类型，没有实现say方法，所以报错
	//data(p1)

}
