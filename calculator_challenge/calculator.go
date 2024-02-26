package main

import "fmt"

func main() {

	var numberToCalculator int
	fmt.Scan(&numberToCalculator)

	score := 0
	var keepOperator string
	for i := 1; i <= numberToCalculator; i++ {

		var number int
		fmt.Printf("Number %d : ", i)
		fmt.Scan(&number)
		fmt.Println("")

		if i == 1 {
			score += number

			fmt.Printf("Operator : ")
			var operator string
			fmt.Scan(&operator)
			fmt.Println("")

			keepOperator = operator

		} else {

			switch keepOperator {
			case "/":
				score /= number
			case "+":
				score += number
			case "-":
				score -= number
			case "*":
				score *= number
			}

			if i != numberToCalculator {
				fmt.Printf("Operator : ")
				var operator string
				fmt.Scan(&operator)
				fmt.Println("")

				keepOperator = operator
			}
		}
	}

	fmt.Println("score : ", score)

}
