package main

import "fmt"

type Person1 struct {
	name string
	age  int
	sex  byte
}

type Student1 struct {
	*Person1 //只有类型 ，没有名字，匿名字段，相当于继承了Person的 ,匿名字段
	score    int
	address  string
}

func main() {

	student1 := Student1{&Person1{"张三", 13, 1}, 100, "哈尔滨"}
	fmt.Println(student1)                                                                    //此方式会打印出地址符号，为了能进行遍历，需要打印出成员的字段
	fmt.Println(student1.name, student1.age, student1.sex, student1.score, student1.address) //此方式可以打印出想要的数据

	//先定义遍历
	var s4 Student1
	s4.Person1 = new(Person1) //分配空间,
	s4.name = "张三"
	s4.sex = 1
	s4.address = "福建省"
	fmt.Println(s4)
	fmt.Println(student1.name, student1.age, student1.sex, student1.score, student1.address) //此方式可以打印出想要的数据

	var s5 Student1
	fmt.Println(s5)

	test(&Person1{"张三", 10, 0})

	person := Person1{name: "王五"}
	p := &Person1{"张三", 10, 0}

	fmt.Println(person)
	fmt.Println(p.name)

	res := add01(1, 1)
	fmt.Println(res)

}

func test(person *Person1) {
	fmt.Println(person.name)
}

func add01(a, b int) int {
	return a + b
}
