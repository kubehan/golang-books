package main

import "fmt"

//嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名。

type Address struct {
	Province    string
	City        string
	UpdatedTime string //嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名。
}

type PersonInfo struct {
	Name  string
	Age   int
	Phone int
	Address
	Email
}

type Email struct {
	EmailAddress string
	UpdatedTime  string
}

func main() {
	PersonInit := PersonInfo{
		Name:  "张三",
		Age:   18,
		Phone: 123456789,
		Address: Address{
			Province:    "广东省",
			City:        "深圳市",
			UpdatedTime: "2022-08-01",
		},
		Email: Email{
			EmailAddress: "123456@qq.com",
			UpdatedTime:  "2023-08-01",
		},
	}
	fmt.Println(PersonInit)
	//fmt.Println(PersonInit.UpdatedTime)
	fmt.Println(PersonInit.Address.UpdatedTime)
	fmt.Println(PersonInit.Email.UpdatedTime)
}
