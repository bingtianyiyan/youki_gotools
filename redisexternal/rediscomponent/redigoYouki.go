package rediscomponent

import (
	"fmt"
	"github.com/bingtianyiyan/youki_gotools/redisexternal/redisenum"
	"github.com/garyburd/redigo/redis"
	"time"
)

/*
 redis组件：github.com/garyburd/redigo/redis 二次封装
 */

type RedisGoYouKi struct {
	redisCli       *redis.Pool
}

// RedisGoConfig 配置文件Model
type RedisGoConfig struct {
	redisAddress string
	pwd string
	maxIdle        int
	maxActive      int
	maxIdleTimeout time.Duration
	maxTimeout     time.Duration
	lazyLimit      bool
	maxSize        int
}

// GetRedisGoPoolInstance 线程池连接
func GetRedisGoPoolInstance(config RedisGoConfig) *redis.Pool{
	return &redis.Pool{
		MaxIdle:     config.maxIdle   , /* 最大的空闲连接数 */
		MaxActive:   config.maxActive,   /* 最大的激活连接数 */
		IdleTimeout: config.maxIdleTimeout,
		Dial: func() (redis.Conn, error) {
			var c redis.Conn
			var err error
			if config.pwd == ""{
				c, err = redis.Dial("tcp", config.redisAddress)
			}else{
				c, err = redis.Dial("tcp", config.redisAddress,redis.DialPassword(config.pwd))
			}
			if err != nil {
				fmt.Println("redis conn err",err)
				return nil, err
			}
			fmt.Println("redis conn success")
			return c, nil
		},
	}
}

// Execute
//  @Description: 执行通用方法(方法名/是否需要设置过期时间/执行参数)
//  @receiver m
//  @param methodName
//  @param setExpireTime
//  @param params
//  @return redisenum.RedisResponse
//  @return redis.Conn
//
func (m *RedisGoYouKi) Execute(methodName string,setExpireTime bool,params ...interface{}) (redisenum.RedisResponse,redis.Conn){
	c := m.redisCli.Get()
	if !setExpireTime {
		defer c.Close()
	}
	res,err := c.Do(methodName,params...)
	if !setExpireTime {
		return redisenum.RedisResponse{
			Result: err == nil,
			Err:    err,
			Data:   res,
		}, nil
	}
	return redisenum.RedisResponse{
		Result: err == nil,
		Err:    err,
		Data:   res,
	}, c
}


//Keys 相关

// KeyExists Key是否存在
func (m RedisGoYouKi) KeyExists(key string) redisenum.RedisResponse {
	response,_ := m.Execute(redisenum.Key_Exists,false,key)
	if response.Result{
	   response.Result,response.Err = redis.Bool(response.Data,response.Err)
	}
	return response
}
