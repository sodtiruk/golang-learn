package main

import "fmt"

//create function
func hello() {
	fmt.Println("hello world")
}

func plus(number1 int, number2 int) int {
	return number1 + number2
}

func main() {
	//call function
	hello()
	fmt.Println(plus(1, 2))

	result := plus(2, 3)
	fmt.Println("result = ", result)

}
