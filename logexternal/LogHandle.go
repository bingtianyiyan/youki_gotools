package logexternal

import (
	"errors"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logenum"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logiface"
)

/*
 日志消息处理中心
 */

var (
	workCount uint32 = 10
	workQueueLen uint32 = 2500
)

type LogHandle struct {
	LogApis map[logenum.Level]logiface.ILogRouter
	LogTaskWorkerQueue []chan logiface.ILogRequest
}

// NewLogHandle 日志处理对象初始化
func NewLogHandle() *LogHandle{
	return &LogHandle{
		LogApis: map[logenum.Level]logiface.ILogRouter{},
		LogTaskWorkerQueue: make([]chan logiface.ILogRequest,workCount),
	}
}

// AddLogRouter 添加日志路由
func (m *LogHandle) AddLogRouter(logLevel logenum.Level, router logiface.ILogRouter) *LogHandle {
	if _,ok := m.LogApis[logLevel];ok{
		panic(errors.New("LogLevel重复注册"))
		return m
	}
	m.LogApis[logLevel] = router
	return m
}

// InitLogTaskWorkerQueuePool 初始化
func (m *LogHandle) InitLogTaskWorkerQueuePool(){
	for i:=0;i<int(workCount);i++{
		m.LogTaskWorkerQueue[i] = make(chan logiface.ILogRequest,workQueueLen)
		go m.StartLogTaskWorkerQueue(i,m.LogTaskWorkerQueue[i])
	}
}

// SendMsgToLogTaskWorkerQueue 外部调用发送日志内容给Channel
func (m *LogHandle) SendMsgToLogTaskWorkerQueue(request logiface.ILogRequest){
	//得到需要处理此条连接的workerID
	workerID := request.GetLogId() % workCount
	m.LogTaskWorkerQueue[workerID] <- request
}

// StartLogTaskWorkerQueue channel消费日志
func (m *LogHandle) StartLogTaskWorkerQueue(worker int,lcQueue chan logiface.ILogRequest){
	for{
		select {
			case msg := <- lcQueue:
				m.DoLogMsgHandler(worker,msg)
		}
	}
}

// DoLogMsgHandler 实际日志处理
func (m *LogHandle) DoLogMsgHandler(worker int,request logiface.ILogRequest) {
	logFunc,ok := m.LogApis[request.GetLogLevel()]
	if !ok{
		panic(errors.New("请注册Log处理事件" + string(request.GetLogLevel())))
	}
	logFunc.LogExecute(request)
}