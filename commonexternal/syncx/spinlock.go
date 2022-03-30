/*
Author:ydy
Date:
Desc:
*/
package syncx

import (
	"runtime"
	"sync/atomic"
)

// A SpinLock is used as a lock a fast execution.
type SpinLock struct {
	lock uint32
}

func (m *SpinLock) Lock() {
	for m.TryLock(){
		runtime.Gosched()//不会阻塞别的grouthine
	}
}


func (m *SpinLock) TryLock() bool{
	return atomic.CompareAndSwapUint32(&m.lock,0,1)
}

func (m *SpinLock) UnLock(){
	atomic.StoreUint32(&m.lock,0)
}
