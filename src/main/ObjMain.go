package main

import (
	"encoding/json"
	"fmt"
	"go_internal_skill_training/src/oop"
)

type human struct {
	Name string
	Age  int
}

type teach struct {
	*human
	Education string
}

func main() {

	var s1 oop.Student = oop.Student{}
	s1.SetName("张三")
	//%+v 显示更加详细的
	name := s1.GetName()
	fmt.Println(name)
	//匿名结构体
	oop.Amonymous()
	fmt.Println()

	teachObj := &teach{Education: "college", human: &human{Name: "张三", Age: 18}}
	fmt.Printf("%v\n", teachObj)
	jsonData, err := json.Marshal(teachObj)
	if err != nil {
		fmt.Println(err)
	}
	sprintf := fmt.Sprintf("%s", jsonData)
	fmt.Println(sprintf)

}
