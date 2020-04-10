package main

import (
	"fmt"
	"math/rand"
	"time"
)

//内建容器  数组   day 2
func main() {
	//定义数组
	var arr1 [5]int //空数组
	//冒号需要给予初值
	arr2 := [3]int{1, 3, 4}                     //3个元素的数组
	arr3 := [...]int{1, 1, 3, 3, 4, 3, 2, 4, 5} //多个元素的数组
	//二维数组
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3, grid)

	//传统数组遍历方式
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	//i 是数组下标  遍历时打印输出
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	//不打印数组下标  遍历时打印输出
	for _, v := range arr3 {
		fmt.Println(v)
	}

	//求和操作
	sum := 0
	for _, v := range arr3 {
		sum += v
	}
	fmt.Println(sum)

	//
	maxValue := -1
	for i, v := range arr3 {
		if v > maxValue {
			_, maxValue = i, v
		}
	}

	printArray(arr1)
	//printArray(arr2)
	//printArray(arr3)

	//step1: 设置种子数
	rand.Seed(time.Now().UnixNano())
	//step2：获取随机数
	//x := rand.Intn(100) //[0,100)
	//向空数组中添加元素
	for i, v := range arr1 {
		//给新元素赋值
		arr1[i] = rand.Intn(100)
		println(v)
	}

	fmt.Println("after add e...........")

	printArray(arr1)

	//[10]int & [20]int  是不同类型的值
	//调用func f(arr [10]int) 会拷贝数组  和其他大部分语言是不一样的  大部分语言这里是引用传递
	//在go语言中一般不直接使用数组  使用切片
}
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
