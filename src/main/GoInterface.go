package main

import "fmt"

//interface是类型，既类似object,又以此实现多态的表现
type Sayer interface {
	Say() string
}

type Cat struct {
}

func (c Cat) Say() string {
	return "喵喵"
}

type Dog struct {
}

func (d Dog) Say() string {
	return "汪汪"
}

func Say(s Sayer) {
	fmt.Println(s.Say())
}

func main() {
	//以上代码太过冗余，并且不利于扩展，使用多态的方式
	//如下代码是正常的书写代码，要想通用性，必须使用
	c := Cat{}
	fmt.Println(c.Say())

	d := Dog{}
	fmt.Println(d.Say())

	//实现代码复用，但是反思一下：比起java这种方式差了点意思
	var x Sayer
	x = c
	fmt.Println(x.Say())
	x = d
	fmt.Println(x.Say())

	Say(c)
	Say(d)

}
