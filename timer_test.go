package belajar_go_routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	waktu := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	waktu2 := <-waktu.C //EMG FUNGSI DR SANA, C maksudnya channel karna disimpen dulu TIMERNYA dalam C
	fmt.Println(waktu2)

}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	waktu2 := <-channel
	fmt.Println(waktu2)
}

var g = sync.WaitGroup{}

func Panggil(nilai string) string {
	a := nilai
	g.Done()
	return "Woyyy" + a
}

func Waktu() {
	fmt.Println(time.Now())
}

func TestAfterFunc(t *testing.T) {

	fmt.Println(time.Now())
	g.Add(1)

	time.AfterFunc(5*time.Second, func() {
		Waktu()
		g.Done()
	})

	g.Wait()
	fmt.Println("Selesai")
}
