package main

import (
	"fmt"
	"reflect"
)

/*

	tag
		主要涉及反引号的使用
		结构体中使用
*/

type Person2 struct {
	name string `field:"name" type:"string"`
}

func main() {

	// 通过反射获取到tag
	person := Person2{name: "小昭"}
	field, _ := reflect.TypeOf(person).FieldByName("name")
	tag := field.Tag
	fmt.Println(tag.Get("field"))

}
