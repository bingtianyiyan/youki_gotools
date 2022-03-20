package logiface

import (
	"github.com/bingtianyiyan/youki_gotools/logexternal/logenum"
)

type ILogRequest interface {
	GetLogId() uint64           //日志随机号
	GetLogData() []byte         //获取日志请求消息的参数数据
	GetLogLevel() logenum.Level //获取日志请求的级别
	GetLogDataStr() string      //获取日志请求消息的参数数据转化成string
}
