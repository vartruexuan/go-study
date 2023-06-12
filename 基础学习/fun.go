package main

import "fmt"

/*
函数/方法/闭包
*/

// 函数
func sum(a, b int) int {
	return a + b
}

// 方法
type book2 struct {
	title string
	price int
}

func (book book2) printBook() {
	fmt.Printf("我是方法执行:书名:%s 价格：%d", book.title, book.price)
}

func main() {
	// 执行函数
	var a, b = 1, 2
	sumValue := sum(a, b)
	fmt.Println(sumValue)

	// 执行方法
	book := book2{title: "php书籍", price: 200}
	book.printBook()
}
