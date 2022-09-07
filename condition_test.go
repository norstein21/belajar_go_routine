package belajar_go_routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var kondisi = sync.NewCond(&locker)
var grup = sync.WaitGroup{}

func WaitCondition(value int) {
	//data yang menunggu diambil
	kondisi.L.Lock()
	kondisi.Wait()
	fmt.Println("Done", value)
	kondisi.L.Unlock()
	grup.Done()

}

func TestKondisi(t *testing.T) {
	for i := 0; i < 10; i++ {
		grup.Add(1)
		go WaitCondition(i)
	}

	for a := 0; a < 10; a++ {
		time.Sleep(1 * time.Second)
		kondisi.Signal()
	}

	grup.Wait() //menunggu dari done

}
