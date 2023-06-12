package main

import "fmt"

/*
		数组
	    多维数组
		切片
*/
func main() {

	// 数组
	var arr1 = [3]int{1, 2, 3}
	arr2 := []string{"a", "b"}
	arr2 = append(arr2, "c")
	fmt.Println(arr1)
	fmt.Println(arr2)

	// 遍历
	for i := range arr2 {
		fmt.Println(i)
	}

	fmt.Println("============== 切片 ===============")
	// 切片
	slice1 := []int{1, 2, 3, 4}
	slice2 := slice1[1:2] // 起始索引:结束索引（不包含）
	slice3 := slice1[:2]  // 默认其实索引为0
	slice4 := slice1[1:]  // 默认结束索引 len(slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	// 切片 复制
	slice5 := make([]int, 3, 5)

	copy(slice5, slice1)
	// cap 获取容量
	fmt.Printf("slice:%v len:%d cap:%d \n", slice5, len(slice5), cap(slice5))

	slice5 = append(slice5, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)

	fmt.Printf("slice:%v len:%d cap:%d \n", slice5, len(slice5), cap(slice5))
}
