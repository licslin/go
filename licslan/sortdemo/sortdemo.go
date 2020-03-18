package main

import (
	"fmt"
	"sort"
)

func main() {

	//create a slice of int  排序
	a := []int{1,2,4,5,789,90,100}
	sort.Ints(a)


	//for循环遍历
	for i,v :=range a{
		fmt.Println(i,v)
	}

	//for循环遍历
	for _,v :=range a{
		fmt.Println(v)
	}
	
}
