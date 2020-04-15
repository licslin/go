package fib
func Fibonacci() intGen {
	a,b:=0,1
	return func() int {
		a,b=b,a+b
		return a
	}
}

//定义一个类型  类型为函数
type intGen func() int
