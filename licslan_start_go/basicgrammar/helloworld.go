package main

import "fmt"
//go语言学习基本语法   再复习敲一次
func main() {

	fmt.Println("hello licslan")
	variable()
	variableTypeDeduction()
	variableshort()
	
}

//定义变量   go定义了变量一定要用到  不用到是不行的
func variable()  {
	//默认值为""
	var a string
	//默认值为0
	var b int
	//赋值初始值
	var c int=8
	fmt.Println(a,b,c)

	fmt.Printf("%d %q\n",b,a)

}

//go语言可以推断我们所定义的的变量类型
func variableTypeDeduction()  {
	//不用写变量的类型
	var a,b,c,d=3,4,true,"xxxxxxxxxx"
	var e=4.0
	fmt.Println(a,b,c,d,e)
	
}

//变量定义简写省略var  用冒号表示  ：
func variableshort()  {
	//带有var定义变量
	var g,f=true,"hhhhhhhhh"
	//不用写变量的类型  省略方式写
	a,b,c,d:=3,4,true,"gggggggggggggg"
	e:=4.0
	//冒号定义后可以再给变量赋值 此时不用加冒号  否则重复定义变量了
	e=8 //不可以e:=8
	fmt.Println(a,b,c,d,e,f,g)
}
