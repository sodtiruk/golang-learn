package main

import "fmt"

//map variable same as dictionary in python

var product = make(map[string]int)

func main() {

	fmt.Println("product: ", product)

	//add
	product["Macbook"] = 40000
	product["Iphone"] = 30000
	product["Ipad"] = 25000
	fmt.Println("product: ", product)

	//delete
	delete(product, "Ipad")
	fmt.Println("product: ", product)

	//update
	product["Iphone"] = 13
	fmt.Println("product: ", product)

	//
	//short format declare variable
	value1 := product["Iphone"]
	fmt.Println("value1 = ", value1)

	//create map and then initiale value
	courseName := map[string]string{"101": "Java", "102": "Python"}
	fmt.Println(courseName)

}
