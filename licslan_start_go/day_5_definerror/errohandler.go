package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
)
import ".//fib"

func main() {

	//我们作为软件工程师 开发代码是一部分工作  那么针对出现代码报错 我们该如何处理呢？
	//资源管理与出错处理？ catch all the errors 打开文件我们需要关闭   连接数据库 我们需要记得释放连接
	//defer调用在函数结束时发生
	//参数在defer语句时计算
	//defer列表先进后出

	//如何实现统一的错误处理逻辑
	//tryDefer()
	//writeFile00("fib.txt")
	writeFile01("fib.txt")
	tryDefer01()
}

//defer
func tryDefer00()  {
	//fmt.Println(1)   打印 13
	defer fmt.Println(1)  //打印 31  会在函数结束时打印（返回） 不会马上打印的
	defer fmt.Println(2)  //defer 里面有一个栈的  先进后出的  所以  321
	fmt.Println(3)
	return
	fmt.Println(4)    //不会打印
}
func tryDefer01()  {
	for i:=0;i<100;i++{
		//倒着打印
		defer fmt.Println(i)
		if i==30{
			panic("printed to many")
		}

	}
}
//writeFile close资源
func writeFile00(filename string) {
	file, err := os.Create(filename)
	if err!=nil{
		panic(err)
	}
	//defer栈  先进后去的
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f:=fib.Fibonacci()
	for i:=0;i<20;i++{
		fmt.Fprintln(writer,f())
	}
}

func writeFile01(filename string) {
	//读文件出错处理方式
	file, err := os.OpenFile(filename,os.O_EXCL|os.O_CREATE,0666)  //0666 linux  可读可写
	if err!=nil{
		//自己创建的error
		err=errors.New("this is a error!")
		//panic(err)
		//00.fmt.Println("file already exists")
		//01.fmt.Println("Error：",err.Error())
		//02.
		if  pathError,ok:=err.(*os.PathError);!ok{
			panic(err)
		}else {
			fmt.Printf("%s,%s,%s\n",pathError.Op,
				pathError.Err,
				pathError.Path)
		}
		return
	}
	//defer栈  先进后去的
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f:=fib.Fibonacci()
	for i:=0;i<20;i++{
		fmt.Fprintln(writer,f())
	}
}