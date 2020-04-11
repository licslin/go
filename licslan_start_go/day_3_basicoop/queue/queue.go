package queue

//封装一个队列

//type Queue []int
//使其可以支持任何类型
type Queue []interface{}

/*func (q *Queue)Push(v int)  {
	*q=append(*q,v)
}*/
func (q *Queue)Push(v interface{})  {
	*q=append(*q,v)
}

/*func (q *Queue)Pop() int  {
	head:=(*q)[0]
	*q=(*q)[1:]
	return head
}*/
func (q *Queue)Pop() interface{}  {
	head:=(*q)[0]
	*q=(*q)[1:]
	return head
}

func (q *Queue)IsEmpty() bool  {
	return len(*q)==0
}

