package belajar_go_routine

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTicker(t *testing.T) {
	ticker := time.NewTicker(3 * time.Second)

	go func() {
		time.Sleep(6 * time.Second)
		ticker.Stop()
	}()

	for i := range ticker.C {
		fmt.Println(i) //print waktu setiap 3 detik
	}
}

func TestTicker(t *testing.T) {
	channel := time.Tick(2 * time.Second)

	for ticker := range channel {
		fmt.Println(ticker)
	}
}
