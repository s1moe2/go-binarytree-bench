package main

import "sync"

// Task encapsulates a work item that should go in a work pool
type Task struct {
	f    func() *treeTestResult
}

// NewTask initializes a new task
func NewTask(f func() *treeTestResult) *Task {
	return &Task{
		f:    f,
	}
}

// Run runs a Task and decrements the WaitGroup
func (t *Task) Run(wg *sync.WaitGroup) *treeTestResult {
	res := t.f()
	wg.Done()
	return res
}

// Pool is a worker group that runs a number of tasks at a configured concurrency
type Pool struct {
	Tasks       []*Task
	concurrency int
	tasksChan   chan *Task
	resultsChan chan *treeTestResult
	wg          sync.WaitGroup
}

// NewPool initializes a new pool with a set of tasks and the concurrency value
func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
		resultsChan: make(chan *treeTestResult),
	}
}

// Run runs the execution pool
func (p *Pool) Run() []*treeTestResult {
	// start concurrent workers
	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}

	// send tasks to channel
	p.wg.Add(len(p.Tasks))
	for _, task := range p.Tasks {
		p.tasksChan <- task
	}
	close(p.tasksChan)

	p.wg.Wait()

	// get results
	var results []*treeTestResult
	for i := 0; i < len(p.Tasks); i++ {
		results = append(results, <- p.resultsChan)
	}

	return results
}

// work picks tasks from the tasks channel and executes them
func (p *Pool) work() {
	for task := range p.tasksChan {
		p.resultsChan <- task.Run(&p.wg)
	}
}
