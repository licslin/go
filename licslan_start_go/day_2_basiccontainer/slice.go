package main

import "fmt"

/**
切片的学习  一般使用数组较少  数组变成一个slice  arr[:]  day2
*/
func main() {
	arr := [...]int{0, 1, 2, 4, 5, 6, 3, 4, 6, 7}
	//是一个视图
	s := arr[2:6]
	fmt.Println(s)
	fmt.Println(arr[2:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:6])
	fmt.Println(arr[:])
	s1 := arr[:]
	fmt.Println("before update", s1)
	updateSlice(s1)
	fmt.Println("after update ", s1)

	//slice 本身是没有数据的是对底层array的一个view

	//reslice
	s1 = s1[:2]
	fmt.Println(s1)

	//slice 的实现  ptr 指向array里面slice指定的开头的那个元素  len 长度  capacity 容量
	//slice <-- ptr --len --cap -->
	//slice计算容量时可以向后扩展  不可以向前扩展
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))

	//向slice中添加元素
	fmt.Println("------------add elements to slice-------------")
	s2 := append(s1, 10)
	fmt.Println(s2)
	//添加元素时如果如果超越cap，系统会重新分配更大的底层数组
	//由于值传递的关系  必须接收append的返回值
	//s = append(s,val)
	var x []int //Zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(x)
		x = append(x, 2*i+1)
	}
	fmt.Println(x)
	xxx := []int{2, 3, 4}
	fmt.Println(xxx)

	//创建一个大小16的slice
	xx := make([]int, 16)
	fmt.Println(xx)
	//创建一个长度是10  cap是32 的slice
	g := make([]int, 10, 32)
	fmt.Println(g)

	//copy slice
	fmt.Println("------------copy slice-------------")
	copy(g, xx)
	fmt.Println(g)

	//delete slice
	fmt.Println("------------delete e'lements from slice-------------")
	//比如说删掉 第3个元素
	d := append(x[:3], x[4:]...)
	fmt.Println(d)

	//删除头尾元素
	fmt.Println("------------popping from front-------------")
	front := x[0]
	x = x[1:]
	fmt.Println("------------popping from back-------------")
	tail := x[len(x)-1]
	x = x[:len(x)-1]
	fmt.Println(front, tail)
	fmt.Println(x)
}

func updateSlice(s []int) {
	s[0] = 100
}
func printSlice(s []int) {
	fmt.Printf("val=%v,len=%d,cap=%d\n", s, len(s), cap(s))
}
