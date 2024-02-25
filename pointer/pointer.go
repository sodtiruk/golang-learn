package main

import "fmt"

//pointer
func setZero(number int) {
	number = 0
}

func zeroPointer(iPointer *int) {
	*iPointer = 0
}

func main() {
	i := 1
	fmt.Println("i = ", i)

	//cannot change
	setZero(i)
	fmt.Println("new i = ", i)

	zeroPointer(&i)
	fmt.Println("new i2 = ", i)
	fmt.Println("i value = ", i)
	fmt.Println("i address = ", &i)

}
