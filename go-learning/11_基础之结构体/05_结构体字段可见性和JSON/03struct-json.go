package main

import "fmt"

// 首字母大写的结构体是公开的，可以被其他包访问到，否则是私有的，只能在本包中使用。
type Student struct {
	Name, Class string //对外部可见
	ID          int
}

type password struct {
	Student
	username, password string //对外部不可见
}

func main() {
	stu1 := Student{"张三", "1", 1}
	fmt.Println(stu1)
	//fmt.Printf("%+v\n", stu1)
	pass := password{Student{"李四", "2", 2}, "lisi", "123456"}
	fmt.Println(pass)
}
