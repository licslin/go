package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/cmplx"
)

//go语言学习基础语法   再复习敲一次

//----------------------变量定义学习--------------------------
func main() {

	fmt.Println("hello licslan")
	variable()
	variableTypeDeduction()
	variableshort()
	fmt.Println(j, k, jj, kk)
	eular()
	triangle()
	consts()
	enums()
	bounded(98)
	play()
	eval(100, 200, "+")

}

//定义变量   go定义了变量一定要用到  不用到是不行的
func variable() {
	//默认值为""
	var a string
	//默认值为0
	var b int
	//赋值初始值
	var c int = 8
	fmt.Println(a, b, c)

	fmt.Printf("%d %q\n", b, a)

}

//go语言可以推断我们所定义的的变量类型
func variableTypeDeduction() {
	//不用写变量的类型  让编译器自动决定类型
	var a, b, c, d = 3, 4, true, "xxxxxxxxxx"
	var e = 4.0
	fmt.Println(a, b, c, d, e)

}

//变量定义简写省略var  用冒号表示  ：
func variableshort() {
	//带有var定义变量
	var g, f = true, "hhhhhhhhh"
	//不用写变量的类型  省略方式写
	a, b, c, d := 3, 4, true, "gggggggggggggg"
	e := 4.0
	//冒号定义后可以再给变量赋值 此时不用加冒号  否则重复定义变量了
	e = 8 //不可以e:=8
	fmt.Println(a, b, c, d, e, f, g)
}

//在函数外面不可以用冒号定义变量  必须用var关键字定义变量
//包内部变量  没有全局变量的说法
var j = 0 //NO  j:=0
var k = "c"

//上面包变量可以改写  当然函数内也可以这样定义变量
var (
	jj = 0
	kk = "cc"
)

//go内建变量类型
//bool string
//(u)int   (u)int8  (u)int16  (u)int32  (u)int64  uintptr<指针的意思>
//byte  rune<go语言的字符型  长度32位的  也就是我们通常理解的char类型>
//float32  float64<浮点型>
//complex64   complex128<高中大学数学里的复数 实部和虚部是64位的>
//i=√-1   复数: 3+4i  |3+4i|=√3²+√4²=5  i²=-1,i³=-i,(i²)²=1....

//euler公式
func eular() {
	//c:=3+4i
	//fmt.Println(cmplx.Abs(c))
	//fmt.Println(cmplx.Pow(math.E,1i*math.Pi+1))
	//fmt.Println(cmplx.Exp(1i*math.Pi)+1)    写成1i会被识别成复数格式 python也会这样
	//取小数点后3位
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)

}

//强制类型转换  类型转换是强制的
var a, b = 3, 4

func triangle() {

	var (
		a, b = 3, 4
		c    int
	)
	//强制转 才可以
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//go语言常量的定义
const filename = "licslan.txt"

func consts() {
	const name = "hh"
	const a, b = 3, 4
	const (
		c = 5
		f = "xxx"
	)
	fmt.Println(name, a, b, filename, c, f)
}

//go语言常量枚举类型
func enums() {
	const (
		golang = 0
		python = 1
		java   = 2
		c      = 3
	)
	const (
		a = 0
		f = 2
		d = 3
	)
	const (
		//自增长
		aa = iota
		bb
		_ //跳过去了
		ff
		dd
	)
	//b,kb,mb,gb,tb,pb
	const (
		//自增长
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(golang, python, java, c, a, f, d, aa, bb, ff, dd, b, kb, mb, gb, tb, pb)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

//总结
//变量类型写变量名之后
//编译器可推测变量类型
//没有char 只有runes
//原生支持复数类型

//----------------------条件语句学习--------------------------
//if条件语句里面不需要括号
func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

func play() {
	const filename = "licslan.in"
	//file, err := ioutil.ReadFile(filename)
	//if err!=nil{
	//	fmt.Println(err)
	//}else {
	//	fmt.Printf("%s\n",file)
	//}
	//上述可以简写

	if file, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", file)
	}

	//if的条件可以赋值
	//if的条件里赋值的变量的作用域就在这个if条件语句里面
}

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsported op" + op)
	}
	return result
}

//----------------------循环语句学习--------------------------

//----------------------函数定义学习--------------------------

//----------------------指针学习--------------------------
