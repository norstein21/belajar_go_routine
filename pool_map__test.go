package belajar_go_routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "kosong" //merubah default data kosong dari nil, menjadi string yg kita inginkan
		},
	}

	pool.Put("Habie")
	pool.Put("Purwokusumo")
	pool.Put("Putrapandowo")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)

		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Selesai")
}

func AddToMap(data *sync.Map, nilai int, wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	data.Store(nilai, nilai)

}

func TestMap(t *testing.T) {
	dt := &sync.Map{}
	grup := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			AddToMap(dt, i, grup)
		}()
	}

	grup.Wait()

	dt.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
