/*
Author:ydy
Date:
Desc:
*/
package syncx

import "sync"

type Barrier struct {
	lock sync.Mutex
}

func (m *Barrier) Guard(f func()){
	m.lock.Lock()
	defer m.lock.Unlock()
	f()
}
