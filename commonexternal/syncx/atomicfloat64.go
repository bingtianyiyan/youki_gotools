/*
Author:ydy
Date:
Desc:
*/
package syncx

import (
	"math"
	"sync/atomic"
)

// An AtomicFloat64 is an implementation of atomic float64.
type AtomicFloat64 uint64

// NewAtomicFloat64 returns an AtomicFloat64.
func NewAtomicFloat64() *AtomicFloat64 {
	return new(AtomicFloat64)
}

// ForAtomicFloat64 returns an AtomicFloat64 with given val.
func ForAtomicFloat64(val float64) *AtomicFloat64 {
	f := NewAtomicFloat64()
	f.Set(val)
	return f
}

func (m *AtomicFloat64) Set(val float64){
	atomic.StoreUint64((*uint64)(m),(uint64)(val))
}

// Add adds val to current value.
func (m *AtomicFloat64) Add(val float64) float64{
	for {
		var old = m.Load()
		var nv = val + old
		if m.CompareAndSwap(old,nv){
			return nv
		}
	}

}

// CompareAndSwap compares current value with old, if equals, set the value to val.
func (m *AtomicFloat64) CompareAndSwap(old, val float64) bool {
	return atomic.CompareAndSwapUint64((*uint64)(m),math.Float64bits(old),math.Float64bits(val))
}

func (m *AtomicFloat64) Load() float64{
	return math.Float64frombits(atomic.LoadUint64((*uint64)(m)))
}
