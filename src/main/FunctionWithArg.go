package main

import "fmt"

func main() {
	argFunc, i, i2 := multiArgFunc()
	fmt.Println(argFunc, i, i2)

	i3 := func07()
	fmt.Println(i3)
	fmt.Printf("%T", i3)
}

func func01() int {
	return 4
}

func func07() interface{} {
	return 10
}

func func03() (res int) {
	res = 555
	return
}

func func04() (res int) {
	return 4
}

func func02() {
	fmt.Println("无参")
}

/**
多参数返回函数 go 官方推荐写法
*/

func multiArgFunc() (int, int, int) {
	return 1, 2, 3
}

func multiArgFunc01() (a int, b int, c int) {
	a, b, c = 1, 2, 3
	return
}
