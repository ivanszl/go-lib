package worker

import (
	"sync"
)

type IRunable interface {
	Do() <-chan interface{}
	Handle(param interface{})
}

type Worker struct {
	wg sync.WaitGroup
}

func (w *Worker) Add(runable IRunable, workerNum int) {
	w.wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go func() {
			defer w.wg.Done()
			for param := range runable.Do() {
				runable.Handle(param)
			}
		}()
	}
}

func (w *Worker) Wait() {
	w.wg.Wait()
}
