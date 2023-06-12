package main

import (
	"../../src/packk"
	"fmt"
	"reflect"
)

/*
1. 声明变量/全局变量/数据类型
2. 常量声明
3. interface 接口定义/多态实现
4. channel 管道实现/结合select
5. 数组/切片
6. map集合
7. pack包引入
8. 结构体/tag 实现
9. 指针 使用
10. for /range
11. ? 类型转换
12  ? 错误处理
*/
var (
	g3 = 1
	g4 = 2
)

// 结构体
type struct1 struct {
	title string
	age   int
}

func main() {
	// 1.声明变量
	var a = 1
	b := 2
	var c, d = 1, "字符串"
	fmt.Printf("a:%d%d%d%s%d%d", a, b, c, d, g3, g4)
	// 2.常量声明
	const c1 = 1
	fmt.Println(c1)
	// 3.数组/切片
	arr := []int{1, 2}
	fmt.Println(arr)
	// 切片
	arr2 := arr[0:1]
	fmt.Print(arr2)
	// 数组复制
	var arr3 []int
	copy(arr3, arr2)
	fmt.Println(arr3)
	// 4. map集合
	map1 := map[string]int{"测试": 1, "delete": 1}
	delete(map1, "delete")
	fmt.Println(map1, map1["测试"])
	// 5. channel/select
	channel1 := make(chan int)
	go func() {
		// 测试管道推入值
		channel1 <- 1
	}()

	// 管道取值
	channelValue1 := <-channel1
	fmt.Println(channelValue1)
	// select 方式
	channel2 := make(chan string)
	channel3 := make(chan string)
	go func() {
		// 推入值
		channel2 <- "测试管道推入值"
	}()
	go func() {
		channel3 <- "我是管道3"
	}()
	select {
	case channel2Value := <-channel2:
		fmt.Println(channel2Value)
	case channel3Value := <-channel3:
		fmt.Println(channel3Value)
	}
	// 结构体
	var struntObject = struct1{title: "测试", age: 12}
	struntObject.title = "测试2"
	fmt.Println(struntObject)

	// interface 接口
	var interfaceObject interface1
	interfaceObject = new(InterfaceClass1)
	interfaceObject.sayHello()
	interfaceObject = InterfaceClass2{name: "测试"}
	interfaceObject.sayHello()
	// 结构体中使用tag/反射获取值
	var class1 = reflect.TypeOf(interfaceObject)
	// 获取的字段必须使用过 - 好严格
	field, result := class1.FieldByName("name")
	if result {
		tag1 := field.Tag.Get("tag1")
		fmt.Printf("我是标签1的值%s", tag1)
	}
	// 指针
	var zhen1 = 1
	var zhen2 *int
	zhen2 = &zhen1
	fmt.Printf("内存地址：%x 指针对应值:%d \n", zhen2, *zhen2)

	// pack 包引用
	fmt.Println(test.Sum3(1, 2))

	// for/range
	map2 := map[int]string{1: "姓名"}
	for k, v := range map2 {
		fmt.Println(k, v)
	}

}

// 接口定义
type interface1 interface {
	// 必须又sayHello方法
	sayHello()
}

type InterfaceClass1 struct {
	name string
	age  int
	sex  string
}

func (interfaceParam InterfaceClass1) sayHello() {
	fmt.Println("我是InterfaceClass1的hello")
}

type InterfaceClass2 struct {
	name string `tag1:"测试"`
	age  int
	sex  string
}

func (interfaceParam InterfaceClass2) sayHello() {
	fmt.Println("我是interfaceClass2的hello")
}
