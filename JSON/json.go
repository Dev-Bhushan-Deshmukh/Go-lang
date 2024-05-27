package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	person1 := Person{Name: "bhushan", Age: 27, City: "Akola"}
	fmt.Println("Person data:", person1)

	jsondata, err := json.Marshal(person1)
	if err != nil {

		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(jsondata))

var person2 Person
	erro:=json.Unmarshal(jsondata,&person2)
	if erro!=nil{

fmt.Println("error occured:",erro)

	}

	fmt.Println("person un parsed data:",person2)
}
