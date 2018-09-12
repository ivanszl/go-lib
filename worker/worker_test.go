package worker

import (
	"fmt"
	"testing"
)

type OneWorker struct {
	In chan interface{}
}

func (w *OneWorker) Do() <-chan interface{} {
	return w.In
}

func (w *OneWorker) Handle(param interface{}) {
	fmt.Printf("%v\n", param)
}
func TestWorker(t *testing.T) {
	test := &OneWorker{
		In: make(chan interface{}, 100),
	}
	worker := Worker{}
	worker.Add(test, 10)
	go func() {
		for i := 0; i < 1000; i++ {
			test.In <- i
		}
		close(test.In)
	}()
	worker.Wait()
	fmt.Printf("all finish")
}
