package main

import "fmt"

func division() {
	i := 100 / 0
	fmt.Println(i)
}

func main() {

	division()

	defer fmt.Println("1313")

}
