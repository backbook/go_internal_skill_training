package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var scoreMap map[string]int64 = make(map[string]int64, 1)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	scoreMap["王五"] = 98
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = int64(value)
	}
	fmt.Println(scoreMap)
	fmt.Println(-98 % 10)

}
