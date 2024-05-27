package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	new := [...]int{1, 2, 3, 4, 5, 6}
	blank := [10]int{1, 2, 3}
	fmt.Println("array of numbers:", numbers)

	numbers[1] = 20
	fmt.Println("array of numbers:", numbers)
	fmt.Println("array of numbers:", new)
	fmt.Println("array of numbers:", blank)
}
