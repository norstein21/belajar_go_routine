package belajar_go_routine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for b := 0; b < 1000; b++ {
		group.Add(1)
		go func() {
			for c := 0; c < 100; c++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()

	fmt.Println(x)
}
