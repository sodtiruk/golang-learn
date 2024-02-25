package main

import "fmt"

func main() {

	//validate grade
	// > 80 A
	// > 70 B
	// > 60 C
	// > 50 D
	//otherwise F
	fmt.Println("grade calculator")
	var grade int
	fmt.Scanf("%d", &grade)

	if grade > 80 {
		fmt.Println("A")
	} else if grade > 70 {
		fmt.Println("B")
	} else if grade > 60 {
		fmt.Println("C")
	} else if grade > 50 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

}
