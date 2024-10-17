package main

import (
	"encoding/json"
	"fmt"
)

//结构体的初始化
//使用键值对进行初始化

type InitPerson struct {
	Age                int
	Name, Sex, Address string
}

func main() {
	p := InitPerson{
		Age:     18,
		Name:    "小明",
		Sex:     "男",
		Address: "北京",
	}
	JsonP, _ := json.Marshal(p)
	fmt.Println(string(JsonP))
	fmt.Printf("%T\n", p.Age) //%T 打印类型
	fmt.Printf("%v\n", p.Age)
}
