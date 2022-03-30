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

func TestWorkerGroup_Start(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	var pool = NewWorkerGroup(func() {
		fmt.Println("wk")
		wg.Done()
	},10)
	go pool.Start()
	wg.Wait()
}
