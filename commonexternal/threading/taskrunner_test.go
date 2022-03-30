/*
Author:ydy
Date:
Desc:
*/
package threading

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewTaskRunner(t *testing.T) {
	var wg sync.WaitGroup
	var taskPool = NewTaskRunner(10)

	for i:=0;i<10;i++{
		wg.Add(1)
		j := i
		taskPool.TaskSchedule(func() {
			defer wg.Done()
			fmt.Println(j)
		})
	}


	wg.Wait()
}
