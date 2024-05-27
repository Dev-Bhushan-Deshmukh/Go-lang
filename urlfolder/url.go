package main

import (
	"fmt"
	"net/url"
)

func main() {
	myurl := "https://jsonplaceholder.typicode.com/posts/1"

	parsedurl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("url parsed", parsedurl)

	fmt.Println(parsedurl.Scheme)

	fmt.Println(parsedurl.Fragment)

	fmt.Println(parsedurl.ForceQuery)

	fmt.Println(parsedurl.Path)

	fmt.Println(parsedurl.RawPath)

	fmt.Println(parsedurl.RawQuery)

	fmt.Println(parsedurl.RawFragment)

	fmt.Println(parsedurl.User)

}
