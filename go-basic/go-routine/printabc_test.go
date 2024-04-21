package slice

import (
	"sync"
	"testing"
)

/**
不同 go-routine 循环打印 123
wait-group
协程池
*/

func TestFoo(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(3)

	ach := make(chan int, 1)
	bch := make(chan int, 1)
	cch := make(chan int, 1)
	ach <- 1
	go printA(ach, 'A', bch, &wg)
	go printA(bch, 'B', cch, &wg)
	go printA(cch, 'C', ach, &wg)

	wg.Wait()
}

func printA(ach chan int, c byte, nextCh chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		x := <-ach
		if x > 0 {
			println(string(c))
		}
		nextCh <- 1
	}
	wg.Done()

}
