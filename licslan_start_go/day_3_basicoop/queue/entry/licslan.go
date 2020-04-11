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
	q.Push("abc")
	q.Push("abcddd")
	q.Push("1")
	q.Push(bool(true))
	fmt.Println(q)


	x:=hi.Licslan{Age: 3}
	fmt.Println(x)
}
