package main

import (
	"fmt"
	"net/http"
)

func main() {
	////开启一个web server  http://localhost:8888
	//http.HandleFunc("/",
	//	    //  参数名在前  参数类型在后
	//	func(writer http.ResponseWriter,
	//		//指针
	//		request *http.Request) {
	//	fmt.Fprint(writer,"<h1>hello licslan</h1>")
	//})
	//http.ListenAndServe(":8888",nil)


	//开启一个web server 这个时候传个参数进来  http://localhost:8888/?name=hwl
	http.HandleFunc("/",
		//  参数名在前  参数类型在后
		func(writer http.ResponseWriter,
		//指针
			request *http.Request) {
			fmt.Fprintf(writer,"<h1>hello licslan %s</h1>",request.FormValue("name"))
		})
	http.ListenAndServe(":8888",nil)
}
