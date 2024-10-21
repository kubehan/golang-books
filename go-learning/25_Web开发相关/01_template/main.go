package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2.解析模板
	template, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	//3. 渲染模板
	//cluster := "Cluster 你好帅"
	type Message struct {
		Name string
		Age  int
		Msg  string
	}
	//msg := Message{
	//	Name: "张三",
	//	Age:  18,
	//	Msg:  "hello world",
	//}
	//myJsonMsg, _ := json.Marshal(msg)
	//fmt.Println(string(myJsonMsg))

	cluster := "hello world"
	name := "张三"
	err = template.Execute(w, cluster)
	err = template.Execute(w, name)
	if err != nil {
		fmt.Printf("execute template failed, err:%v\n", err)
		return
	}
	return
}
func main() {
	//1. 调用模板
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server start failed, err:", err)
	}

}
