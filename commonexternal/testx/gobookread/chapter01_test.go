package gobookread

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
	"testing"
	"time"
)

func TestChapter(t *testing.T){
	inputerader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入")
	input,err := inputerader.ReadString('\n')
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	input = input[:len(input)-1]
	fmt.Printf("Hello %s!\n",input)
}

type myTest1 struct {
  v int
}

func (self myTest1) Print(){
	self.v = self.v +1
	fmt.Println("Normal" , self.v)
}

func (self *myTest1) PrintPoint(){
	self.v = self.v +1
	fmt.Println("Point",self.v)
}

func TestFunc01(t *testing.T)  {
	runtime.GOMAXPROCS(4)
	debug.SetMaxThreads(1000)
//	//var v myTest1
//	//v.v = 1
//	//v.PrintPoint()
//	//v.Print()
//
//
//	//var v = new(myTest1)
//	//v.v = 1
//	//v.PrintPoint()
//	//v.Print()
//	//fmt.Println(v.v)
//
//	defer func() {
//		if e := recover();e != nil{
//			fmt.Println(e)
//		}
//	}()
//
//	fmt.Println(os.Getpid())
//
//	signalRec := make(chan os.Signal,1)
//	sigs := []os.Signal{syscall.SIGINT,syscall.SIGKILL,syscall.SIGQUIT,syscall.SIGTERM}
//	signal.Notify(signalRec,sigs...)
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		for sig := range signalRec{
//			fmt.Println(sig)
//		}
//		fmt.Println("end signal")
//		wg.Done()
//	}()
//	time.Sleep(time.Second *2)
//	signal.Stop(signalRec)
//	close(signalRec)
//wg.Wait()
//	cmd0 := exec.Command("cmd","-n","this ssss")
//	studo0,err := cmd0.StderrPipe()
//	if err != nil{
//		fmt.Println("cmd err",err)
//		return
//	}
//	if err := cmd0.Start();err != nil{
//		fmt.Println("cmd err",err)
//		return
//	}
//
//	//output := make([]byte,1024)
//	outbuff0 := bufio.NewReader(studo0)
//	out0,_,err := outbuff0.ReadLine()
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//	fmt.Printf("%s\n",string(out0))
//	//n,err := studo0.Read(output)
//	//if err != nil{
//	//	if err != io.EOF {
//	//		fmt.Println("cmd err", err)
//	//		return
//	//	}
//	//}
//	//fmt.Println(output[:n])
//
//	var ch chan int
//	ch = make(chan int,1)
//	ch <- 1
//	close(ch)
//	for x := range ch {
//		fmt.Println(x)
//	}

	var wg1 sync.WaitGroup
	wg1.Add(2)
	go server(wg1)
	go client(wg1)
    wg1.Wait()

}

func server(wg1 sync.WaitGroup){
	//net 服务端
	listen ,err := net.Listen("tcp","127.0.0.1:8089")
	if err != nil{
		return
	}
	fmt.Println("server conn")
	defer listen.Close()
	//wg1.Add(1)
	for{
		conn,err := listen.Accept()
		if err != nil{
			return
		}
		fmt.Println("server accept client")
		//处理信息
		go func(sConn net.Conn) {
			defer wg1.Done()

			//read
			rd := bufio.NewReader(sConn)
			var buf [126]byte
			n,err := rd.Read(buf[:])
			if err != nil {
				if err == io.EOF {
					//return
				}else {
					fmt.Println("server read err", err)
					return
				}
			}
			fmt.Println("server receive msg-->",string(buf[:n]))

			//write
			sConn.Write([]byte("server send msg\n"))
		}(conn)
	}

}

func client(wg1 sync.WaitGroup){
	//net 客户端
	cConn,err := net.Dial("tcp","127.0.0.1:8089")
	if err != nil{
		return
	}
	defer cConn.Close()
	fmt.Println("client dial conn")
	//wg1.Add(1)
	for{
		defer wg1.Done()
		//写入数据
		go func() {
			for i:=0;i<5;i++ {
				cConn.Write([]byte("client send msg\n"))
				time.Sleep(time.Second*3)
			}
		}()


		var buf [1024]byte
		//read
		rd :=bufio.NewReader(cConn)
		n,err := rd.Read(buf[:])
		if err != nil{
			if err == io.EOF{
				//return
			}else {
				fmt.Println("client read err", err)
				return
			}
		}
		recvStr := string(buf[:n])
		fmt.Println("client receive data->",recvStr)

	}
}

var strChan = make(chan string,3)
func TestChannel(t *testing.T){
	syncchan1 := make(chan struct{},1)
	syncchan2 := make(chan struct{},2)
	//receive
	go func() {
		<- syncchan1
		fmt.Println("receive...")
		time.Sleep(time.Second*2)
		for{
            if item,ok := <-strChan;ok{
				fmt.Println("receive item--",item)
			}else {
				break
			}
		}
		fmt.Println("stop receive")
		syncchan2 <- struct{}{}
	}()

	//send
	go func() {
		for _,item := range []string{"a","b","c","d"}{
			strChan <- item
			fmt.Println("send item-->",item)
			if item == "c"{
				syncchan1 <- struct{}{}
			}
		}
		time.Sleep(time.Second * 2)
		close(strChan)
		syncchan2 <- struct{}{}
	}()
	<- syncchan2
	<- syncchan2
}

var chan1 chan int
var chan2 chan int
var channels =[]chan int{chan1,chan2}
var numbers = []int{1,2,3,4,5}
func TestSelect(t *testing.T){
	Loop:
	for {
		select {
		case getChan(0) <- getNumber(0):
			fmt.Println("chan0")
			break
		case getChan(1) <- getNumber(1):
			fmt.Println("chan1")
			break
		default:
			fmt.Println("default")
			break Loop
		}
	}
}

func getNumber(i int)int{
	fmt.Println("getnumber:",i)
return numbers[i]
}

func getChan(i int) chan int{
	fmt.Println("getchan:",i)
	return channels[i]
}

func TestTimer(t *testing.T){
	//var pTimer = time.NewTimer(time.Second*2)
	//fmt.Println(time.Now())
	//fmt.Println(pTimer.Stop())
	//exptime1 := <-pTimer.C
	//fmt.Println(exptime1)
	//fmt.Println(pTimer.Stop())

	ti := time.AfterFunc(time.Second, func() {
		fmt.Println("time after func")
	})
	defer ti.Stop()
	time.Sleep(time.Second*10)

	var reChan chan int = make(chan int)
	tk := time.Tick(time.Second * 2)//time.NewTicker(time.Second *2)
	//defer tk.Stop()
	go func() {
		for _ = range tk{
			select {
			case reChan <- 1:
			case reChan <- 2:
			case reChan <- 3:
		   }
		}
	}()
	for item := range reChan{
		fmt.Println(item)
	}
}

var locker sync.Mutex
var cond = sync.NewCond(&locker)
func TestLockCond(t *testing.T){
	locker.Lock()
	for i := 0; i < 10; i++ {
		go func(x int) {
			cond.L.Lock()         // 获取锁
			defer cond.L.Unlock() // 释放锁
			cond.Wait()           // 等待通知，阻塞当前 goroutine
			// 通知到来的时候, cond.Wait()就会结束阻塞, do something. 这里仅打印
			fmt.Println(x)
		}(i)
	}
	time.Sleep(time.Second * 1) // 睡眠 1 秒，等待所有 goroutine 进入 Wait 阻塞状态
	fmt.Println("Signal...")
	cond.Signal()               // 1 秒后下发一个通知给已经获取锁的 goroutine
	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal()               // 1 秒后下发下一个通知给已经获取锁的 goroutine
	time.Sleep(time.Second * 1)
	cond.Broadcast()            // 1 秒后下发广播给所有等待的goroutine
	fmt.Println("Broadcast...")
	time.Sleep(time.Second * 1) // 睡眠 1 秒，等待所有 goroutine 执行完毕
	time.Sleep(time.Second*100)
}

func TestYYP(t *testing.T){
	
}