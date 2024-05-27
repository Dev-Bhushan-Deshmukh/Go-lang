package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("crud api")
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {

		fmt.Println("err::", err)
		return

	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println("The error is:",response.Status)

	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("err occured:", err)
		return
	}
	fmt.Println("data:", string(data))
}
