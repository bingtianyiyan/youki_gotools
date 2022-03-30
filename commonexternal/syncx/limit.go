/*
Author:ydy
Date:
Desc:
*/
package syncx

import "errors"

type Limit struct {
	pool chan struct{}
}

func NewLimit(captity int) Limit{
	return Limit{
		pool: make(chan struct{},captity),
	}
}

func (m Limit) Borrow(){
	m.pool <- struct{}{}
}

func (m Limit) TryBorrrow() bool{
	select {
	case m.pool <- struct{}{}:
		return true
	default:
		return false

	}
}

func (m Limit) Return() error{
	select {
	case <- m.pool:
     return nil
	default:
		return errors.New("pool is full")
	}
}
