package main

import "fmt"

func main() {

	// Perulangan Menggunakan Keyword for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	//Penggunaan Keyword for Dengan Argumen Hanya Kondisi
	var i = 0

	for i < 5 {
		fmt.Println("Nilai", i)
		i++
	}

	//Penggunaan Keyword for Tanpa Argumen

	var j = 0

	for {
		fmt.Println("Point", j)
		j++
		if j == 5 {
			break
		}
	}

	// Penggunaan Keyword for - range

	var xs = "123" // string
	for ukuran, v := range xs {
		fmt.Println("Index=", ukuran, "Value=", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("value=", v)
	}

	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("valuee=", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Println("Key=", k, "valueee=", v)
	}

	// boleh juga baik k dan atau v nya diabaikan

	for range kvs {
		fmt.Println("Done")
	}

	// Penggunaan Keyword break & continue
	for a := 1; a <= 10; a++ {
		if a%2 == 1 {
			continue
		}
		if a > 8 {
			break
		}

		fmt.Println("Hasil", a)
	}

	// Perulangan Bersarang

	for b := 0; b < 5; b++ {
		for c := b; c < 5; c++ {
			fmt.Println(c, "")
		}
		fmt.Println()
	}

	// Pemanfaatan Label Dalam Perulangan
outerLoop:
	for d := 0; d < 5; d++ {
		for e := 0; e < 5; e++ {
			if d == 3 {
				break outerLoop
			}
			fmt.Println("matriks [", d, "][", e, "]", "\n")
		}
	}

}
