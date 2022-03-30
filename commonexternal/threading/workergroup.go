/*
Author:ydy
Date:
Desc:
*/
package threading

// A WorkerGroup is used to run given number of workers to process jobs.
type WorkerGroup struct {
	job     func()
	workers int
}

// NewWorkerGroup returns a WorkerGroup with given job and workers.
func NewWorkerGroup(job func(),workers int) WorkerGroup{
	return WorkerGroup{
		job:job,
		workers: workers,
	}
}

// Start starts a WorkerGroup.
func (m WorkerGroup) Start(){
    group := NewRoutineGroup()
	for i := 0;i <m.workers;i++{
		group.RunSafe(m.job)
	}
	group.Wait()
}
