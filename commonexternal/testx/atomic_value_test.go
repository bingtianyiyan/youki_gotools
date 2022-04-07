/*
Author:ydy
Date:
Desc:
*/
package testx

import (
	"fmt"
	"sync/atomic"
	"testing"
)

type Value struct {
	Key string
	Val interface{}
}

type Noaway struct {
	Movies atomic.Value
	Total  atomic.Value
}

func NewNoway() *Noaway{
	n := new(Noaway)
	n.Movies.Store(&Value{Key: "movie",Val: "wolf"})
	n.Total.Store("2.66")
	return n
}

func TestAtomicVal(t *testing.T){
	n := NewNoway()
	val := n.Movies.Load().(*Value)
	total := n.Total.Load().(string)
	fmt.Printf("Movies %v domestic total as of Aug. 27, 2017: %v \n", val.Val, total)
}
