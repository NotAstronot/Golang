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
	var slice5 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	var aSlice5 = slice5[0:3]
	fmt.Println(len(aSlice5))
	fmt.Println(cap(aSlice5))

	var bSlice5 = slice5[1:4]
	fmt.Println(len(bSlice5))
	fmt.Println(cap(bSlice5))

	// Fungsi append()
	var slice6 = []string{"apple", "grape", "banana"}
	var cSlice = append(slice6, "pepaya")

	fmt.Println(slice6)
	fmt.Println(cSlice)

	// agar lebih jelas perhatikan contoh dibawah ini

	var slice7 = []string{"apple", "grape", "banana"}
	var dSlice = slice7[0:2]

	fmt.Println(cap(slice7))
	fmt.Println(len(dSlice))

	fmt.Println(slice7)
	fmt.Println(dSlice)

	var eSlice = append(dSlice, "pepaya")

	fmt.Println(slice7)
	fmt.Println(dSlice)
	fmt.Println(eSlice)

	// Fungsi copy()
	// Fungsi copy() digunakan untuk men-copy elements slice pada src (parameter ke-2), ke dst (parameter pertama).

	// copy(dst, src)

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

}

/*
	detail type data slice

	type data slice memiliki 3 data yaitu pointer, length dan capacity
	pointer adalah penunjuk data pertama di array para slice
	Length adalah panjang dari slice
	Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity
*/
