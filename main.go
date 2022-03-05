package main

import (
	"fmt"
	"math-formula/aritmatika"
	bangunruang "math-formula/bangun-ruang"
)

func main() {
	fmt.Println("Masukkan Formula Matematika yang ingin dioperasikan")
	var formula string
	fmt.Scanf("%s\n", &formula)
	switch formula {
	case "aritmatika":
		var number1 int
		var number2 int
		var operator string
		var hasil int

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
		fmt.Printf("Hasil dari formula %s tersebut adalah %d\n", formula, hasil)

	case "bangunruang":
		var bangun string
		fmt.Println("Masukkan bangun ruang")
		fmt.Scanf("%s\n", &bangun)

		switch bangun {
		case "kubus":
			var sisi int
			fmt.Println("Masukkan sisi")
			fmt.Scanf("%d\n", &sisi)
			fmt.Printf("Volume Kubus : %d", bangunruang.Volume_Kubus(sisi))
			fmt.Printf("Volume Kubus : %d", bangunruang.Luas_Permukaan(sisi))
			fmt.Printf("Volume Kubus : %d", bangunruang.Keliling_Kubus(sisi))
			fmt.Printf("Volume Kubus : %d", bangunruang.Luas_Sisi(sisi))

		}

	}
}
