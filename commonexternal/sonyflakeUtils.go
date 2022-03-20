package commonexternal

import (
	"github.com/sony/sonyflake"
)

/*
雪花算法
*/

var (
	SFlake *SnowFlake
)

// SnowFlake SnowFlake算法结构体
type SnowFlake struct {
	sFlake *sonyflake.Sonyflake
}

func init() {
	SFlake = NewSnowFlake()
}

// 模拟获取本机的机器ID
func getMachineID() (mID uint16, err error) {
	mID = 10
	return
}

func NewSnowFlake() *SnowFlake {
	st := sonyflake.Settings{}
	// machineID是个回调函数
	st.MachineID = getMachineID
	return &SnowFlake{
		sFlake: sonyflake.NewSonyflake(st),
	}
}

func (s *SnowFlake) GetID() (uint64, error) {
	return s.sFlake.NextID()
}
