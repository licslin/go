package mock

import "fmt"

//结构体
type Retriever struct {
	Contents string
}



func (r Retriever) Post(url string, form map[string]string) string {
	r.Contents="licslan"
	//panic("implement me")
	return "yes"
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever :{Contents=%s}",r.Contents)
	//panic("implement me")
}

func (r Retriever) Connect(host string) {
	panic("implement me")
}

//实现
func (r Retriever)Get(url string)string  {
	return r.Contents
}
