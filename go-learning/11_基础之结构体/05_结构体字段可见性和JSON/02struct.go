package main

import "fmt"

type People struct {
	bool bool
	Sex  string
}

func (p *People) IsMoney() bool {
	fmt.Println("人都会挣钱的")
	return true
}

type WuMan struct {
	*People //女人继承了人这个属性
}

type Man struct {
	*People //男人继承了人这个属性
}

func (w *WuMan) IsChild() bool {
	fmt.Println("女人会生孩子")
	return true
}

func (m *Man) IsChild() bool {
	fmt.Println("男人不会生孩子")
	return false
}

func main() {
	isPeople := People{true, "男"}
	//fmt.Println(isPeople.IsMoney())
	//fmt.Println(isPeople.Sex)
	isWoman := WuMan{&isPeople}
	fmt.Println("女人--", isWoman.IsChild())
	fmt.Println("女人--", isWoman.IsMoney())

	isMan := Man{&isPeople}
	fmt.Println("男人--", isMan.IsChild())
	fmt.Println("男人--", isMan.IsMoney())
}
