/*
Author:ydy
Date:
Desc:
*/
package syncx

import (
	"errors"
	"time"
)

// ErrTimeout is an error that indicates the borrow timeout.
var ErrTimeout = errors.New("borrow timeout")

// A TimeoutLimit is used to borrow with timeouts.
type TimeoutLimit struct {
	limit Limit
	cond  *Cond
}

// NewTimeoutLimit returns a TimeoutLimit.
func NewTimeoutLimit(n int) TimeoutLimit {
	return TimeoutLimit{
		limit: NewLimit(n),
		cond:  NewCond(),
	}
}

func (m *TimeoutLimit) Borrow(timeout time.Duration) error{
    var result = m.TryBorrow()
    if result {
		return nil
	}
	var ok bool
    for{
	   timeout,ok = m.cond.WaitWithTimeout(timeout)
	   if ok && m.TryBorrow(){
	   	return nil
	   }
       if timeout <0 {
       	return ErrTimeout
	   }
	}
}

func (m *TimeoutLimit) TryBorrow() bool{
  return m.limit.TryBorrrow()
}

func (m *TimeoutLimit) Return() error{
    err := m.limit.Return()
    if err != nil{
    	return  err
	}

	m.cond.Signal()
    return nil
}
