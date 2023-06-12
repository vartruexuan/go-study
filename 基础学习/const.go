package main

import (
	"fmt"
	"unsafe"
)

/*
	常量
*/

// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过
const (
	ga = "abc"
	gb = len(ga)
	gc = unsafe.Sizeof(ga)
)

func main() {

	// 声明
	const a = 1
	// 枚举
	const (
		boy  = 1
		girl = 2
	)
	// 多重赋值
	const b, c, d = 1, "string", 2

	// iota 关键字 自动+1
	const (
		iota1 = iota
		iota2
		iota3
	)
	// iota 关键字 自动+1 >> 计算模式
	const (
		iota4 = 2 << iota // 2 左移 0 位
		iota5             // 2 左移 1 位
		iota6             // 2 左移 2 位
		iota7 = 3 << iota // 3 左移3 位
	)
	fmt.Println(a, boy, girl, b, c, d, ga, gb, gc)
	fmt.Println(iota1, iota2, iota3)
	fmt.Println(iota4, iota5, iota6, iota7)
}
