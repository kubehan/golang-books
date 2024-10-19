package main

import (
	"encoding/json"
	"fmt"
)

// 嵌套结构体包含多个相同的字段的时候,需要使用具体嵌套的结构体
type GrafanaTools struct {
	Url        string
	User       string
	Password   string
	Datasource string
	UpdateTime string // 更新时间 结构体嵌套出现相同的字段的时候
}

type Address struct {
	Ip           string
	Port         string
	UpdateTime   string // 更新时间 结构体嵌套出现相同的字段的时候
	GrafanaTools GrafanaTools
}
type Mydata struct {
	Address      Address
	GrafanaTools GrafanaTools
	UpdateTime   string
}

func main() {
	MyAddress := Address{
		Ip:         "39.108.137.1",
		Port:       "3000",
		UpdateTime: "2021-08-01",
		GrafanaTools: GrafanaTools{
			Url:        "https://grafana.kubehan.cn",
			User:       "admin",
			Password:   "123456",
			Datasource: "Prometheus",
			UpdateTime: "2021-08-01",
		},
	}
	fmt.Println(MyAddress.UpdateTime)
	JsonMyAddress, _ := json.Marshal(MyAddress)
	fmt.Println(string(JsonMyAddress))
}
