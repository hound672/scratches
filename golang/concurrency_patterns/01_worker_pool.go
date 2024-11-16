package main

type WorkerPool struct {
	taskCh chan func()
}

func NewWorkerPool(workerNumber int) *WorkerPool {
	return &WorkerPool{}
}

func (wp *WorkerPool) AddTask(task func()) error {
	return nil
}

func (wp *WorkerPool) Shutdown() {}

func main() {

}
