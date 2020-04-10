package main

import (
	"fmt"
)

func main() {

	//定义一个channel
	ch :=make(chan string)

	//hello world 并发版编程 channel通信

	//并发加上go  go starts a goroutine

	//加上for循环
	for i :=0;i<10000;i++{
		//普通其他语言根本很难达到开启这么多线程（10000个） 需要异步IO才可以做到 当然这个goroutine 并不是线程
		//利用go语言要开多少个goroutine 都可以轻松做到
		go printhelloworldx(i,ch)
		//如何和同时的跑函数goroutine  进行通信呢？   普通语言需要返回了再进行通信   channel
		//go 语言main函数运行完了 就自动退出  这个时候还没有运行 go printhelloworld()
	}


	for{
		//死循环
		//接受刚才那个channel里面的数据  channel里面提取数据
		msg :=<-ch
		fmt.Println(msg)
	}

	//sleep一下
	//time.Sleep(1*time.Millisecond)

	fmt.Print("finished!")

}

func printhelloworldx(i int,ch chan string)  {
	//不断的输出  死循环
	for{
		//打印的值发送到channel里面去  数据进入到channel
		ch <- fmt.Sprintf("hello world from goroutine %d\n",i)
	}

}
