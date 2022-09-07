package belajar_go_routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var counter = 0

func OnlyOnce() {
	counter = counter + 1

}

// INI SALAH HASILNYA MASIH GA NENTU
func TestOnce(t *testing.T) {
	wg := sync.WaitGroup{}
	//var once sync.Once

	for h := 0; h < 100; h++ {
		wg.Add(1) //kalo dihilang sync lainnya dia rada tanda kuning , knapa ya??
		go func() {
			//once.Do(OnlyOnce)

			OnlyOnce()
			fmt.Println(counter)
			wg.Done()
		}()
	}

	wg.Wait()
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
	fmt.Println("Counter", counter)

}
