/*
Author:ydy
Date:
Desc:
*/
package threading

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"testing"
	"time"
)

func TestAntsDemo(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10000)
	p, _ := ants.NewPool(10)
	defer p.Release()
	for i := 0; i < 10000; i++ {
		if i%2 == 0 || i%7 == 0 {
			p.Submit(func() {
				fmt.Println("task 2-->", i)
				time.Sleep(time.Millisecond * 10)
				wg.Done()
			})
		} else {
			p.Submit(func() {
				fmt.Println("task 1-->", i)
				time.Sleep(time.Millisecond * 10)
				wg.Done()
			})
		}
	}
	wg.Wait()
	fmt.Println("finsid-->")
}

func TestAntsNewPoolWithFuncDemo(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(100)
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		fmt.Println(i)
		wg.Done()
	})
	defer p.Release()
	for i := 0; i < 100; i++ {
		p.Invoke(i)
		time.Sleep(time.Millisecond * 10)
	}
	wg.Wait()
	fmt.Println("finsid-->")
}
