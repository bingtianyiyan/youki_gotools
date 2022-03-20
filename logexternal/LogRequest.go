package logexternal

import (
	"github.com/bingtianyiyan/youki_gotools/commonexternal"
	"github.com/bingtianyiyan/youki_gotools/logexternal/logenum"
)

type LogRequest struct {
	logId    uint64
	logData  []byte
	logLevel logenum.Level
}


func SetLogRequestData(data string,level logenum.Level) *LogRequest{
	logId,_ := commonexternal.SFlake.GetID()
	return &LogRequest{
		logId: logId,
		logData: []byte(data),
		logLevel:level,
	}
}

func (m *LogRequest) GetLogId() uint64{
	return m.logId
}

func (m *LogRequest) GetLogData() []byte {
	return m.logData
}

func (m *LogRequest) GetLogLevel() logenum.Level {
	return m.logLevel
}

func (m *LogRequest) GetLogDataStr() string{
	return string(m.logData)
}

