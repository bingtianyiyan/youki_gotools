package gogoroutine_open_test

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
	"testing"
	"time"

	"gopkg.in/go-playground/pool.v3"
)

/*
控制goroutine数量
 */
func TestGoroutineUseChannelSync(t *testing.T){
    poolTest()
}

var (
	// channel长度
	poolCount      = 5
	// 复用的goroutine数量
	goroutineCount = 10
)

func poolTest() {
	jobsChan := make(chan int, poolCount)

	// workers
	var wg sync.WaitGroup
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range jobsChan {
				// ...
				fmt.Println(item)
			}
		}()
	}

	// senders
	for i := 0; i < 1000; i++ {
		jobsChan <- i
	}

	// 关闭channel，上游的goroutine在读完channel的内容，就会通过wg的done退出
	close(jobsChan)
	wg.Wait()
}

func TestGoroutineUseSemaphore(t *testing.T){
	semphoreDemo()
}

const (
	// 同时运行的goroutine上限
	Limit = 3
	// 信号量的权重
	Weight = 1
)

func semphoreDemo() {
	names := []string{
		"小白",
		"小红",
		"小明",
		"小李",
		"小花",
	}

	sem := semaphore.NewWeighted(Limit)
	var w sync.WaitGroup
	for _, name := range names {
		w.Add(1)
		go func(name string) {
			sem.Acquire(context.Background(), Weight)
			// ... 具体的业务逻辑
			fmt.Println(name, "-吃饭了")
			time.Sleep(2 * time.Second)
			sem.Release(Weight)
			w.Done()
		}(name)
	}
	w.Wait()

	fmt.Println("ending--------")
}

func TestGoPlayground(t *testing.T){
	p := pool.NewLimited(10)
	defer p.Close()

	user := p.Queue(getUser(13))
	other := p.Queue(getOtherInfo(13))

	user.Wait()
	if err := user.Error(); err != nil {
		// handle error
	}

	// do stuff with user
	username := user.Value().(string)
	fmt.Println(username)

	other.Wait()
	if err := other.Error(); err != nil {
		// handle error
	}

	// do stuff with other
	otherInfo := other.Value().(string)
	fmt.Println(otherInfo)
}

func getUser(id int) pool.WorkFunc {

	return func(wu pool.WorkUnit) (interface{}, error) {

		// simulate waiting for something, like TCP connection to be established
		// or connection from pool grabbed
		time.Sleep(time.Second * 1)

		if wu.IsCancelled() {
			// return values not used
			return nil, nil
		}

		// ready for processing...

		return "Joeybloggs", nil
	}
}

func getOtherInfo(id int) pool.WorkFunc {

	return func(wu pool.WorkUnit) (interface{}, error) {

		// simulate waiting for something, like TCP connection to be established
		// or connection from pool grabbed
		time.Sleep(time.Second * 1)

		if wu.IsCancelled() {
			// return values not used
			return nil, nil
		}

		// ready for processing...

		return "Other Info", nil
	}
}

func TestMySemphore(t *testing.T){
	//var wg sync.WaitGroup
	var num int = 0
	//wg.Add(20)
	for i:=0;i<20;i++ {
		go func() {
			_ = NewMyMutex()
			num++
			//wg.Done()
		}()
	}
	//wg.Wait()
	fmt.Println(num)
}

type Mutex Mysemphore

func NewMyMutex() Mutex {
	return Mutex(NewMySemphore(1))
}

//使用channel实现同步
type Mysemphore chan struct{}

func NewMySemphore(size int) Mysemphore{
	return make(chan struct{},size)
}

func (self Mysemphore) Mylock(){
	self <- struct {}{}
}

func (self Mysemphore) Myunlock(){
	<- self
}