package main

import "fmt"

func main() {

	type NoKTP string

	var ktpFeri NoKTP = "12333333333333333"

	var contoh string = "22222222222222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpFeri)
	fmt.Println(contohKTP)
}
