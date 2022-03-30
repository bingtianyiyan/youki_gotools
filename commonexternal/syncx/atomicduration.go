/*
Author:ydy
Date:
Desc:
*/
package syncx

import (
	"sync/atomic"
	"time"
)

// An AtomicDuration is an implementation of atomic duration.
type AtomicDuration int64

// NewAtomicDuration returns an AtomicDuration.
func NewAtomicDuration() *AtomicDuration {
	return new(AtomicDuration)
}

func (m *AtomicDuration) ForAtomicDuration(time time.Duration) *AtomicDuration{
     var o = NewAtomicDuration()
     o.Set(time)
     return m
}

func (m *AtomicDuration) Set(time time.Duration){
	atomic.StoreInt64((*int64)(m),int64(time))
}


// CompareAndSwap compares current value with old, if equals, set the value to val.
func (m *AtomicDuration) CompareAndSwap(old, val time.Duration) bool {
	 return atomic.CompareAndSwapInt64((*int64)(m),int64(old),int64(val))
}


func (m *AtomicDuration)  Load() time.Duration{
	return time.Duration( atomic.LoadInt64((*int64)(m)))
}
