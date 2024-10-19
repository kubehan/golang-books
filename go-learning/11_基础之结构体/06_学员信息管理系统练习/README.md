# GoLang如何获取用户输入
### fmt.Scanln() 与 fmt.Scanf()
> 需求场景：
数就非常有用。比如，创建一个任务管理工具，可能需要用户输入任务名称、截止日期等信息，这时就可以使用 ```fmt.Scanln ()``` 或 ```fmt.Scanf ()``` 来接收用户输入。

>对于一些需要用户交互的小型应用程序，比如简单的计算器程序，可能需要用户输入数字和运算符，fmt.Scanf () 可以按照指定的格式接收这些输入，确保输入的准确性

## fmt.Scanln()包的使用
```fmt.Scanln()```函数是 Go 语言中用于接收用户输入的函数之一。使用时，首先需要声明变量来存储用户输入的数据

然后，通过```fmt.Println()```函数提醒用户输入相应的内容，当程序执行到```fmt.Scanln(&name)```时，程序会停止在这里，等待用户输入，并在用户回车后继续执行。对于不同类型的数据，都需要分别调用```fmt.Scanln()```来接收用户输入。
代码案例：
```go
package main

import "fmt"

func main() {
	// 定义用户输入数据的变量
	var name string
	var age int
	var salary float64
	var isOffer bool

	// 提醒用户输入姓名
	fmt.Println("嘿，朋友！先告诉我你的大名吧：")
	fmt.Scanln(&name)

	// 提醒用户输入年龄
	fmt.Println("哇，" + name + "，你今年高寿啊？")
	fmt.Scanln(&age)

	// 提醒用户输入薪水
	fmt.Println("听说你月薪过万了？来，告诉我你的真实收入吧：")
	fmt.Scanln(&salary)

	// 提醒用户输入是否拿到 offer
	fmt.Println("最后一个问题，你最近有没有收到什么神秘的 offer？（输入 true 或 false）")
	fmt.Scanln(&isOffer)

	// 输出用户所有的信息
	fmt.Println("好啦，总结一下：")
	fmt.Println("名字是", name, "，年龄是", age, "岁，薪水是", salary, "元，是否拿到 offer 是", isOffer)
	fmt.Println("哇塞，" + name + "，你真是个传奇人物！")
}
```

## fmt.Scanf() 包的使用
```fmt.Scanf()```函数可以按指定的格式输入输入值，使用**空格隔开**。其语法格式为```fmt.Scanf(format string, a...interface{})```，其中```format```是指定的格式字符串，后面的可变参数是接收输入值的变量地址。这个函数会根据指定的格式去读取由空白符分隔的值，并保存到传递给本函数的参数中。例如，如果要接收一个字符串和一个整数，可以使用```fmt.Scanf("%s %d", &str, &num)```这样的格式。
**需要注意的是如果%d后面不加\n，则会出现异常**
来个实际案例：
```go
package main

import "fmt"

func main() {
	// 定义用户输入数据的变量
	var name string
	var age int
	var salary float64
	var isOffer bool

	// 提醒用户输入姓名
	fmt.Println("嘿，朋友！先告诉我你的大名吧：")
	fmt.Scanf("%s\n", &name)

	// 提醒用户输入年龄
	fmt.Println("哇，" + name + "，你今年高寿啊？")
	fmt.Scanf("%d\n", &age)

	// 提醒用户输入薪水
	fmt.Println("听说你月薪过万了？来，告诉我你的真实收入吧：")
	fmt.Scanf("%f\n", &salary)

	// 提醒用户输入是否拿到 offer
	fmt.Println("最后一个问题，你最近有没有收到什么神秘的 offer？（输入 true 或 false）")
	fmt.Scanf("%t\n", &isOffer)

	// 输出用户所有的信息
	fmt.Println("好啦，总结一下：")
	fmt.Println("名字是", name, "，年龄是", age, "岁，薪水是", salary, "元，是否拿到 offer 是", isOffer)
	fmt.Println("哇塞，" + name + "，你真是个传奇人物！")

	// 突出 fmt.Scanf 的特点：一行输入多个参数
	fmt.Println("现在我们来玩个高级的，一次输入多个信息！")
	fmt.Println("请输入你的姓名、年龄、薪水和是否拿到 offer（用空格分隔）：")
	fmt.Scanf("%s %d %f %t\n", &name, &age, &salary, &isOffer)

	// 再次输出用户的所有信息
	fmt.Println("再次总结一下：")
	fmt.Println("名字是", name, "，年龄是", age, "岁，薪水是", salary, "元，是否拿到 offer 是", isOffer)
	fmt.Println("你真是太棒了，" + name + "，简直是个天才！")
}
```
感兴趣的用这两个代码区执行一下看看区别！

### 特点比对：
```fmt.Scanln()```的主要特点是在遇到换行时停止扫描。与```fmt.Scanf()```相比，```fmt.Scanln()```是逐行接收用户输入，用户需要逐行输入不同类型的数据，并且在输入完一行数据后回车，程序才会继续接收下一行数据。而```fmt.Scanf()```可以通过指定格式一次性接收多个不同类型的数据。与```fmt.Scan()```相比，```fmt.Scanln()```在遇到换行时会立刻停止扫描，而```fmt.Scan()```在遇到换行时可能会继续等待更多的输入，直到满足输入的数量要求。总的来说，```fmt.Scanln()```在需要用户逐行输入数据的场景下非常有用，它可以确保输入的准确性和完整性。


1. ```fmt.Scanln()```是逐行输入，用户在输入完每一行数据后回车，程序才会继续接收下一行数据。例如，在接收用户的姓名、年龄、薪水和是否拿到 offer 的信息时，需要分别在不同的行输入每个数据项。

2. ```fmt.Scanf()```则是按照指定的格式一次性输入多个不同类型的数据，数据之间用空格隔开，最后以回车键结束输入。比如在接收相同的信息时，可以在一行中按照指定格式输入所有数据。

