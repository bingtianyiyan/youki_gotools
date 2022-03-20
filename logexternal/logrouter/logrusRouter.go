package logrouter

import (
	"github.com/bingtianyiyan/youki_gotools/logexternal/logcomponent"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logenum"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logiface"
)

var (
	logrusObj *logcomponent.LogrusYouki
)

func LogrusInit(){
	//logrus组件全局对象初始化
	logcomponent.InitFileLog()
	logrusObj = new(logcomponent.LogrusYouki)
}

/*
 定义logrus日志注册的路由实现方法 ,继承基类
*/

type LogrusRouter struct {
	logiface.ILogRouter
}

func (m *LogrusRouter) LogExecute(request logiface.ILogRequest){
	if request.GetLogLevel() == logenum.TraceLevel {
		logrusObj.Trace(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.DebugLevel {
		logrusObj.Debug(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.InfoLevel {
		logrusObj.Info(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.WarnLevel {
		logrusObj.Warn(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.ErrorLevel {
		logrusObj.Error(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.FatalLevel {
		logrusObj.Fatal(request.GetLogDataStr())
	} else if request.GetLogLevel() == logenum.PanicLevel {
		logrusObj.Panic(request.GetLogDataStr())
	}
}

