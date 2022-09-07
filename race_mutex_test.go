package belajar_go_routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for h := 1; h <= 100; h++ {
				x = x + 1 // x merupakan variabel yg disharing, diakses oleh beberapa goroutine
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter", x)
	// hasilnya seharusnya 1000 * 100, namun hasilnya tidak akan tepat karena menggunakan goroutine(parallel) sehingga ada perhitungan yang sama menyebabkan hasilnya tidak tepat(balapan antara 1 core dgn yg lain)
}

func TestWithMutex(t *testing.T) {
	z := 0
	var mutx sync.Mutex //definisikan variabel mutex

	for a := 1; a <= 500; a++ {
		go func() {
			for b := 1; b <= 100; b++ {
				mutx.Lock() //mengunci
				z = z + 1
				mutx.Unlock() //membuka, jadi tidak akan ada bentrok perhitungan yang sama, karna hanya 1 goroutine yg bisa mengakses
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Nilai dari z = ", z) //hasilnya sesuai dengan yg diinginkan

}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(nilaiTambah int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + nilaiTambah
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int { //struct method lebih baik menggunakan huruf kapital
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

// MARI KITA BUAT TESTNYA!!!!!!
func TestReadWriteMutex(t *testing.T) {
	//pertama mari kita buat variabel utk mendefinisikan structnya
	akun := BankAccount{}

	//mari kita test
	for i := 0; i < 100; i++ {
		go func() {
			for c := 0; c < 100; c++ {
				akun.AddBalance(1)
				fmt.Println(akun.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance : ", akun.GetBalance())

}

type UserBalance struct {
	Mut     sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Unlock() {
	user.Mut.Unlock()
}

func (user *UserBalance) Lock() {
	user.Mut.Lock()
}

func (user *UserBalance) Change(nilai int) {
	user.Balance = user.Balance + nilai
}

//pola transfernya adalah user 1 mengirim uang ke user2

func Transfer(user1 *UserBalance, user2 *UserBalance, nilai int) {
	user1.Lock()
	fmt.Println("Lock User 1", user1.Name)
	user1.Change(-nilai)

	time.Sleep(1 * time.Second) //biar ada jeda

	user2.Lock()
	fmt.Println("Lock User 2", user2.Name)
	user2.Change(nilai)
	time.Sleep(1 * time.Second) //biar ada jeda

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Habie",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Radit",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 50000) //kalau hanya satu tidak masalah, namun jika dua akan terjadi deadlock, CASE TERBURUKNYA jika user1 dan user2 melakukan transfer satu sama lain
	go Transfer(&user2, &user1, 20000) //masing2 sama2 ngelock user1 sehingga user2(yg merupakan user1 di panggilan fungsi lain) tidak bisa diakses karna masih dilock
	time.Sleep(5 * time.Second)        //penting untuk berhenti sesaat, jika ini ditempatkan dipaling bawah, fmtPrintln akan tampil duluan sebelum goroutine selesai diproses

	fmt.Println(user1.Balance)
	fmt.Println(user2.Balance)

}
