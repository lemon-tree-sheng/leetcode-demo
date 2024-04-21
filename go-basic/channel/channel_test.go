package slice

import (
	"sync"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	a := make(chan int, 10)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go produce(a)

	go recv(a)

	wg.Wait()
}

func recv(a chan int) {
	for {
		x := <-a
		println(x)
	}
}

func produce(a chan int) {
	for i := 0; i < 100; i++ {
		a <- i
		time.Sleep(1)
	}
}
