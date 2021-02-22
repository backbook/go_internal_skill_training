package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//告知编译器应该以多大的内存存储
	var a int64 = 10 //long
	var b int8 = 10  //short
	var e int = 1314 //int
	var c byte = 'b'
	var d uint = 'b'
	var f float64 = 3141.2341
	var name string = "张三"
	//常量
	const age int64 = 10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(name)
	var for_range string = "abcde张三"
	for index, value := range for_range {
		/**
		目前printf支持以下格式的输出，例如：
		printf("%c",a)；输出单个字符。
		printf("%d",a)；输出十进制整数。
		printf("%f",a)；输出十进制浮点数.
		printf("%o",a)；输出八进制数。
		printf("%s",a)；输出字符串。
		printf("%u",a)；输出无符号十进制数。
		printf("%x",a)；输出十六进制数。
		*/
		fmt.Printf("idnex:%d Unicode:%c %d ", index, value, value)
		var result string = fmt.Sprintf("%c", value)
		fmt.Println(result)

	}
	fmt.Println()
	var str1, str2 string = "a", "21"
	var var1, var2, var3 = 1, "2", true
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(unsafe.Sizeof(var1)) //查看使用的数据空间占用
	fmt.Println(unsafe.Sizeof(var2))

	//计算换算小规律
	const (
		B  = 1 << (10 * iota) // 0
		KB = 1 << (10 * iota) //1
		MB = 1 << (10 * iota) //2
		GB = 1 << (10 * iota) //3
		TB = 1 << (10 * iota) //4
		PB = 1 << (10 * iota) //4
	)
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)

	const (
		a1 = iota //0
		b1        //1
		c1        //2
		d1 = "ha" //独立值，iota += 1
		e1        //"ha"   iota += 1
		f1 = 100  //iota +=1
		g         //100  iota +=1
		h  = iota //7,恢复计数
		i         //8
	)

	fmt.Println(a1, b1, c1, d1, e1, f1, g, h, i)

	var str = `
     hello
     world
     世界`
	fmt.Println(str)

}
