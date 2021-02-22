package main

import "fmt"

//使用结构体代替类
type Tool struct {
	values int
}

var instance *Tool

func getInstance() *Tool {
	if instance == nil {
		instance = new(Tool)
	}
	return instance
}

type Tool1 struct {
}

var tool1 *Tool1

func init() {
	tool1 = new(Tool1)
}

func getTool() *Tool1 {
	return tool1
}

func main() {
	tool := getInstance()
	tool.values = 100
	fmt.Println((*tool).values)
}
