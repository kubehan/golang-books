package main

import "fmt"

func main() {
	// 定义用户输入数据的变量
	var name string
	var age int
	var salary float64
	var isgift bool

	// 提醒用户输入姓名
	fmt.Println("嘿，朋友！先告诉我你的大名吧：")
	fmt.Scanf("%s\n", &name)

	// 提醒用户输入年龄
	fmt.Println("哇，" + name + "，你今年高寿啊？")
	fmt.Scanf("%d\n", &age)

	// 提醒用户输入薪水
	fmt.Println("听说你月薪过万了？来，告诉我你的真实收入吧：")
	fmt.Scanf("%f\n", &salary)

	// 提醒用户输入是否单身
	fmt.Println("最后一个问题，你最近有没有收到什么神秘的 gift？（输入 true 或 false）")
	fmt.Scanf("%t\n", &isgift)

	// 输出用户所有的信息
	fmt.Println("好啦，总结一下：")
	fmt.Println("名字是", name, "，年龄是", age, "岁，薪水是", salary, "元，是否单身", isgift)
	fmt.Println("哇塞，" + name + "，你真是个传奇人物！")

	// 突出 fmt.Scanf 的特点：一行输入多个参数
	fmt.Println("现在我们来玩个高级的，一次输入多个信息！")
	fmt.Println("请输入你的姓名、年龄、薪水和是否单身（用空格分隔）：")
	fmt.Scanf("%s %d %f %t\n", &name, &age, &salary, &isgift)

	// 再次输出用户的所有信息
	fmt.Println("再次总结一下：")
	fmt.Println("名字是", name, "，年龄是", age, "岁，薪水是", salary, "元，是否单身", isgift)
	fmt.Println("你真是太棒了，" + name + "，简直是个天才！")
}
