package logrouter

import (
	"github.com/bingtianyiyan/youki_gotools/logexternal/logcomponent"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logenum"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logiface"
)

var (
	zapObj logiface.ILog = new(logcomponent.ZapYouki)
)

/*
 定义zap日志注册的路由实现方法 ,继承基类
*/

type ZapRouter struct {
	logiface.ILogRouter
}

func (m *ZapRouter) LogExecute(request logiface.ILogRequest){
	if request.GetLogLevel() == logenum.TraceLevel {
		zapObj.Trace(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.DebugLevel {
		zapObj.Debug(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.InfoLevel {
		zapObj.Info(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.WarnLevel {
		zapObj.Warn(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.ErrorLevel {
		zapObj.Error(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.FatalLevel {
		zapObj.Fatal(request.GetLogDataStr())
	} else if request.GetLogLevel() == logenum.PanicLevel {
		zapObj.Panic(request.GetLogDataStr())
	}
}

