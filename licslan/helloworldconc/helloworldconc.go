package main

import (
	"fmt"
	"time"
)

func main() {

	//hello world 并发版编程

	//并发加上go  go starts a goroutine

	//加上for循环
	for i :=0;i<10000;i++{
		//普通其他语言根本很难达到开启这么多线程（10000个） 需要异步IO才可以做到 当然这个goroutine 并不是线程
		//利用go语言要开多少个goroutine 都可以轻松做到
		go printhelloworld(i)
		//如何和同时的跑函数goroutine  进行通信呢？   普通语言需要返回了再进行通信   channel




		//go 语言main函数运行完了 就自动退出  这个时候还没有运行 go printhelloworld()
	}

	//sleep一下
	time.Sleep(1*time.Millisecond)

	fmt.Print("finished!")

}

func printhelloworld(i int)  {
	//不断的输出  死循环
	for{
		fmt.Printf("hello world from goroutine %d\n",i)
	}

}
