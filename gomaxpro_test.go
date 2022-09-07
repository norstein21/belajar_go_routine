package belajar_go_routine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomax(t *testing.T) {
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU = ", totalCPU)

	totalTHREAD := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread =", totalTHREAD)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total GoRoutine =", totalGoroutine)
}

func TestGomaxLebihJelas(t *testing.T) {
	grup := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		grup.Add(1)
		go func() {
			time.Sleep(2 * time.Second)
			fmt.Println(i)
			grup.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU = ", totalCPU)

	totalTHREAD := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread =", totalTHREAD)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total GoRoutine =", totalGoroutine)

	grup.Wait() //kalo waitnya di atas kode lain, tidak kelihatan berapa yg lg dieksus goroutine
}
