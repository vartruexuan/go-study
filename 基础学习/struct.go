package main

import "fmt"

/*
	结构体
*/

type books struct {
	title string
	price int
}

func main() {
	// 创建一个结构体
	var book books
	book.title = "书名"
	book.price = 2
	// 第二种方式
	var book2 = books{title: "书名2", price: 200}
	fmt.Println(book, book2)
}
