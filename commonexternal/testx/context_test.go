/*
Author:ydy
Date:
Desc:
*/
package testx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContextDemo(t *testing.T){
    ctx,cancel := context.WithCancel(context.Background())
    defer cancel()
    for i := range gen(ctx){
		fmt.Println(i)
		if i == 10{
			cancel()
			break
		}
	}
	fmt.Println("finish.....")
}

func gen(ctx context.Context) <-chan int{
	ch := make(chan int)
	go func() {
		var n int
		for{
			select {
			    case <- ctx.Done():
					return
					case ch <- n:
						n++
						time.Sleep(time.Second)
			}
		}
	}()
	return ch
}
