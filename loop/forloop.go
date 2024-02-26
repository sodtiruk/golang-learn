package main

import "fmt"

func main() {

	// while loop
	i := 0
	for i < 30 {
		fmt.Println("i = ", i)
		i += 1
	}

	// for loop
	count := 30
	for j := 0; j < count; j++ {
		fmt.Println("j = ", j)
	}

	// infinity loop
	for {
		fmt.Println("Hello Sutthirak")
	}

}
