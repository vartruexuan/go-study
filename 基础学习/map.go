package main

import "fmt"

/*
map集合

	创建/遍历/赋值/删除
*/
func main() {
	// 构建集合
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := make(map[string]string, 5)
	map2["key1"] = "value1"
	map2["key2"] = "value2"
	map2["key3"] = "value3"
	fmt.Println(map1)
	fmt.Println(map2)
	// 删除成员
	delete(map1, "a")
	fmt.Printf("删除成员后：%v \n", map1)
	// 遍历
	for k, v := range map2 {
		fmt.Printf("遍历 key:%s value:%s \n", k, v)
	}
}
