package go_channel_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Task struct {
	ID int
}

type Worker struct {
	ID       int
	TaskChan chan Task
	Done     chan bool
}

type Pool struct {
	Workers    []*Worker
	TaskQueue  chan Task
	ResultChan chan string
	wg         sync.WaitGroup // 用于等待所有协程完成任务后关闭结果通道和工作协程通道。
}

func NewPool(numWorkers, queueSize int) *Pool {
	pool := &Pool{
		TaskQueue:  make(chan Task, queueSize),
		ResultChan: make(chan string),
	}
	for i := 0; i < numWorkers; i++ {
		w := &Worker{
			ID:       i,
			TaskChan: make(chan Task),
			Done:     make(chan bool),
		}

		pool.Workers = append(pool.Workers, w)

		go func() { // 启动工作协程处理任务队列中的任务。
			for task := range w.TaskChan {
				result := fmt.Sprintf("worker %d processed task %d", w.ID, task.ID)
				pool.ResultChan <- result

				// 模拟耗时操作，这里可以替换为实际需要执行的业务逻辑。
				time.Sleep(time.Second)

				w.Done <- true // 告知主线程该任务已完成。
			}
		}()
	}

	return pool
}

func (p *Pool) AddTask(task Task) {
	p.TaskQueue <- task // 将新任务添加到任务队列中。
}

func (p *Pool) Start() {
	go func() {
		for task := range p.TaskQueue {
			worker := p.getAvailableWorker()
			if worker == nil { // 如果没有可用的工作协程，则等待一个任务完成后再重新分配。
				<-p.Workers[0].Done
				p.Workers[0].TaskChan <- task
			} else {
				worker.TaskChan <- task
			}
		}

		// 关闭结果通道和工作协程通道。
		close(p.ResultChan)
		for _, w := range p.Workers {
			close(w.TaskChan)
			<-w.Done
		}

		p.wg.Wait() // 等待所有协程完成任务。
	}()
}

func (p *Pool) getAvailableWorker() *Worker {
	for _, w := range p.Workers {
		select {
		case <-w.Done: // 检查是否有空闲的工作协程。
			return w
		default:
		}
	}

	return nil
}

func TestFoo(t *testing.T) {
	pool := NewPool(3, 10)

	go func() { // 启动一个 goroutine 监听并打印处理结果。
		for result := range pool.ResultChan {
			fmt.Println(result)
		}
	}()

	for i := 1; i <= 20; i++ { // 添加一些任务到任务队列中，总共添加了20个任务。
		pool.AddTask(Task{ID: i})
	}

	pool.Start()
}
