package main

import (
	"bufio"
	"fmt"
	"os"
)
import "../../pipeline"
func main() {

	const filename  ="licslan.in"
	const t  =50000
	//mergedemo()

	//创建文件licslan.in
	file, err := os.Create(filename)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	//随机数
	source := pipeline.RandomSource(t)
	//写入文件licslan.in
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer,source)
	writer.Flush()

	//后面我们再来读这个文件里面的数据
	file, err = os.Open("licslan.in")
	if err!=nil{panic(err)}
	defer file.Close()
	//bufio.NewReader 很快
	p := pipeline.ReaderSource(bufio.NewReader(file),-1)
	count:=0
	for v:=range p{

		fmt.Println(v)
		count++
		//打印前100个数据
		if count>10{
			break
		}

	}


}

func mergedemo()  {
	//这个时候是一个channel
	p :=pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3,2,4,6,9)),
		pipeline.InMemSort(pipeline.ArraySource(30,20,40,60,90)))

	//for {
	//	if num,ok :=<- p;ok{
	//		fmt.Println(num)
	//	}else {
	//		break
	//	}
	//
	//}



	//将已经合并的好的多个channel里面的值获取并打印出来
	for v := range p{
		//此时发送方要close掉  fatal error: all goroutines are asleep - deadlock!
		fmt.Println(v)
	}


	fmt.Println("-----------------------")
	fmt.Println(pipeline.StrGet(1))
}

