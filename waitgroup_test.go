package belajar_go_routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(grup *sync.WaitGroup) {
	defer grup.Done() //jika sudah selesai, maka done ini akan mengirim sinyal ke wait bahwa proses fungsi ini telah selesai dilakukan

	grup.Add(1) //kenapa addnya disini yak?!
	fmt.Println("Fira")

	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	grupMakMak := &sync.WaitGroup{}
	x := 0
	for i := 0; i < 100000; i++ {
		go RunAsynchronous(grupMakMak)
		x = x + 1
	}

	grupMakMak.Wait()

	fmt.Println(x) //karena dia gamasuk dalam goroutine
	fmt.Println("Selesai")
}
