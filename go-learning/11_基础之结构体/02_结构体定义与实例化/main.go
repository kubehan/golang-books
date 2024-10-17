package main

import (
	"encoding/json"
	"fmt"
)

//定义结构体

type Student struct {
	Name  string
	Age   int
	Sex   string
	Class string
}

// 类型相同的字段可以写在同一行,会比较省内存
type Student2 struct {
	Name, Sex, School string
	Age, Class        int
}

func main() {
	//实例化结构体：定义变量内容，只有实例化后才能使用，才会实际占用内存
	var stu1 Student

	stu1.Name = "张三"
	stu1.Age = 18
	stu1.Sex = "男"
	stu1.Class = "1班"
	fmt.Println(stu1)
	fmt.Printf("%+v\n", stu1)
	var stu2 Student2
	stu2.Name = "李四"
	stu2.Age = 19
	stu2.Sex = "女"
	stu2.School = "清华北大"
	fmt.Println(stu2)
	fmt.Printf("%+v\n", stu2)

	// 将结构体转换为 JSON 字符串
	jsonStu1, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}
	fmt.Println(string(jsonStu1))
}

//指针类型的结构体
