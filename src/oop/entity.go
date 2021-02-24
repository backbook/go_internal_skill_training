package oop

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  byte
}

type Student struct {
	Person  //只有类型，没有名字，匿名字段，相当于继承了Person的 ,匿名字段
	Name    string
	Score   int
	Address string
}

func (student *Student) SetName(name string) {
	student.Name = name
}

func (student *Student) GetName() string {
	return student.Name
}

func (student *Student) PrintInfoT() {
	fmt.Println("this is student", student)
}

func (person *Person) PrintInfo() {
	fmt.Println("this is person", person)
}

func Amonymous() {
	//匿名结构体
	var s struct {
		name string
	}
	s.name = "王五"
	fmt.Print(s.name)
}
