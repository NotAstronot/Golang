package main

import "fmt"

func main() {

	var John = false
	var Wick = true

	var johnAndwick = John && Wick
	fmt.Println("John && Wick :", johnAndwick)

	var johnOrwick = John || Wick
	fmt.Println("John || Wick :", johnOrwick)

	var JohnReverse = !John
	fmt.Println("John :", JohnReverse)
}
