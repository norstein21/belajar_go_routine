package belajar_go_routine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan int)
	defer close(channel)

	//bikin anonymous function
	go func() { //penting sekali menjadikan nya go routine, kalau tidak akan terjadi error
		time.Sleep(2 * time.Second)
		channel <- 12
		fmt.Println("Data Channel sudah berhasil diterima")
	}()

	//var data int
	data := <-channel //lebih simple dan fungsinya sama
	fmt.Println(data)
	time.Sleep(5 * time.Second)

}

func GoChannel(isiChannel chan string) { //tidak perlu pake pointer di *chan
	time.Sleep(2 * time.Second)
	isiChannel <- "Habie Purwokusumo Putrapandowo"
}

func TestChannelWithParameter(t *testing.T) {
	kotakChannel := make(chan string)
	defer close(kotakChannel)

	go GoChannel(kotakChannel)

	penerimaData := <-kotakChannel

	fmt.Println(penerimaData)
	time.Sleep(5 * time.Second)
}

func ChannelIn(isiChannel chan<- string) { //chan<- hanya untuk data yang masuk ke dalam channel
	time.Sleep(2 * time.Second)
	isiChannel <- "Ini merupakan data yang dipindahkan dari channel in"
}

func ChannelOut(ambilChannel <-chan string) { // <-chan hanya untuk data yang diambil/dikirim dr channel
	isiData := <-ambilChannel
	fmt.Println(isiData)
}

func TestChannelMasukKeluar(t *testing.T) {
	chann := make(chan string)
	defer close(chann)

	go ChannelIn(chann)
	go ChannelOut(chann)

	time.Sleep(5 * time.Second)

}

func TestBufferedChann(t *testing.T) {
	fmt.Println("****MULAI****")
	cha := make(chan string, 4) // 4 disini maksudnya channel memiliki kapasitas sebesar 4 untuk menampung data antrian
	defer close(cha)

	go func() {
		time.Sleep(2 * time.Second)
		cha <- "Naruto"
		cha <- "Sasuke"
		cha <- "Sakura"
		cha <- "Inzaghi"
		cha <- "Nagatomo" //walaupun ada 5 tidak masalah karna ini menggunakan go routine jd data yang terkirim akan memberi ruang kosong pada buffered(antrian) sehingga data ke-5 dapat masuk
		cha <- "Kaka"
		cha <- "Luffy"
		cha <- "Ronaldo" //maksimal 8 karena kapasitas nya 4, jika semua data sudah keluar, maka 4 lainnya bisa mengisi ruang kosong dalam antrian(buffered)
	}()

	go func() {
		data_1 := <-cha
		data_2 := <-cha
		data_3 := <-cha
		data_4 := <-cha
		data_5 := <-cha
		data_6 := <-cha
		data_7 := <-cha
		data_8 := <-cha //ternyata buffered channel bisa digunakan 2x lipat dari kapasitas yang kita tentukan

		fmt.Println(data_1)
		fmt.Println(data_2)
		fmt.Println(data_3)
		fmt.Println(data_4)
		fmt.Println(data_5)
		fmt.Println(data_6)
		fmt.Println(data_7)
		fmt.Println(data_8)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("****SELESAI****")
}

func TestWithLoop(t *testing.T) {
	c := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			c <- "Ini perhitungan ke-" + strconv.Itoa(i) //ini dari int dikonversikan menjadi string
		}
		close(c)
	}()

	for dataLoop := range c {
		fmt.Println(dataLoop)
	}
}

func TestWithSelect(t *testing.T) {
	chan1 := make(chan string)
	chan2 := make(chan string)
	defer close(chan1)
	defer close(chan2)

	go GoChannel(chan1)
	go GoChannel(chan2)

	counter := 0
	angka := 0
	for {
		select { //kalau selectnya hanya 1 maka yang keluar juga hanya 1, solusiya bisa menambahkan select lagi atau menggunakan looping
		case data := <-chan1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-chan2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			angka++
			fmt.Println(angka)
		}
		if counter == 2 {
			break
		} //code untuk menghentikan perulangan
	}

}
