package logcomponent

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

/*
Zap的日志组件
*/

type ZapYouki struct {

}

var (
	zapInstance *zap.Logger
)

const (
	logTmFmt = "2006-01-02 15:04:05"
)

func InitZapLog(){
	GetLogger()
}

func GetLogger() {
	Encoder := GetEncoder()
	WriteSyncer := GetWriteSyncer()
	LevelEnabler := GetLevelEnabler()
	ConsoleEncoder := GetConsoleEncoder()
	newCore := zapcore.NewTee(
		zapcore.NewCore(Encoder, WriteSyncer, LevelEnabler),                          // 写入文件
		zapcore.NewCore(ConsoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel), // 写入控制台
	)
	zapInstance = zap.New(newCore, zap.AddCaller())
	zap.ReplaceGlobals(zapInstance)
}

// GetEncoder 自定义的Encoder
func GetEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller_line",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     "  ",
			EncodeLevel:    cEncodeLevel,
			EncodeTime:     cEncodeTime,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   cEncodeCaller,
		})
}

// GetConsoleEncoder 输出日志到控制台
func GetConsoleEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
}

// GetWriteSyncer 自定义的WriteSyncer
func GetWriteSyncer() zapcore.WriteSyncer {
	//lumberJackLogger := &lumberjack.Logger{
	//	Filename:   "./zap.log",
	//	MaxSize:    200,
	//	MaxBackups: 10,
	//	MaxAge:     30,
	//}
	//return zapcore.AddSync(lumberJackLogger)
	file, _ := os.Create("./zap.log")
	return zapcore.AddSync(file)
}

// GetLevelEnabler 自定义的LevelEnabler
func GetLevelEnabler() zapcore.Level {
	return zapcore.InfoLevel
}

// cEncodeLevel 自定义日志级别显示
func cEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// cEncodeTime 自定义时间格式显示
func cEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format(logTmFmt) + "]")
}

// cEncodeCaller 自定义行号显示
func cEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}


func (m *ZapYouki) Print(args ...interface{}){
	//for k, v := range args {
	//	str := fmt.Sprintf("%v", v)
	//	fmt.Println(k, str)
	//}
	zapInstance.Debug(args[0].(string))
}

func (m *ZapYouki) Trace(args ...interface{}){
	zapInstance.Debug(args[0].(string))
}

func (m *ZapYouki) Debug(args ...interface{}){
	zapInstance.Debug(args[0].(string))
}

func (m *ZapYouki) Info(args ...interface{}){
	zapInstance.Info(args[0].(string))
}
func (m *ZapYouki) Warn(args ...interface{}){
	zapInstance.Warn(args[0].(string))
}
func (m *ZapYouki) Error(args ...interface{}){
	zapInstance.Error(args[0].(string))
}
func (m *ZapYouki) Fatal(args ...interface{}){
	zapInstance.Fatal(args[0].(string))
}
func (m *ZapYouki) Panic(args ...interface{}){
	zapInstance.Panic(args[0].(string))
}

