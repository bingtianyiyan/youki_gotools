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
	logObj *logexternal.LogHandle
)

func initlogrus(){
	logcomponent.InitFileLog()
	//初始化日志队列
	logObj = logexternal.NewLogHandle()
	//注册日志路由（根据需要写
	logObj.AddLogRouter(logenum.TraceLevel,&logrouter.LogrusRouter{}).
		AddLogRouter(logenum.DebugLevel,&logrouter.LogrusRouter{}).
		AddLogRouter(logenum.InfoLevel,&logrouter.LogrusRouter{}).
		AddLogRouter(logenum.WarnLevel,&logrouter.LogrusRouter{}).
		AddLogRouter(logenum.ErrorLevel,&logrouter.LogrusRouter{}).
		AddLogRouter(logenum.FatalLevel,&logrouter.LogrusRouter{}).
		AddLogRouter(logenum.PanicLevel,&logrouter.LogrusRouter{})
	   logObj.InitLogTaskWorkerQueuePool()
}

//测试logrus
func TestLogrusLogInfo(t *testing.T){
	initlogrus()
	msgData := time.Now().Format("2006-01-02 15:04:05")
	var request = logexternal.SetLogRequestData("Hello World logurs error->" + msgData, logenum.ErrorLevel)
	logObj.SendMsgToLogTaskWorkerQueue(request)

	var requestErr = logexternal.SetLogRequestData("Hello World logurs warn->" + msgData, logenum.WarnLevel)
	logObj.SendMsgToLogTaskWorkerQueue(requestErr)
	time.Sleep(time.Second * 20)
}

