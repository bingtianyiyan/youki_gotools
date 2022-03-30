/*
Author:ydy
Date:
Desc:
*/
package threading

import "github.com/bingtianyiyan/youki_gotools/commonexternal/rescue"

// A TaskRunner is used to control the concurrency of goroutines.
type TaskRunner struct {
    limitChan chan struct{}
}

func NewTaskRunner(concurrency int) *TaskRunner{
    return &TaskRunner{
    	limitChan: make(chan struct{},concurrency),
	}
}

func (m *TaskRunner) TaskSchedule(task func()){
	m.limitChan <- struct{}{}

    go func() {
    	defer rescue.Recover(func() {
			<- m.limitChan
		})

		task()
	}()

}
