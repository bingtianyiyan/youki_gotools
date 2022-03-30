/*
Author:ydy
Date:
Desc:
*/
package syncx

import "sync/atomic"

// A OnceGuard is used to make sure a resource can be taken once.
type OnceGuard struct {
	done uint32
}

// CheckTake checks if the resource is taken.
func (m *OnceGuard) CheckTake() bool{
	return atomic.LoadUint32(&m.done) == 1
}

func (m *OnceGuard) Take() bool{
	return atomic.CompareAndSwapUint32(&m.done,0,1)
}
