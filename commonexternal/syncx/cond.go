/*
Author:ydy
Date:
Desc:
*/
package syncx

import (
	"github.com/bingtianyiyan/youki_gotools/commonexternal"
	"time"
)

// A Cond is used to wait for conditions.
type Cond struct {
	signal chan struct{}
}

func NewCond() *Cond{
	return &Cond{
		signal: make(chan struct{}),
	}
}

// WaitWithTimeout wait for signal return remain wait time or timed out.
func (m *Cond) WaitWithTimeout(timeout time.Duration) (time.Duration, bool) {
   timer := time.NewTimer(timeout)
   defer timer.Stop()

   begin := commonexternal.Now()

	select {
        case <- m.signal:
        	elapsed := commonexternal.Since(begin)
        	remainTimeout := timeout - elapsed
        	return  remainTimeout,true
	case <- timer.C:
		return  0,false
   }
}

func (m *Cond) Wait(){
	<- m.signal
}

// Signal wakes one goroutine waiting on c, if there is any.
func (m *Cond) Signal(){
	select {
	    case m.signal <- struct{}{} :
	    default:

	}
}

