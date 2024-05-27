package main

import "fmt"

type person struct {
	name        string
	age         int
	designation string
	salary      float32
}

func main() {
	var person1 person

	person1.name = "bhushan"
	person1.age = 27
	person1.designation = "senior conultant"
	person1.salary = 300000.00

	fmt.Println(person1.name)
	fmt.Println(person1.age)
	fmt.Println(person1.designation)
	fmt.Println(person1.salary)
	fmt.Println("---------")
	fmt.Println(person1)
}
