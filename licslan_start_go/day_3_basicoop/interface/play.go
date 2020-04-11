package main

import (
	"fmt"
	"time"
)
import "../interface/mock"
import "../interface/real"

func main() {
	//duck typing（动态绑定）
	//传统系统类型
	//像鸭子走路 像鸭子叫 （长得像鸭子） 那么就是鸭子
	//描述事务的外部行为而非内部结构
	//严格来说go语言属于结构化类型系统（编译时绑定）  类似duck typing

	//duck typing  ---> python  运行时才知道传入的retriever有没有get  c++ 是编译时才知道传入的retriever有没有get  需要注释来说明接口
	//Java不需要注释来说明接口  传入的参数必须实现Retriever接口  不是duck typing

	//同时需要Readable Appendable 怎么办  （apache polygene）
	//同时具有python c++的duck typing的灵活性
	//又具有Java的类型检查

	//go语言中接口是由使用者定义的  和传统语言不一样的   使用者  实现者

	var r Retriever
	r=mock.Retriever{"this is fake www.licslan.com"}
	fmt.Printf("%T %v\n",r,r)

	//r=real.Retriever{}
	r=real.Retriever{
		UserAgent:"Mozilla/5.0.......cccc.....",
		TimeOut:time.Minute,
	}
	fmt.Printf("%T %v\n",r,r)
	//fmt.Println(download(r))

	//接口的实现是隐式的  实现里面的方法
	inspect(r)
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxx")
	//realRetriever:=r.(*real.Retriever)
	//fmt.Println(realRetriever)

}

type Retriever interface {
	Get(url string) string
}

//使用者
func download(r Retriever) string {
	return r.Get("http://www.licslan.com")
}
func inspect(r Retriever)  {
	switch v:= r.(type) {
	case real.Retriever:
		fmt.Println("UserAgent:",v.UserAgent)
	case mock.Retriever:
		fmt.Println("Contents:",v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:",v.UserAgent)
	}
}

