package main

import "fmt"

func main() {
	var input int // Variabel untuk nyimpen input dari user
	fmt.Print("Input	: ") // Minta user untuk input
	fmt.Scan(&input) // Baca input dari user dan simpan ke variabel input

	// Cek apakah inputnya sama dengan 42
	if input == 42 {
		fmt.Println("Output	: Hello Universe") // Jika iya, cetak pesan khusus
	} else {
		// Jika bukan 42, cetak input yang dimasukkan
		fmt.Println("Output	:", input)
	}
}
