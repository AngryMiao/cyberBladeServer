package async

import (
	"context"
	"sync"
)

type Pool struct {
	workers    int
	tasks      chan Task
	results    chan Result
	wg         sync.WaitGroup
	ctx        context.Context
	cancelFunc context.CancelFunc
}

type Task struct {
	Handler func() (interface{}, error)
}

type Result struct {
	Data  interface{}
	Error error
}

func NewPool(workers int) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	return &Pool{
		workers:    workers,
		tasks:      make(chan Task, workers*2),
		results:    make(chan Result, workers*2),
		ctx:        ctx,
		cancelFunc: cancel,
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go p.worker()
	}
}

func (p *Pool) worker() {
	defer p.wg.Done()

	for {
		select {
		case task, ok := <-p.tasks:
			if !ok {
				return
			}

			data, err := task.Handler()

			p.results <- Result{
				Data:  data,
				Error: err,
			}
		case <-p.ctx.Done():
			return
		}
	}
}

func (p *Pool) Submit(handler func() (interface{}, error)) Result {
	task := Task{
		Handler: handler,
	}

	p.tasks <- task
	return <-p.results
}

func (p *Pool) Stop() {
	p.cancelFunc()
	close(p.tasks)
	p.wg.Wait()
	close(p.results)
}
