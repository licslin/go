package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/list/",
		func(writer http.ResponseWriter,
			request *http.Request) {
			path :=request.URL.Path[len("/list/"):]  // /list/fib.txt
			file,err:=os.Open(path)
			/*if err!=nil{
				panic(err)
			}*/
			if err!=nil{
				http.Error(writer,err.Error(),http.StatusInternalServerError)
				return
			}
			defer file.Close()
			//byte[]  error
			all, err := ioutil.ReadAll(file)
			if err!=nil{
				panic(err)
			}

			//将byte[]写到ResponseWriter里面去
			writer.Write(all)
		})

	fmt.Println("the server is running at port 8888----------")
	//http server监听端口 8888
	err:= http.ListenAndServe(":8888", nil)
	if err!=nil{
		panic(err)
	}

}
