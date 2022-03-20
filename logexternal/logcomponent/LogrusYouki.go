package logcomponent

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sync"
)

/*
logrus日志组件二次封装
 */

// LogrusYouki
type LogrusYouki struct {

}

var (
	instance *logrus.Logger
	once sync.Once
)

// NewLogrusInstance 创建logrus实例
func NewLogrusInstance() *logrus.Logger{
	if instance == nil {
		once.Do(func() {
			instance = logrus.New()
		})
	}
	return instance
}

// InitFileLog 按文件保存初始化配置
func InitFileLog() {
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	NewLogrusInstance().SetFormatter(&logrus.JSONFormatter{})
	//写屏幕设置
	NewLogrusInstance().SetOutput(os.Stdout)
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	file, err := openOrCreate("1.log")
	//NewLogrusInstance().SetOutput(file)
	//同时写文件和屏幕
	writers := []io.Writer{
		file,
		os.Stdout}
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		NewLogrusInstance().SetOutput(fileAndStdoutWriter)
	} else {
		NewLogrusInstance().Info("failed to log to file.")
	}
	//设置最低loglevel
	NewLogrusInstance().SetLevel(logrus.TraceLevel)
}

func openOrCreate(name string) (*os.File, error) {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return f, err
}


func (m *LogrusYouki) Print(args ...interface{}){
	NewLogrusInstance().Print(args...)
}

func (m *LogrusYouki) Trace(args ...interface{}){
	NewLogrusInstance().Trace(args...)
}

func (m *LogrusYouki) Debug(args ...interface{}){
	NewLogrusInstance().Debug(args...)
}

func (m *LogrusYouki) Info(args ...interface{}){
	NewLogrusInstance().Info(args...)
}
func (m *LogrusYouki) Warn(args ...interface{}){
	NewLogrusInstance().Warn(args...)
}
func (m *LogrusYouki) Error(args ...interface{}){
	NewLogrusInstance().Error(args...)
}
func (m *LogrusYouki) Fatal(args ...interface{}){
	NewLogrusInstance().Fatal(args...)
}
func (m *LogrusYouki) Panic(args ...interface{}){
	NewLogrusInstance().Panic(args...)
}
