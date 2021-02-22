package oop

import "fmt"

//go语言的封装性，体现
type Animal struct {
	name string
}

func (animal *Animal) setName(name string) {
	animal.name = name
}

func (animal *Animal) GetName() string {
	return animal.name
}

func main() {
	animal := Animal{name: "李四"}
	//animal.setName("张三")
	fmt.Println(animal.GetName())
}
