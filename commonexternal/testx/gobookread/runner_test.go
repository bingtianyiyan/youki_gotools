package gobookread

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestRunner(t *testing.T)  {
	var tk = NewRunner(time.Second*20)
	tk.AddTask(CreateTask(),CreateTask(),CreateTask())

	err := tk.Start()
	if err != nil{
		switch err {
		case errorInterrupt:
			fmt.Println("pr--interput")
			os.Exit(1)
		case errorTimeout:
			fmt.Println("pr--timeout")
			os.Exit(2)
		}
	}

  fmt.Println("process end")
}

func CreateTask() func(int){
	return func(i int) {
		fmt.Println("task--",i)
	}
}

type Runner struct {
	interrupt chan os.Signal

	complete chan error

	timeout <-chan time.Time

	tasks []func(int)
}

var errorTimeout = errors.New("receive timeout")
var errorInterrupt = errors.New("received interrupt")

func NewRunner(d time.Duration)*Runner{
	return &Runner{
		interrupt: make(chan os.Signal,1),
		complete: make(chan error),
		timeout:time.After(d),
	}
}

func (self *Runner) AddTask(tasks ...func(int)){
	self.tasks = append(self.tasks,tasks...)
}

func (self *Runner) Start() error{
	signal.Notify(self.interrupt,os.Interrupt)

	go func() {
		self.complete <- self.run()
	}()

	select {
	case err := <- self.complete:
		return  err
	case <- self.timeout :
		return errorTimeout
	}
}

func (self *Runner) run() error {
    for id,task := range self.tasks{
    	if self.gointerrupt(){
    		return errorInterrupt
		}
		task(id)
	}
	return nil
}

func (self *Runner) gointerrupt() bool {
	select {
	case <- self.interrupt:
		signal.Stop(self.interrupt)
		return true
	default:
		return  false
 }
}
