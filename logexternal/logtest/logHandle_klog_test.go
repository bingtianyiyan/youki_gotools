package logtest

import (
	"flag"
	"github.com/bingtianyiyan/youki_gotools/logexternal"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logcomponent"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logenum"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logrouter"
	"testing"
	"time"
)

var (
	logObj2 *logexternal.LogHandle
)

func initTestLocal(){
	//初始化日志队列
	logObj2 = logexternal.NewLogHandle()
	//注册日志路由（根据需要写
	logObj2.
		AddLogRouter(logenum.InfoLevel,&logrouter.KLogRouter{}).
		AddLogRouter(logenum.WarnLevel,&logrouter.KLogRouter{}).
		AddLogRouter(logenum.ErrorLevel,&logrouter.KLogRouter{}).
		AddLogRouter(logenum.FatalLevel,&logrouter.KLogRouter{})
	logObj2.InitLogTaskWorkerQueuePool()

	//初始化KLog
	logcomponent.InitKLog()
}

//测试klog
func TestKlogLogInfo(t *testing.T){
	initTestLocal()
	flag.Set("logtostderr", "false") //日志输出到stderr，不输出到日志文件。false为关闭
	flag.Set("log_file", "myKlogfile.log")
	flag.Parse()

	msgData := time.Now().Format("2006-01-02 15:04:05")
	var request = logexternal.SetLogRequestData("Hello World klog error->" + msgData, logenum.ErrorLevel)
	logObj2.SendMsgToLogTaskWorkerQueue(request)

	var requestErr = logexternal.SetLogRequestData("Hello World klog warn->" + msgData, logenum.WarnLevel)
	logObj2.SendMsgToLogTaskWorkerQueue(requestErr)
	time.Sleep(time.Second * 20)
}
