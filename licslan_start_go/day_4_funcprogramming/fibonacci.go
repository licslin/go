package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	//斐波拉切数列
	f:=fibonacci()
	printFileContents(f)
	//fmt.Println(f())//1
	//fmt.Println(f())//1
	//fmt.Println(f())//2
	//fmt.Println(f())//3
	//fmt.Println(f())//5
	//fmt.Println(f())//8
	//fmt.Println(f())//13



}
/*func fibonacci() func()int  {
	a,b:=0,1
	return func() int {
		a,b=b,a+b
		return a
	}
}*/
func fibonacci() intGen {
	a,b:=0,1
	return func() int {
		a,b=b,a+b
		return a
	}
}

//为函数实现接口
func printFileContents(reader io.Reader)  {
	scanner:=bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}
//定义一个类型  类型为函数
type intGen func() int

//函数可以作为返回值  也可以作为接收者  返回值类型int  如果出错 就是error
func (g intGen) Read(p []byte) (n int, err error) {
	next:=g()
	if next>10000{
		return 0,io.EOF
	}
	s:=fmt.Sprintf("%d\n",next)
	return strings.NewReader(s).Read(p)
}
