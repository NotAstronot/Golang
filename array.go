package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Feri"
	names[1] = "Riski"
	names[2] = "Muh"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// array secara langsung

	var values = [...]int{
		90, 80, 70, 80, 20}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}

/*
function Array

	len(array) 					untuk mendapatkan panjang array
	array[index] 				untuk mendapatkan data di posisi index
	array[index] = value 		untuk mengubah data di posisi index
*/
