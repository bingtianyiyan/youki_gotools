/*
Author:ydy
Date:
Desc:
*/
package testx

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"
)

var ErrTimeOut = errors.New("执行者执行超时")
var ErrInterrupt = errors.New("执行者被中断")

//一个执行者，可以执行任何任务，但是这些任务是限制完成的，
//该执行者可以通过发送终止信号终止它
type Runner struct {
	tasks     []func(int)      //要执行的任务
	complete  chan error       //用于通知任务全部完成
	timeout   <-chan time.Time //这些任务在多久内完成
	interrupt chan os.Signal   //可以控制强制终止的信号

}

func NewRunner(d time.Duration) *Runner{
	return &Runner{
		complete: make(chan error),
		timeout: time.After(d),
		interrupt: make(chan os.Signal,1),
	}
}

func (self *Runner) Add(task ...func(int)){
	self.tasks = append(self.tasks,task...)
}
//检查是否接收到了中断信号
func (self *Runner) isInterrupt() bool {
	select {
	case <-self.interrupt:
		signal.Stop(self.interrupt)
		return true
	default:
		return false
	}
}

func (self *Runner) run() error{
     for k,v := range self.tasks{
     	if self.isInterrupt(){
     		return nil
		}
		v(k)
	 }
	 return nil
}

func (self *Runner) Start() error{
	signal.Notify(self.interrupt,os.Interrupt)

	go func() {
		self.complete <- self.run()
	}()

	select {
	    case err := <- self.complete:
	    	fmt.Println("comple..")
	    	return err
	case <- self.timeout:
		return ErrTimeOut

	}
}

func TestWorker(t *testing.T) {
	log.Println("...开始执行任务...")

	timeout := 3 * time.Second
	r := NewRunner(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err:=r.Start();err!=nil{
		switch err {
		case ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}
	log.Println("...任务执行结束...")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("正在执行任务%d", id)
		time.Sleep(time.Duration(id)* time.Second)
	}
}
