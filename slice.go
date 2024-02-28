package main

import "fmt"

func main() {

	//Inisialisasi Slice
	var slice = []string{"apple", "orange", "banana", "grape"}
	fmt.Println(slice[0])

	// Hubungan Slice Dengan Array & Operasi Slice
	var slice1 = []string{"apple", "orange", "banana"}
	var newSlice = slice1[0:3]
	fmt.Println(newSlice) // [apple orange]

	// Slice Merupakan Tipe Data Reference

	var slice3 = []string{"apple", "grape", "banana", "melon"}

	var aSlice = slice3[0:3]
	var bSlice = slice3[1:4]

	var aaSlice = aSlice[1:2]
	var bbSlice = bSlice[0:1]

	fmt.Println(slice3)
	fmt.Println(aSlice)
	fmt.Println(bSlice)
	fmt.Println(aaSlice)
	fmt.Println(bbSlice)

	// Buah "grape" diubah menjadi "pinnaple"

	bbSlice[0] = "pinnaple"

	fmt.Println(slice3)
	fmt.Println(aSlice)
	fmt.Println(bSlice)
	fmt.Println(aaSlice)
	fmt.Println(bbSlice)

	// Fungsi len()
	var slice4 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(slice4))

	//Fungsi cap()

}

/*
	detail type data slice

	type data slice memiliki 3 data yaitu pointer, length dan capacity
	pointer adalah penunjuk data pertama di array para slice
	Length adalah panjang dari slice
	Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity
*/
