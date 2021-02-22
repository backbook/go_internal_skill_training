package main

import "fmt"

func main() {

	var name []int
	//添加append函数
	name = append(name, 10)
	fmt.Println(len(name), name[0])
	if name == nil {
		fmt.Printf("%s", name)
	}
	//copy 函数将a切片复制到c切片函数
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a)
	fmt.Println(a)
	fmt.Println(c)
	c[0] = 1000
	fmt.Println(a)
	fmt.Println(c)

}
