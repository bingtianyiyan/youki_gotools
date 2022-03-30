/*
Author:ydy
Date:
Desc:
*/
package syncx

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCond_WaitWithTimeout(t *testing.T) {
	var cond = NewCond()
    var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10)
		fmt.Println("before")
		cond.WaitWithTimeout(time.Second*10)
		fmt.Println("after")
	}()
	wg.Wait()
}

func TestSignalNoWait(t *testing.T) {
	cond := NewCond()
	cond.Signal()
}


func TestTimeoutCondWaitTimeoutRemain(t *testing.T) {
	var cond = NewCond()
	var myReChan chan time.Duration = make(chan time.Duration,1)
	defer close(myReChan)
   var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer 	wg.Done()
		remain,_ := cond.WaitWithTimeout(time.Second * 200)
		myReChan <- remain
	}()

	go func(){
		defer wg.Done()
        cond.Signal()
	}()

	wg.Wait()

   remainTime := <- myReChan
   fmt.Println("remaintime-->",remainTime)
}
