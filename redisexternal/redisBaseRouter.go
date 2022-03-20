package redisexternal

import "github.com/bingtianyiyan/youki_gotools/redisexternal/redisenum"

/*
 redis基类方法 实现iredis
 */
type RedisBaseRouter struct {

}

func (m *RedisBaseRouter) KeyExists(key string) redisenum.RedisResponse{
	return redisenum.RedisResponse{}
}
