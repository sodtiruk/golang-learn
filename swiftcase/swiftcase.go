package main

import "fmt"

func main() {
	input := 2
	switch input {
	case 1:
		fmt.Println("i am one")
	case 2:
		fmt.Println("i am two")
	case 3:
		fmt.Println("i am three")
	default:
		fmt.Println("default case")
	}

	var str string
	fmt.Scan(&str)
	switch str {
	case "blue":
		fmt.Println("#0000FF")
	case "green":
		fmt.Println("#008000")
	case "pink":
		fmt.Println("#FFC0CB")
	case "yellow":
		fmt.Println("#FFFF00")
	default:
		fmt.Println("No color")

	}

}
