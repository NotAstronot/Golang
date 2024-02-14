package main

import "fmt"

func main() {
	var name1 = "Feri"
	var name2 = "Riski"

	var result bool = name1 == name2

	fmt.Println(result)

	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Println("nilai", value, isEqual)

}
