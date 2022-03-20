package redisiface

import "github.com/bingtianyiyan/youki_gotools/redisexternal/redisenum"

/*
redis通用方法接口
 */

type IRedis interface {
   // KeyExists Key是否存在
   KeyExists(key string) redisenum.RedisResponse
}
