package main

import "fmt"

func init() {
	fmt.Println(1)
}

func main() {

	a := 10
	b := &a
	fmt.Println(b)
	c := *b
	fmt.Println(c)
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)

}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 1000
}
