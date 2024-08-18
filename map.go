package main

import "fmt"

// Penggunaan Map
func main() {

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["Januari"] = 50
	chicken["Februari"] = 40

	fmt.Println("Januari", chicken["Januari"])
	fmt.Println("Mei", chicken["Mei"])

}
