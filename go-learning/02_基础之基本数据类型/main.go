package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var name = "基础数据类型"
	fmt.Println(name)

	// 获取当前目录并打印
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("无法获取当前目录: %v", err)
	}
	fmt.Println("当前目录为：", currentDir)

	// 读取当前目录下的所有文件
	files, err := ioutil.ReadDir(currentDir)
	if err != nil {
		log.Fatalf("无法读取当前目录下的文件: %v", err)
	}

	// 遍历文件并打印文件名
	for _, file := range files {
		fileName := file.Name()
		fmt.Println("文件名：", fileName)
	}
}
