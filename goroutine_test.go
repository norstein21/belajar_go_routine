package belajar_go_routine

import (
	"fmt"
	"testing"
	"time"
)

func MenulisKedua() {
	fmt.Println("Ini adalah tulisan yang tidak panjang")
}

func MenulisRia() {
	fmt.Println("Ini adalah tulisan yang tidak panjang-panjang amat")
}

func TestGoRoutine(t *testing.T) {
	//di jalankan secara asyncronous, dilakukan secara paralel sehingga mendahulukan mana yan lebih dulu selesai
	go MenulisRia() //menggunakan go routine, tidak disarankan untuk yg mengembalikan return value
	go MenulisKedua()
	time.Sleep(1 * time.Second)

}

func CobaTampil(angka int) {
	fmt.Println("Display", angka)
}

func TestCobaTampil(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go CobaTampil(i) // kalau diperhatakan angka yang keluar tidak berurutan, walaupun go mendukung concurrency, namun akan menggunakan parallel jika prosessor kita memiliki lebih  dari  1 core
	}

	//time.Sleep(1 * time.Second)
}
