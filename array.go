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

	// Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen

	var values = [...]int{
		90, 80, 70, 80, 20}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))

	//Inisialisasi Nilai Awal Array

	var fruits = [4]string{"apple", "orange", "banana", "melon"}

	fmt.Println("Jumlah Element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// Array Multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 4}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	//Perulangan Elemen Array Menggunakan Keyword for
	var buah = [4]string{"apple", "orange", "banana", "melon"}

	for i := 0; i < len(buah); i++ {
		fmt.Println("elements", i, buah[i])
	}

	//Perulangan Elemen Array Menggunakan Keyword for - range
	var buah2 = [4]string{"apple", "orange", "banana", "semangka"}

	for i, buah2 := range buah2 {
		fmt.Println("elements", i, buah2)
	}

	//Penggunaan Variabel Underscore _ Dalam for - range
	var buah3 = [4]string{"apple", "orange", "banana", "semangka"}

	for _, buah4 := range buah3 {
		fmt.Println("nama", buah4)
	}
}

/*
function Array

	len(array) 					untuk mendapatkan panjang array
	array[index] 				untuk mendapatkan data di posisi index
	array[index] = value 		untuk mengubah data di posisi index
*/
