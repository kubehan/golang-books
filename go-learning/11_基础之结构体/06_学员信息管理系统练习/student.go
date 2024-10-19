package main

import "fmt"

type student struct {
	name  string
	class string
	id    int
}
type studentManager struct {
	studentList []*student
}

func NewStudentManager() *studentManager {
	return &studentManager{
		studentList: make([]*student, 0, 100), //初始化切片
	}
}

// 添加学生
func (s *studentManager) AddStudent(NewStudent *student) {
	s.studentList = append(s.studentList, NewStudent) //在已有的切片后面添加元素
	fmt.Println("添加成功", NewStudent.name, NewStudent.class, NewStudent.id)
	fmt.Println(s.studentList)
}

// 显示学生
func (s *studentManager) ShowStudent() {
	for _, v := range s.studentList {
		fmt.Println(v.name, v.class, v.id)
	}
}

// 修改学生
func (s *studentManager) EditStudent(NewStudent *student) {
	for i, v := range s.studentList {
		if v.id == NewStudent.id {
			s.studentList[i] = NewStudent
			fmt.Println("修改成功", NewStudent.name, NewStudent.class, NewStudent.id)
			return
		}
	}
	fmt.Println("输入的学员id信息有误", NewStudent.id)
}

// 删除学生
func (s *studentManager) DeleteStudent(id int) {
	for i, v := range s.studentList {
		if v.id == id {
			s.studentList = append(s.studentList[:i], s.studentList[i+1:]...)
			fmt.Println("删除成功")
			return
		}
	}
	fmt.Println("输入的学员id信息有误", id)
}

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1. 添加学员信息")
	fmt.Println("2. 编辑学员信息")
	fmt.Println("3. 展示所有学员信息")
	fmt.Println("4. 删除学员信息")
	fmt.Println("5. 退出系统")
}
func getInput() *student {
	var (
		name  string
		class string
		id    int
	)
	fmt.Println("请输入学员信息姓名:")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入学员班级信息:")
	fmt.Scanf("%s\n", &class)
	fmt.Println("请输入学员信息id:")
	fmt.Scanf("%d\n", &id)
	NewStudent := &student{
		name:  name,
		class: class,
		id:    id,
	}
	return NewStudent
}

func main() {
	sm := NewStudentManager()
	for {
		//fmt.Println(sm)
		showMenu()
		var choice int
		fmt.Println("请输入您的选择:  ")
		//fmt.Scanln(&choice)
		fmt.Scanf("%d\n", &choice)
		fmt.Println("您选择了", choice)
		switch choice {
		case 1:
			fmt.Println("添加学员信息")
			NewStudent := getInput()
			fmt.Println("输入的学员信息:", NewStudent.name, NewStudent.class, NewStudent.id)
			sm.AddStudent(NewStudent)
		case 2:
			fmt.Println("编辑学员信息")
			NewStudent := getInput()
			fmt.Println("输入的学员信息:", NewStudent.name, NewStudent.class, NewStudent.id)
			sm.EditStudent(NewStudent)
		case 3:
			fmt.Println("展示所有学员信息")
			sm.ShowStudent()
		case 4:
			fmt.Println("删除学员信息")
			NewStudent := getInput()
			sm.DeleteStudent(NewStudent.id)

		case 5:
			fmt.Println("退出系统")
		default:
			fmt.Println("无效的选择，请重新选择")
		}
		fmt.Println("=======================================")
	}
}
