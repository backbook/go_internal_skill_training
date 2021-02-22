package main

import (
	"fmt"
	"log"
)

/**
本类是探寻go语言方法的多态性
*/

func main() {
	i := add(1, 1)
	fmt.Println(i)

	var fTest funcType //多态的体现，回调函数
	fTest = add
	test := fTest(10, 12)
	fTest = minus
	test = fTest(10, 12)
	log.Println(test)

	//多态的体现
	i2 := calc(1, 1, add)
	fmt.Println(i2)
	i3 := calc(1, 1, minus)
	fmt.Println(i3)

}

func add(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

type funcType func(int, int) int

/**
同一个函数可以使用不一样的计算逻辑,多态的表现
*/
func calc(a int, b int, fTest funcType) int {
	i := fTest(a, b)
	return i
}
