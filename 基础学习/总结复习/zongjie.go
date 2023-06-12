package main

import (
	"fmt"
	"time"
)

/*
go 总结

	数据类型
	变量声明（局部变量/全局变量）
	常量
	运算符（逻辑运算符/二进制运算符）
	func
	数组/切片
	map集合
	struct结构体
	指针定义
	interface 接口
	channel管道/select
	for/range/switch
*/

// 全局变量
var (
	g1 int
	g2 string
)

func main() {
	// 变量声明
	fmt.Println("========== 变量声明 ==============")
	// 1.
	var a int
	a = 1
	// 2.
	var b = "我是第二种"
	// 3.
	c := "我是第三种"
	// 4. 多变量声明
	var d, e = 1, "字符串"
	// 全局变量/局部变量同名 优先级局部变量最高
	fmt.Println(a, b, c, d, e)
	g1 = 100
	g2 = "我是全局变量郭g2"
	fmt.Printf("我是全局变量:%d %s \n", g1, g2)
	fmt.Println("========== 常量声明 ==============")
	const const1 int = 1
	const const2 string = "姓名"
	fmt.Println(const1, const2)
	fmt.Println("========== 函数/方法 ==============")
	// 函数
	var s1, s2 = 1, 2
	sumV := sum_value(s1, s2)
	fmt.Printf("函数求和:%d+%d=%d \n", s1, s2, sumV)
	// 方法
	person := Person{name: "小郭", age: 18, sex: "男"}
	person.printPerson()

	fmt.Println("========== 数组/切片 ==============")
	arr1 := []int{1, 2, 3}
	arr2 := [3]string{"小郭", "小黄"}
	fmt.Println(arr1, arr2)
	// 切片
	arr3 := make([]int, 3, 4)
	arr3[0] = 1
	arr3[1] = 2
	arr4 := arr1[:3]
	fmt.Printf("我是切片:arr3:%v arr4:%v \n", arr3, arr4)
	// 切片复制
	arr5 := make([]int, 3, 9)
	copy(arr5, arr1) // 参数1：目标数组 参数2：复制数组
	fmt.Printf("我是切片复制arr5:%v \n", arr5)

	fmt.Println("========== map集合 ==============")
	map1 := map[string]string{"name": "小昭", "age": "11"}
	delete(map1, "name")
	map1["sex"] = "男"
	fmt.Println(map1)

	fmt.Println("========== channel管道/select ==============")
	channel1 := make(chan int)
	channel2 := make(chan string)

	// 协程
	go func() {
		time.Sleep(1000) // 停顿1秒
		// 数据推入管道
		channel1 <- 1
	}()
	go func() {
		// 数据推入管道
		channel2 <- "管道2"
	}()

	// select 关键字接收管道值
	select {
	case channelValue := <-channel1:
		fmt.Printf("我是管道中值：%d \n", channelValue)
	case channelValue := <-channel2:
		fmt.Printf("我是管道2中值:%s \n", channelValue)
	}
	fmt.Println("========== 指针 ==============")
	var z = 1
	var zhizhen1 *int
	zhizhen1 = &z
	fmt.Printf("指针：内存地址：%x 指针对应值:%d \n", zhizhen1, *zhizhen1)

	fmt.Println("========== interface 接口定义（隐式实现） ==============")
	var pen Pen
	pen = RedPen{title: "我是红笔的title"}
	pen.printTitle()
	pen = new(BlackPen)
	pen.printTitle()
}

// 函数
func sum_value(a, b int) int {
	return a + b
}

// Person 方法
type Person struct {
	name string
	age  int
	sex  string
}

func (person Person) printPerson() {
	fmt.Printf("我的姓名是:%s,年龄:%d,性别：%s \n", person.name, person.age, person.sex)
}

/* interface-笔 实现*/
type Pen interface {
	printTitle()
}
type RedPen struct {
	title string
}

func (pen RedPen) printTitle() {
	fmt.Printf("我是红笔调用的:%s \n", pen.title)
}

type BlackPen struct {
}

func (pen BlackPen) printTitle() {
	fmt.Printf("我是黑笔调用的 \n")
}
