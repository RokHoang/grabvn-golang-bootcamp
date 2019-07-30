package main

import "fmt"

type Pool struct {
	pool    chan chan Job
	job     chan Job
	workers []Worker
}

func NewWorkerPool(count int) *Pool {
	p := Pool{
		pool: make(chan chan Job, count),
		job:  make(chan Job),
	}
	for i := 0; i < count; i++ {
		p.workers = append(p.workers, Worker{
			index: i + 1,
			pool:  p.pool,
			job:   make(chan Job),
		})
	}
	for _, worker := range p.workers {
		go worker.start()
	}
	go p.start()
	return &p
}

func (p *Pool) dispatch(job Job) {
	p.job <- job
}

func (p *Pool) start() {
	for {
		job := <-p.job
		worker := <-p.pool
		//step 2
		worker <- job
	}
}

type DoAction interface {
	Do()
}

type Job DoAction

type Worker struct {
	index int
	job   chan Job
	pool  chan chan Job
}

func (w Worker) start() {
	for {
		//push job to pool
		//step 1
		w.pool <- w.job
		//step 3
		value := <-w.job
		// value.Print(w.index)
		fmt.Println("This is worker:", w.index)
		value.Do()
	}
}
