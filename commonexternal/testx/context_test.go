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

var (
	key = "name"
)

func TestContextValues(t *testing.T){
   ctx,cancel := context.WithCancel(context.Background())
   ctxVal :=context.WithValue(ctx,key,"ctxVal--->")
   go watch(ctxVal)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context){
	for {
		select {
		    case <- ctx.Done():
		    	fmt.Println(ctx.Value(key),"exit")
				return
		default:
			fmt.Println(ctx.Value(key),"default")
			time.Sleep(2 * time.Second)
		}
	}
}

func TestContextwittimeout(t *testing.T){
	ctx,cancel := context.WithTimeout(context.Background(),time.Second * 10)
	defer cancel()

	for{
		select {
		     case <- ctx.Done():
		     	fmt.Println("exit")
				 return
		default:
			   fmt.Println("default")
			time.Sleep(time.Second*2)
		}
	}
	fmt.Println("finish...")
}

func TestContextwitDeadline(t *testing.T){
	ctx,cancel := context.WithDeadline(context.Background(),time.Now().Add(time.Second*10))
	defer cancel()

	for{
		select {
		case <- ctx.Done():
			fmt.Println("exit,with no later then time")
			return
		default:
			fmt.Println("default")
			time.Sleep(time.Second*2)
		}
	}
	fmt.Println("finish...")
}

func TestContextwitCancel(t *testing.T){
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()
	go childCtx(ctx)
     i := 1
	for{
		select {
		case <- ctx.Done():
			fmt.Println("exit,cancel",i)
			return
		default:
			fmt.Println("default")
			i++
			if i == 10{
				cancel()
			}
			time.Sleep(time.Second*2)
		}
	}
	fmt.Println("finish...")
}

func childCtx(ctx context.Context){
	for {
		select {
		case <-ctx.Done():
			fmt.Println("parent is cancel")
			return
		default:
			fmt.Println("child print..")
			time.Sleep(time.Second)
		}
	}
}
