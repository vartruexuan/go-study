package main

/*
	变量声明
*/

import "fmt"

// 全局变量声明
var name = "全局变量"
var (
	age int
	sex int
)

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main() {
	// 自动识别类型
	var a = "字符串"
	// 声明指定类型
	var b int
	b = 1
	// 等同于 var  语法糖
	c := 2
	// 使用全局变量
	age = 1
	sex = 1
	// 多变量声明
	var d, e = 1, "string"
	var f, g int = 1, 2
	h, i := 1, 2
	fmt.Println(a, b, c, name, age, sex, d, e, f, g, h, i)

}
