package main

import "fmt"

func myfunc() {
	fmt.Println("function called")

}

func arg(name string) {

	fmt.Println("my name is:", name)

}

func addition(num1, num2 int) int {
	var res int = num1 + num2
	return res
}

func multiply(num1, num2 int) (result1, result2 int) {
	result1 = num1 * 2
	result2 = num2 * 2
	return
}

func main() {
	myfunc()
	arg("bvd")
	res := addition(2, 4)
	fmt.Println("result:", res)
	a, b := multiply(2, 3)
	fmt.Println(a, b)
	c, _ := multiply(2, 3)
	fmt.Println(c)
}
