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
		fmt.Println("tidak lulus:", point, "\n")
	}

	// Variabel Temporary Pada if - else
	var nilai = 8559.0

	if nilai := nilai / 100; nilai >= 100 {
		fmt.Println("perfect", nilai)
	} else if nilai >= 70 {
		fmt.Println("good", nilai, "\n")
	} else {
		fmt.Println("not bad", nilai)
	}

	// Switch - case
	var pointKedua = 8

	switch pointKedua {
	case 8:
		fmt.Println("perfect \n")
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
			fmt.Println("you can be better! \n")
		}
	}

	// Switch Dengan Gaya if - else
	var angka = 6

	switch {
	case angka == 8:
		fmt.Println("perfect")
	case (angka < 8) && (angka > 3):
		fmt.Println("awesome \n")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can need learn more")
		}
	}

	// Penggunaan Keyword fallthrough Dalam switch
	var dapatAngka = 8

	switch {
	case dapatAngka == 8:
		fmt.Println("perfect \n")
	case (dapatAngka < 8) && (dapatAngka > 3):
		fmt.Println("awesome")
		fallthrough
	case dapatAngka < 5:
		fmt.Println("you can learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can learn more")
		}
	}

	// Seleksi Kondisi Bersarang
	var pointPertama = 0

	if pointPertama > 7 {
		switch pointPertama {
		case 4:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if pointPertama == 5 {
			fmt.Println("not bad!")
		} else if pointPertama == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if pointPertama == 0 {
				fmt.Println("try harder")
			}
		}
	}
}
