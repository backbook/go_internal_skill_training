package main

import "fmt"

/**
go 语言匿名函数语法,和闭包使用，闭包指的是可以直接引用这参数，如i
*/
func main() {

	i := 10

	anonymousF1 := func() {
		fmt.Println(i)
	}
	anonymousF1()

	type funcType func()

	var f2 funcType

	f2 = anonymousF1

	f2()

	func(a int) {
		fmt.Printf("this is %c", a)
	}(99)
	fmt.Println()

	//匿名函数有参数返回值
	max, min := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 23)

	fmt.Println(max, min)

}
