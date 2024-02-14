package main

import "fmt"

func main() {

	//Seleksi Kondisi Menggunakan Keyword if, else if, & else
	var point = 5

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Println("tidak lulus: ", point)
	}

	// Variabel Temporary Pada if - else
	var nilai = 8559.0

	if nilai := nilai / 100; nilai >= 100 {
		fmt.Println("perfect", nilai)
	} else if nilai >= 70 {
		fmt.Println("good", nilai)
	} else {
		fmt.Println("not bad", nilai)
	}
}
