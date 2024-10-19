package main

import (
	"encoding/json"
	"fmt"
)

// 结构体中的字段首字母大小写敏感，大写的字段才可以被外部访问，小写的字段只能在本包中使用。
// json 与Str的相互转换
// 创建一个10个学生的班级
type Student struct {
	Name string
	ID   int
}

func NewStudent(name string, id int) Student {
	return Student{
		Name: name,
		ID:   id,
	}
}

//type class struct {
//	Title    string //变成小写后将对外部不可见
//	Students []Student
//}

type class struct {
	Title    string    `json:"title"`                         //json标签，可以修改json的字段名
	Students []Student `json:"students_list" db:"student_db"` //json标签，可以修改json的字段名
}

func GetStrclass(jsonData string) {
	var Newclass class                          //无法直接修改原来的值，因此需要定义一个变量
	json.Unmarshal([]byte(jsonData), &Newclass) //反序列化
	fmt.Printf("%+v\n", Newclass)
	//fmt.Println(Newclass)
}

func main() {
	C1 := class{
		Title:    "Go语言高级火箭班", //变成小写后将对外部不可见
		Students: make([]Student, 0, 100),
	}
	// 添加10个学生
	for i := 1; i < 11; i++ {
		tmpStu := NewStudent(fmt.Sprintf("第%d个学生", i), i)
		C1.Students = append(C1.Students, tmpStu)
	}
	//fmt.Printf("%+v\n", C1)
	jsonData, err := json.Marshal(C1) //序列化
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}
	fmt.Println(string(jsonData)) //打印出被修改的json
	GetStrclass(string(jsonData))
}
