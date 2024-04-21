package slice

type MyWaitGroup struct {
	num int
	ch  chan int
}

func (w MyWaitGroup) Add(delta int) {
	w.num += delta
}

func (w MyWaitGroup) Done() {
	w.ch <- 1
}

func (w MyWaitGroup) Wait() {
	for {
		x := <-w.ch
		if x > 0 {
			w.num--
			if w.num == 0 {
				break
			}
		}
	}
}
