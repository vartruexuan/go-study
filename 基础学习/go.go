package main

import (
	"errors"
	"fmt"
	"time"
)

func test(p int) (int, error) {
	if p <= 0 {
		return p, errors.New("错误信息")
	}
	return 1, errors.New("")
}

func main() {

	go func() {
		defer func() {
			fmt.Println("测试defer")
		}()
		fmt.Println("我是协程")

	}()

	fmt.Println("hello")
	time.Sleep(1000)

	result, errorStr := test(0)
	fmt.Println(result, errorStr)
}
