package mock
//结构体
type Retriever struct {
	Contents string
}
//实现
func (r Retriever)Get(url string)string  {
	return r.Contents
}
