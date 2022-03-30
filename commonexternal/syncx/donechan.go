/*
Author:ydy
Date:
Desc:
*/
package syncx

import "sync"

// A DoneChan is used as a channel that can be closed multiple times and wait for done.
type DoneChan struct {
	once sync.Once
	done chan struct{}
}

func NewDoneChan() *DoneChan{
	return &DoneChan{
		done: make(chan struct{}),
	}
}

func (m *DoneChan) Close(){
	m.once.Do(func() {
		close(m.done)
	})
}

func (m *DoneChan) Done() chan struct{}{
	return m.done
}
