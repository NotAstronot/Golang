package main

import "fmt"

func main() {
	var a = 50
	var b = 3
	var c = a / b
	fmt.Println(c)

	var i = 10
	i += 10
	i++
	fmt.Println(i)

	var j = 20
	j++
	fmt.Println(j)

	// Jika dikurangkan
	j--
	fmt.Println(j)

	var value = (((2+6)%3)*4 - 2) / 3

	fmt.Println(value)
}

/* (+) adalah penjumlahan
(-) adalah pengurangan
(*) adalah perkalian
(/) adalah pembagian
(%) adalah sisa pembagian atau modulo*/
