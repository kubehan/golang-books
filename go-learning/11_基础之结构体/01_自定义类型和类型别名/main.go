package main

import "fmt"

// 自定义类型
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Myint int

type Myint2 Myint

func main() {
	var Myage Myint
	Myage = 1999
	fmt.Println(Myage)
	var Val
}
