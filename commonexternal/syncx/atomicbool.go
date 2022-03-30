/*
Author:ydy
Date:
Desc:
*/
package syncx

import "sync/atomic"

// An AtomicBool is an atomic implementation for boolean values.
type AtomicBool uint32

// NewAtomicBool returns an AtomicBool.
func NewAtomicBool() *AtomicBool {
	return new(AtomicBool)
}

// ForAtomicBool returns an AtomicBool with given val.
func ForAtomicBool(val bool) *AtomicBool {
	b := NewAtomicBool()
	b.Set(val)
	return b
}

func (m *AtomicBool) Set(val bool){
	if val{
		atomic.StoreUint32((*uint32)(m),1)
	}else {
		atomic.StoreUint32((*uint32)(m),0)
	}
}

// CompareAndSwap compares current value with given old, if equals, set to given val.
func (m *AtomicBool) CompareAndSwap(old, val bool) bool {
	var ol,vl uint32
	if old{
		ol = 1
	}
	if val{
		vl = 1
	}
    return atomic.CompareAndSwapUint32((*uint32)(m),ol,vl)
}

func (m *AtomicBool) True() bool{
	return atomic.LoadUint32((*uint32)(m)) == 1
}