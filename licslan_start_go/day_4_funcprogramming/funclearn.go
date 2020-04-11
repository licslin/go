package main

import "fmt"

func main() {
	//函数与闭包

	//函数式编程VS指针编程
	//函数一等公民 参数  变量 返回值都可以是函数
	//高阶函数  其参数也可以是函数
	//正统式函数式编程  不可变性  不能有状态 只有常量和函数
	//函数只能有一个函数


	//go语言闭包  更为自然   不要修饰如何访问自由变量
	//go语言没有Lambda表达式  但是有匿名函数


	a:=adder()//返回值是个函数  参数为int类型
	for i:=0;i<10;i++{
		fmt.Printf("0+1+2+...+%d=%d\n",i,a(i))
	}

	//闭包：函数体里面有局部变量+自由变量
}

//返回值是函数  返回值是函数并且其参数是int
func adder() func(int) int  {
	//自由变量 可以改变
	sum:=0
	//返回一个闭包
	return func(v int) int {
		//v局部变量
		sum+=v
		return sum
	}
	
}

//稍微正统一点的函数式编程  没有变量  只有函数+常数
type iAdder func(int) (int,iAdder)
func adder2(base int)iAdder  {
	return func(v int) (int, iAdder) {
		return base+v,adder2(base+v)
	}
}

//python原生支持闭包  使用_closure_来查看闭包内容
//c++
//java1.8 使用function接口和Lambda表达式来创建函数对象 匿名类也支持闭包
