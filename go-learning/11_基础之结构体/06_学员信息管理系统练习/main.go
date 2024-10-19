//# 学员信息 管理系统 练习
//1. 添加学员信息
//2. 编辑学员信息
//3. 展示所有学员信息

package main

import "C"
import "fmt"

type students struct {
	name   string
	age    int
	sex    string
	classs string
}

type studentManage struct {
	allStudent []*students
}

func newStudent(name string, age int, sex string, classs string) *students {
	fmt.Println("有新同学需要加入。。")
	s := &students{
		name:   name,
		age:    age,
		sex:    sex,
		classs: classs,
	}
	s.addStudent()
	return s
}

func (s students) addStudent() {
	fmt.Println("开始添加学员信息")
	fmt.Println("学员姓名:", s.name)
	fmt.Println("学员年龄:", s.age)
	fmt.Println("学员性别:", s.sex)
	fmt.Println("学员班级:", s.classs)
	S1 := students{
		name:   s.name,
		age:    s.age,
		sex:    s.sex,
		classs: s.classs,
	}
	fmt.Println("学员信息添加成功", S1)

}
func (s students) editStudent() (name string, age int, sex string, classs string) {
	fmt.Println("开始编辑学员信息")
	return s.name, s.age, s.sex, s.classs
}

func (s students) showAllStudent() (name string, age int, sex string, classs string) {
	fmt.Println("开始展示所有学员信息")
	return s.name, s.age, s.sex, s.classs
}

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1. 添加学员信息")
	fmt.Println("2. 编辑学员信息")
	fmt.Println("3. 展示所有学员信息")
	fmt.Println("4. 退出系统")
}

func main() {

	//1. 打印菜单等待用户选择
	showMenu()
	//2. 根据用户选择执行相应功能
	//使用scanf()函数获取用户输入
	var choice int
	fmt.Println("请输入您的选择:  ")
	fmt.Scanf("%d", &choice) //等待用户的输入
	fmt.Println("您选择了", choice)
	//3. 执行对应功能
	switch choice {
	case 1:
		fmt.Println("添加学员信息")
		newStudent("张三", 18, "男", "1班")
	case 2:
		fmt.Println("编辑学员信息")
	case 3:
		fmt.Println("展示所有学员信息")
	case 4:
		fmt.Println("退出系统")
	default:
		fmt.Println("无效的选择，请重新选择")
	}
}
