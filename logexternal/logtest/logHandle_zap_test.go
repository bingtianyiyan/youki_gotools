package logtest

import (
	"github.com/bingtianyiyan/youki_gotools/logexternal"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logcomponent"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logenum"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logrouter"
	"testing"
	"time"
)

var (
	logObj3 *logexternal.LogHandle
)

func initZap(){
	logcomponent.InitZapLog()
	//初始化日志队列
	logObj3 = logexternal.NewLogHandle()
	//注册日志路由（根据需要写
	logObj3.AddLogRouter(logenum.TraceLevel,&logrouter.ZapRouter{}).
		AddLogRouter(logenum.DebugLevel,&logrouter.ZapRouter{}).
		AddLogRouter(logenum.InfoLevel,&logrouter.ZapRouter{}).
		AddLogRouter(logenum.WarnLevel,&logrouter.ZapRouter{}).
		AddLogRouter(logenum.ErrorLevel,&logrouter.ZapRouter{}).
		AddLogRouter(logenum.FatalLevel,&logrouter.ZapRouter{}).
		AddLogRouter(logenum.PanicLevel,&logrouter.ZapRouter{})
	logObj3.InitLogTaskWorkerQueuePool()
}

//测试zap
func TestZapLogInfo(t *testing.T){
	initZap()
	msgData := time.Now().Format("2006-01-02 15:04:05")
	var request = logexternal.SetLogRequestData("Hello World zap error->" + msgData, logenum.ErrorLevel)
	logObj3.SendMsgToLogTaskWorkerQueue(request)

	var requestErr = logexternal.SetLogRequestData("Hello World zap warn->" + msgData, logenum.InfoLevel)
	logObj3.SendMsgToLogTaskWorkerQueue(requestErr)
	time.Sleep(time.Second * 20)
}
