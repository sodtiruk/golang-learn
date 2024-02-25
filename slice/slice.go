package main

import "fmt"

func main() {

	//create slice variable, slice variable same as list in python
	var courseName []string
	courseName = []string{"Java", "Python"}
	fmt.Println(courseName)
	courseName = append(courseName, "C", "Javascript", "Html", "css", "ruby", "rust")
	fmt.Println(courseName)

	courseWeb := courseName[4:7]
	fmt.Println(courseWeb)

	courseWeb = courseName[:4]
	fmt.Println(courseWeb)

}
