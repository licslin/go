package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

//结构体
type Retriever struct {
	UserAgent string
	TimeOut time.Duration
}

//Get implement from play.go  值接收者
func (r Retriever) Get(url string) string {
	//panic("implement me")
	resp, err := http.Get(url)
	if err!=nil{panic(err)}
	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err!=nil{panic(err)}
	return string(result)
}

