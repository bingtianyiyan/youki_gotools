/*
Author:ydy
Date:
Desc:
*/
package syncx

import (
	"github.com/bingtianyiyan/youki_gotools/commonexternal/timex"
	"sync"
	"time"
)

const defaultRefreshInterval = time.Second

type (
	// ImmutableResourceOption defines the method to customize an ImmutableResource.
	ImmutableResourceOption func(resource *ImmutableResource)

	// An ImmutableResource is used to manage an immutable resource.
	ImmutableResource struct {
		fetch           func() (interface{}, error)
		resource        interface{}
		err             error
		lock            sync.RWMutex
		refreshInterval time.Duration
		lastTime        *AtomicDuration
	}
)

func NewImmutableResource(fn func() (interface{}, error), opts ...ImmutableResourceOption) *ImmutableResource{
	var resource  = new(ImmutableResource)
	resource.fetch = fn
	resource.refreshInterval = defaultRefreshInterval
    resource.lastTime = NewAtomicDuration()
	for _,v := range opts{
        v(resource)
	}
	return resource
}


// Get gets the immutable resource, fetches automatically if not loaded.
func (m *ImmutableResource) Get() (interface{}, error) {
     m.lock.RLock()
     rs := m.resource
     m.lock.RUnlock()
     if rs != nil{
     	return rs,nil
	 }

    m.maybeRefresh(func() {
		var res,err = m.fetch()
		m.lock.Lock()
		if err != nil{
			m.err = err
		}else{
			m.resource,m.err = res,err
		}
		m.lock.Unlock()
	})

    m.lock.RLock()
     rs = m.resource
    m.lock.RUnlock()
     return rs,m.err
}

func (m *ImmutableResource) maybeRefresh(execute func()) {
	now := timex.Now()
	lastTime := m.lastTime.Load()
	if lastTime == 0 || lastTime+m.refreshInterval < now {
		m.lastTime.Set(now)
		execute()
	}
}

// WithRefreshIntervalOnFailure sets refresh interval on failure.
// Set interval to 0 to enforce refresh every timex if not succeeded, default is time.Second.
func WithRefreshIntervalOnFailure(interval time.Duration) ImmutableResourceOption {
	return func(resource *ImmutableResource) {
		resource.refreshInterval = interval
	}
}