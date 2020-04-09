package main

import "fmt"

/**
切片的学习
*/
func main() {

	arr := [...]int{0, 1, 2, 4, 5, 6, 3, 4, 6, 7}
	s := arr[2:6]
	fmt.Println(s)
	fmt.Println(arr[2:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:6])
	fmt.Println(arr[:])

}
