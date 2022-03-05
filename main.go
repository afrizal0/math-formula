package main

import (
	"fmt"
	"math-formula/aritmatika"
)

func main() {
	fmt.Println("Masukkan Formula Matematika yang ingin dioperasikan")
	var formula string
	fmt.Scanf("%s\n", &formula)
	var hasil int
	if formula == "aritmatika" {
		var number1 int
		var number2 int
		var operator string
		fmt.Println("Masukkan angka yang ingin dioperasikan")
		fmt.Scanf("%d\n", &number1)
		fmt.Println("Masukkan angka yang ingin dioperasikan")
		fmt.Scanf("%d\n", &number2)
		fmt.Println("Masukkan operator")
		fmt.Scanf("%s\n", &operator)
		switch operator {
		case "+":
			hasil = aritmatika.Penjumlahan(number1, number2)

		case "-":
			hasil = aritmatika.Pengurangan(number1, number2)

		case "*":
			hasil = aritmatika.Perkalian(number1, number2)

		case "/":
			hasil = aritmatika.Pembagian(number1, number2)

		case "*%":
			hasil = aritmatika.Modulus(number1, number2)
		}

	}
	fmt.Printf("Hasil dari formula %s tersebut adalah %d\n", formula, hasil)
}
