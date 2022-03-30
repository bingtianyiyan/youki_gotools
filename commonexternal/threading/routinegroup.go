/*
Author:ydy
Date:
Desc:
*/
package threading

import "sync"

// A RoutineGroup is used to group goroutines together and all wait all goroutines to be done.
type RoutineGroup struct {
	waitGroup sync.WaitGroup
}

func NewRoutineGroup() *RoutineGroup{
	return new(RoutineGroup)
}


func (m *RoutineGroup) Run(f func()){
	m.waitGroup.Add(1)

	go func() {
		defer m.waitGroup.Done()
		f()
	}()
}


func (m *RoutineGroup) RunSafe(f func()){
	m.waitGroup.Add(1)

	GoSafe(func() {
		defer m.waitGroup.Done()
		f()
	})
}

func(m *RoutineGroup) Wait(){
	m.waitGroup.Wait()
}
