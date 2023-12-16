package go_channel

import (
	"sync"
	"testing"
)

type Pool struct {
	tasks chan func()
}

func initPool() Pool {
	p := Pool{
		tasks: make(chan func(), 10),
	}

	go func() {
		for {
			select {
			case t, ok := <-p.tasks:
				if !ok {
					go t()
				}
			case p.tasks <- nil:
				close(p.tasks)

			}
		}
	}()

	return p
}

func (p Pool) addTask(f func()) {
	p.tasks <- f
}

func TestPool(t *testing.T) {
	pool := initPool()
	for i := 0; i < 10; i++ {
		pool.addTask(func() {
			println("foo")
		})
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
