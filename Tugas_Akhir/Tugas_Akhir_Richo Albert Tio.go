package main

import "fmt"

// program main program utama
func main() {
	// Mendeklarasikan variabel
	var a = 5
	var b = 5
	fmt.Println("Hasil Aritmatika nilai pertama 5, dan nilai kedua 5")
	// Memanggil fungsi serta parameternya
	jumlah(a, b)
	kurang(a, b)
	kali(a, b)
	bagi(a, b)
	modulus(a, b)
}

// fungsi penambahan 
func jumlah(a int, b int) {
	var hasil = a + b
	fmt.Println("Penjumlahan", a, "+", b, "=", hasil)
}

// fungsi pengurangan
func kurang(a int, b int) {
	var hasil = a - b
	fmt.Println("Pengurangan", a, "-", b, "=", hasil)
}

// fungsi perkalian
func kali(a int, b int) {
	var hasil = a * b
	fmt.Println("Perkalian", a, "*", b, "=", hasil)
}

// fungsi Pembagian
func bagi(a int, b int) {
	var hasil = a / b
	fmt.Println("Pembagian", a, "/", b, "=", hasil)
}

// fungsi modulus
func modulus(a int, b int) {
	var hasil = a % b
	fmt.Println("Sisa Hasil Bagi Pembagian MODULUS", a, "%", b, "=", hasil)
}
