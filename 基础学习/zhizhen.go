package main

import "fmt"

/*
	指针
*/

func main() {

	var a int = 1
	var ip *int
	ip = &a
	fmt.Printf("指针地址:%x 指针对应值:%d", ip, *ip)
}
