package main

import "bufio"
import "fmt"
import "os"
import "strings"

// Fungsi ini untuk membalikkan string
func reverse(s string) string {
	// Ubah string jadi slice rune biar bisa ngedukung karakter non-ASCII
	runes := []rune(s) 
	// Balikkan runes dari depan ke belakang
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) // Kembalikan hasilnya sebagai string
}

func main() {
	reader := bufio.NewReader(os.Stdin) // Buat reader untuk input

	fmt.Print("Input minimal 3 kata: ")
	input, _ := reader.ReadString('\n') // Baca input dari user
	input = strings.TrimSpace(input)      // Hapus spasi di awal dan akhir

	// Pisahkan input jadi kata-kata
	words := strings.Fields(input)

	// Cek apakah jumlah katanya minimal 3
	if len(words) < 3 {
		fmt.Println("Input harus minimal 3 kata") // Kasih tau kalau kurang
		return
	}

	var reversed_words []string // Slice untuk nyimpen kata-kata yang udah dibalik

	// Balikkan setiap kata dan tambahin ke slice
	for _, word := range words {
		reversed_words = append(reversed_words, reverse(word))
	}

	// Cetak kata-kata yang udah dibalik dalam satu baris
	fmt.Println(strings.Join(reversed_words, " "))
}
