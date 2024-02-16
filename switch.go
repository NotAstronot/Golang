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
		fmt.Println("tidak lulus:", point)
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

	// Switch - case
	var pointKedua = 8

	switch pointKedua {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Kurung Kurawal Pada Keyword case & default
	var keyword = 3

	switch keyword {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// Switch Dengan Gaya if - else
	var angka = 6

	switch {
	case angka == 8:
		fmt.Println("perfect")
	case (angka < 8) && (angka > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can need learn more")
		}
	}

	// Penggunaan Keyword fallthrough Dalam switch

}
