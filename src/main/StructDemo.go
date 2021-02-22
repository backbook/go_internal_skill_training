package main

import "fmt"

type person struct {
	name   string
	age    int64
	height int64
}

type person1 struct {
	name, city string
	age        int64
	height     int64
}

func main() {
	var p person
	p3 := &person{}
	p3.age = 10

	p.name = "张三"
	p.age = 18
	p.height = 175
	fmt.Println((*p3).age)

}
