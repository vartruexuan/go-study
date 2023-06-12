package main

import (
	"fmt"
	"time"
)

/*
	管道channel/switch/select/for
*/

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)
	var a = 1
	// 协程
	go func() {
		time.Sleep(1000)
		chan1 <- a // 向管道chan1推值
	}()
	go func() {
		time.Sleep(1000)
		chan2 <- "hello world" // 向管道chan1推值
	}()
	var b int
	b = <-chan1  // 向管道chan1取值
	c := <-chan2 // 向管道chan2取值
	fmt.Println(b)
	fmt.Println(c)
	// select + channel 使用
	go func() {
		for {
			chan1 <- 100 // 向管道chan1推值
		}
	}()
	go func() {
		time.Sleep(2000)
		chan2 <- "测试chan2-select" // 向管道chan1推值
	}()
	select {
	case d := <-chan1:
		fmt.Printf("我是select-chan1:%d \n", d)
	case e := <-chan2:
		fmt.Printf("我是select-chan2:%s \n", e)
	}
	// for
	i := 100
	for i > 0 {
		i--
		switch i {
		case 1:
			fmt.Println("我是case1")
		case 2:
			fmt.Println("我是case2")
		default:
			//fmt.Println("我是default")
		}

	}

}
