package main

import "fmt"

//不同包下面如何调用  注意大写字母开头才是 public  小写是private
import ".."
import "../../../day_2_basiccontainer/hi"

func main() {

	q:=queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	x:=hi.Licslan{Age: 3}
	fmt.Println(x)
}
