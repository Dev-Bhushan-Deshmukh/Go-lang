package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {

		fmt.Println("error occured", err)
		return

	}
	defer res.Body.Close()
	// fmt.Println("resonse", res)
data,err:=ioutil.ReadAll(res.Body)
if err !=nil{

	fmt.Println("response error:",err)
return
}
fmt.Println("response:",string(data))
}
