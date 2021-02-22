package main

import (
	"fmt"
	"os"
)

/**
此操作类似java的args，只不过这边返回了一个list
*/
func main() {

	args := os.Args

	fmt.Println(args)

	//遍历args
	for i := 0; i < len(args); i++ {
		fmt.Println(i, args[i])
	}

}
