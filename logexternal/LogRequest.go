package logexternal

import "github.com/bingtianyiyan/youki_gotools/logexternal/logenum"

type LogRequest struct {
	logId    uint32
	logData  []byte
	logLevel logenum.Level
}


func SetLogRequestData(data string,level logenum.Level) *LogRequest{
	return &LogRequest{
		logData: []byte(data),
		logLevel:level,
	}
}

func (m *LogRequest) GetLogId() uint32{
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

