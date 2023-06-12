package main

import (
	"encoding/json"
	"fmt"
)

type struct3 struct {
	title string
	sex   int
}

/*
json解析/格式化
*/
func main() {
	b, _ := json.Marshal(true)
	str := []string{"测试"}
	map1 := map[string]string{"key": "值"}
	obj := struct3{title: "测试", sex: 1}
	strJson, _ := json.Marshal(str)
	mapJson, _ := json.Marshal(map1)
	objJson, _ := json.Marshal(obj)
	fmt.Println(string(b))       // bool
	fmt.Println(string(strJson)) // 数组/切片
	fmt.Println(string(mapJson)) // 集合
	fmt.Println(string(objJson)) // 结构体
}
