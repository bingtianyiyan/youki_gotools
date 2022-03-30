/*
Author:ydy
Date:
Desc:
*/
package syncx

import "sync"

// Once returns a func that guarantees fn can only called once.
func Once(f func()) func(){
	once := new(sync.Once)
	return func() {
		once.Do(f)
	}
}
