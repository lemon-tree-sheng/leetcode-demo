package slice

/*
*
两种方式

1. 传统方式
创建 n 个goroutine，阻塞等待任务
2. 用 channel 来控制同时运行的 goroutine 数量即可
*/
type MyRoutinePool struct {
	cha chan int
}

func Init(num int) *MyRoutinePool {
	r := MyRoutinePool{
		cha: make(chan int, num),
	}
	return &r
}

func (r *MyRoutinePool) addTask(f func()) {
	r.cha <- 1
	go f()
	<-r.cha
}
