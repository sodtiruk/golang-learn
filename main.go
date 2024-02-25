package main

import "fmt"

// how create variable
var numberInt int = 2
var numbersInt1, numbersInt2 int = 1000, 2000
var msg string = "message"

//create array
var productName [4]string
var price [4]float32

func main() {
	//day 1

	//output hello world
	fmt.Println("hello world")

	//output number which you create variable
	fmt.Printf("%d\n", numberInt)
	fmt.Println(numbersInt1, numbersInt2)
	fmt.Println(msg)

	fmt.Println(numbersInt1 + numbersInt2)

	//casting variable
	numberfloat := 25.4 //create float number
	fmt.Println(numberfloat)
	fmt.Println(int(numberfloat) + numbersInt2)

	//concate string
	fmt.Println(msg+"World = ", numbersInt2)

	// call array
	productName[0] = "Macbook"
	productName[1] = "Ipad"
	productName[2] = "Iphone"
	productName[3] = "Airpods"

	price := [4]float32{2000, 3000, 4000, 50}

	fmt.Println(productName)
	fmt.Println(price)

}
