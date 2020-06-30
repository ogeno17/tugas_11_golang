package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input1, input2 string

	fmt.Print("Masukan angka ke-1: ")
	fmt.Scanln(&input1)

	fmt.Print("Masukan angka ke-2: ")
	fmt.Scanln(&input2)

	angka1, err := strconv.Atoi(input1)
	if err != nil {
		panic(err.Error())
	}

	angka2, err := strconv.Atoi(input2)
	if err != nil {
		panic(err.Error())
	}

	penjumlahan(angka1, angka2)
	pengurangan(angka1, angka2)
	perkalian(angka1, angka2)
}

func penjumlahan(angka1, angka2 int) {
	hasil := angka1 + angka2
	fmt.Println("hasil penjumlahan: ", hasil)
}

func pengurangan(angka1, angka2 int) {
	hasil := angka1 - angka2
	fmt.Println("hasil pengurangan: ", hasil)
}

func perkalian(angka1, angka2 int) {
	hasil := angka1 * angka2
	fmt.Println("hasil perkalian: ", hasil)
}
