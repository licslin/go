package main

import "fmt"

//内建容器  数组
func main() {
	//定义数组
	var arr1 [5]int                             //空数组
	arr2 := [3]int{1, 3, 4}                     //3个元素的数组
	arr3 := [...]int{1, 1, 3, 3, 4, 3, 2, 4, 5} //多个元素的数组
	fmt.Println(arr1, arr2, arr3)

}

