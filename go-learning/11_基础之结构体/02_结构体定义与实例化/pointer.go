package main

import (
	"encoding/json"
	"fmt"
)

//指针类型的结构体new
//使用场景：需要初始化一个结构体，但是不知道结构体的具体值，可以使用指针类型的结构体new

type Person struct {
	Name, Sex, Address string
	Age                int
}

func main() {
	var NewPerson = new(Person)
	(*NewPerson).Age = 18    //指针类型
	(*NewPerson).Name = "张三" //指针类型
	(*NewPerson).Sex = "男"   //指针类型
	NewPerson.Address = "北京" //结构体类型
	fmt.Println(*NewPerson)
	fmt.Printf("%T\n", NewPerson) // *main.Person
	JsonNewPerson, _ := json.Marshal(NewPerson)
	fmt.Println(string(JsonNewPerson))

	//取结构体的值进行实例化
	NewPerson2 := &Person{}
	NewPerson2.Sex = "女"
	NewPerson2.Age = 18
	NewPerson2.Name = "李四"
	fmt.Printf("%T\n", NewPerson2) //%T 打印类型
	fmt.Printf("%v\n", NewPerson2) //%v 打印值
	JsonNerPeron2, _ := json.Marshal(NewPerson2)
	fmt.Println(string(JsonNerPeron2))
}
