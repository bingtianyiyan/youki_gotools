package logrouter

import (
	"github.com/bingtianyiyan/youki_gotools/logexternal/logcomponent"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logenum"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logiface"
)

var (
	klogObj logiface.ILog = new(logcomponent.KLogYouki)
)

/*
 定义klog日志注册的路由实现方法 ,继承基类
*/

type KLogRouter struct {
	logiface.ILogRouter
}

func (m *KLogRouter) LogExecute(request logiface.ILogRequest){
	 if request.GetLogLevel() == logenum.InfoLevel {
		 klogObj.Info(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.WarnLevel {
		 klogObj.Warn(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.ErrorLevel {
		 klogObj.Error(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.FatalLevel {
		 klogObj.Fatal(request.GetLogDataStr())
	}
	logcomponent.KLogFlush()
}
