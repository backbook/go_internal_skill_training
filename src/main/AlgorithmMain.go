package main

import "go_internal_skill_training/src/algorithm"

func main() {
	// :=  局部变量，类型自动推导
	recursive := algorithm.Recursive(10)
	println(recursive)
}
