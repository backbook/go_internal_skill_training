package main

import "fmt"

type Mover interface {
	move()
}

type dog struct {
	name string
}

func (d dog) move() {
	fmt.Println(d.name, "会跳")
}

func main() {

	//考虑下go语言的指针地址的位置
	var m Mover
	var wangcai = dog{name: "wangcai"}
	m = wangcai
	m.move()
	var fugui = &dog{name: "fugui"}
	m = fugui
	fmt.Printf("%p\n", fugui)
	m.move()
	var tianyuan = &dog{name: "tianyuan"}
	m = tianyuan
	fmt.Printf("%p\n", tianyuan)
	m.move()

}
