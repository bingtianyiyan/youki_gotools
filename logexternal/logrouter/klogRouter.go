package logrouter

import (
	"github.com/bingtianyiyan/youki_gotools/logexternal/logenum"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logiface"
	"k8s.io/klog"
)

func KLogInit(){
	klog.InitFlags(nil)
}

/*
 定义klog日志注册的路由实现方法 ,继承基类
*/

type KLogRouter struct {
	logiface.ILogRouter
}

func (m *KLogRouter) LogExecute(request logiface.ILogRequest){
	 if request.GetLogLevel() == logenum.InfoLevel {
		klog.Info(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.WarnLevel {
		klog.Warning(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.ErrorLevel {
		klog.Error(request.GetLogDataStr())
	}else if request.GetLogLevel() == logenum.FatalLevel {
		klog.Fatal(request.GetLogDataStr())
	}
	klog.Flush()
}
