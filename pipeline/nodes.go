package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
)

//作为一个库  没有main函数


//<- chin int  只能拿数据  所有里面操作只能放东西 out <- v   数据源是一个数组
func ArraySource(a ...int) <-chan int {

	//创建一个channel
	out :=make(chan int)

	//channel通信  starts a goroutine
	go func() {
		for _,v := range a{
			//将遍历的 v送入channel out
			out <- v
		}
		//在pipeline并行计算里 还是要关闭掉channel  fatal error: all goroutines are asleep - deadlock!
		close(out)
	}()

	return out
}


//排序操作    返回只出不进   相当于别人调用的时候拿东西  只出来
func InMemSort(in <-chan int) <- chan int{
	out :=make(chan int)
	go func() {
		//read into memory  将channel的内存放入到数组里面
		var a []int
		//a :=[]int{}
		for v:=range in{
			a = append(a,v)
		}
		//Sort  排序  对数组排序
		sort.Ints(a)
		//output  输出   将排序好的内存又写回channel
		for _,v :=range a {
			out <- v
		}
		close(out)
	}()
	return out
}


//归并2个channel计算结果（排序结果）
func Merge(in1,in2 <-chan int) <-chan int{


	//从参数可以知道  对多个channel的里面的数据合并成一个channel
	out:=make(chan int)
	go func() {

		//从in1数据存入v1
		v1,ok1:=<-in1
		//从in2数据存入v2
		v2,ok2:=<-in2


		//遍历之前上面的数据其实已经先各自排序好了

		//遍历时只要v1 or v2不为空
		for ok1||ok2{

			//先存v1结果到out中  说明此时v1v2不为空且v1数据小于v2的  排序先排小的数据 或则v2为空了 先存v1数据
			if !ok2||(ok1&&v1<v2){
				out<-v1
				//更新数据 让in1进入到v1  这个时候不能加冒号了：
				v1,ok1=<-in1
			}else {
				out<-v2
				v2,ok2=<-in2
			}
		}

		//记得关闭channel
		close(out)
	}()
	return out


	//原始数据 ---->读取  ----->内部排序  ----->归并  ------>多路归并  ---->展示排序结果
}






//数组数据源节点 channel节点的关闭和检查
//内部排序节点
//归并节点




//数据源是一个文件中的数据  chunkSize 文件分为多少块
func ReaderSource(reader io.Reader,chunkSize int) <-chan int {
	out:= make(chan int)

	go func() {

		//开辟一个大小为8 的字节数组
		buffer:=make([]byte,8)
		//不断从reader去读
		byteReader :=0

		for{
			//第一个参数表示读了多少字节  第二个参数表示读取过程中是否又错误
			r, err := reader.Read(buffer)
			byteReader +=r
			if r>0{
				v:=int(binary.BigEndian.Uint64(buffer))
				out<-v
			}
			if err!=nil || (chunkSize!=-1&&byteReader>=chunkSize){
				//如果出错了 就退出循环
				break
			}
		}


		//记得关闭channel
		close(out)
	}()


	return out

}


//只进来数据  不出  只进不去
func WriteSink(writer io.Writer,in<-chan int)  {
	//写数据
	for v:=range in{
		buffer:=make([]byte,8)
		//将v转换成buffer
		binary.BigEndian.PutUint64(buffer,uint64(v))
		writer.Write(buffer)
	}
}



//随机数据源
func RandomSource(count int)<-chan int  {

	out:=make(chan int)
	go func() {
		//输出随机数
		for i:=0;i<count;i++{
			out<-rand.Int()
		}
		close(out)
	}()
	return out
}



//多路归并
func MergeN(inputs ...<-chan int) <-chan int{

	if len(inputs)==1{
		return inputs[0]
	}
	m:=len(inputs)/2
	//merge inputs[0..m) and inputs(m..end]
	return Merge(MergeN(inputs[:m]...),MergeN(inputs[m:]...))
	
}

func StrGet(s int) int {
	ss := make(chan int)
	go func() {
		var b = 1
		//将b送入channel
		ss <-b
		c :=<-ss
		fmt.Printf("````````````%d",c)
		close(ss)
	}()
	a :=s+1
	return a
}
