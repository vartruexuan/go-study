package main

import "fmt"

/*
	interface 接口定义
			隐式实现/和php不一样没有实现关键字
*/

// Phone 接口
type Phone interface {
	call()
}

// Iphone 苹果手机
type Iphone struct {
}

func (phone Iphone) call() {
	fmt.Println("我是Iphone")
}

// MiPhone 小米手机
type MiPhone struct {
}

func (phone MiPhone) call() {
	fmt.Println("我是小米手机")
}

func main() {
	var phone Phone

	// 苹果实现（隐式实现了接口）
	phone = new(Iphone)
	phone.call()

	phone = new(MiPhone)
	phone.call()
}
