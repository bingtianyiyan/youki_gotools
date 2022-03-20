package logcomponent

import "k8s.io/klog"

/*
K8s的日志组件
 */

type KLogYouki struct {

}

func InitKLog(){
	klog.InitFlags(nil)
}

func KLogFlush(){
	klog.Flush()
}


func (m *KLogYouki) Print(args ...interface{}){
	klog.Info(args...)
}

func (m *KLogYouki) Trace(args ...interface{}){
	klog.Info(args...)
}

func (m *KLogYouki) Debug(args ...interface{}){
	klog.Info(args...)
}

func (m *KLogYouki) Info(args ...interface{}){
	klog.Info(args...)
}
func (m *KLogYouki) Warn(args ...interface{}){
	klog.Warning(args...)
}
func (m *KLogYouki) Error(args ...interface{}){
	klog.Error(args...)
}
func (m *KLogYouki) Fatal(args ...interface{}){
	klog.Fatal(args...)
}
func (m *KLogYouki) Panic(args ...interface{}){
	klog.Fatal(args...)
}

