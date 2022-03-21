package rediscomponent

import (
	"errors"
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

//
//  ConvertTypeOk
//  @Description: 部分特殊情况结果转换
//  @param value
//  @return int64
//  @return bool
//
func convertTypeOk(value interface{}) (int64,bool){
	switch value.(type) {
	case string:
		// 将interface转为string字符串类型
		op, ok := value.(string)
		if ok {
			if op == "OK"{
				return 1,ok
			}
		}
		return 0,ok
	case int32:
		// 将interface转为int32类型
		op, ok := value.(int32)
		return int64(op),ok
	case int64:
		// 将interface转为int64类型
		op, ok := value.(int64)
		return op,ok
	default:
		return 0,false
	}
}

func setSingleKeyExpire(cli redis.Conn,key string,expireTimeSecond int)(interface{},error){
	result,err := cli.Do(redisenum.Key_Expire, key, expireTimeSecond)
	if err != nil {
		cli.Do(redisenum.Key_DelKey, key)
	}
	return result,err
}

func setMoreKeysExpire(cli redis.Conn,expireTimeSecond int,params []interface{})(interface{},error){
	var result interface{}
	var err error
	var line = len(params)
	if line % 2 == 0 {
		for i := 0; i < line; i++ {
			if i%2 == 0 {
				cli.Send(redisenum.Key_Expire,params[i],expireTimeSecond)
			}
		}
		cli.Flush()
		result,err = cli.Receive()
		//设置失败则再次删除
		if err != nil{
			for i := 0; i < line; i++ {
				if i%2 == 0 {
					cli.Send(redisenum.Key_DelKey,params[i])
				}
			}
			cli.Flush()
			result,err = cli.Receive()
		}
	}else {
		err = errors.New("参数不能为奇数")
	}
	return result,err
}

/*
Keys 相关
 */

// KeyExists Key是否存在
func (m RedisGoYouKi) KeyExists(key string) redisenum.RedisResponse {
	response,_ := m.Execute(redisenum.Key_Exists,false,key)
	if response.Result{
	   response.Result,response.Err = redis.Bool(response.Data,response.Err)
	}
	return response
}

// SetExpire 设置过期时间
func (m RedisGoYouKi) SetExpire(key string,expireTime int) redisenum.RedisResponse {
	response,_ := m.Execute(redisenum.Key_Expire,false,key,expireTime)
	return response
}

// DelKeys 删除Key
func (m RedisGoYouKi) DelKeys(keys ...interface{}) redisenum.RedisResponse {
	response,_ := m.Execute(redisenum.Key_DelKey,false,keys...)
	return response
}

// Dump 序列化给定 key ，并返回被序列化的值(如果 key 不存在，那么返回 nil 。 否则，返回序列化之后的值
func (m RedisGoYouKi) Dump(key string) redisenum.RedisResponse {
	response,_ := m.Execute(redisenum.Key_Dump,false,key)
	return response
}

/*
string 相关
 */

// Set 设置单个数据 数据格式 key/data/"EX"固定过期时间字符串/expireTime:过期时间1000,如果expireTime==0则不设置过期时间
func (m RedisGoYouKi) Set(key string,data interface{},expireTime int) redisenum.RedisResponse{
	if expireTime == 0 {
		response,_ := m.Execute(redisenum.Str_Set,false,key,data)
		return response
	}
	response,_ := m.Execute(redisenum.Str_Set,false,key,data,"EX",expireTime)
	return response
}

// Get 获取单个Set的数据
func (m RedisGoYouKi) Get(key string)redisenum.RedisResponse{
	response,_ := m.Execute(redisenum.Str_Get,false,key)
	return response
}

// GetSet 将给定 key 的值设为 value ，并返回 key 的旧值(old value)
func (m RedisGoYouKi) GetSet(key string,data interface{})redisenum.RedisResponse{
	response,_ := m.Execute(redisenum.Str_GetSet,false,key,data)
	return response
}

// MGet 获取所有(一个或多个)给定 key 的值,返回结果可以用redis 封装的 ByteSlices 直接帮你转换
func (m RedisGoYouKi) MGet(keys ...interface{})redisenum.RedisResponse{
	response,_ := m.Execute(redisenum.Str_MGet,false,keys...)
	return response
}

// SetEx 将值 value 关联到 key ，并将 key 的过期时间设为 seconds (以秒为单位)。
func (m RedisGoYouKi) SetEx(key string,data interface{},expireTimeSecond int)redisenum.RedisResponse{
	response,_ := m.Execute(redisenum.Str_SetEx,false,key,expireTimeSecond,data)
	return response
}

// SetNx 只有在 key 不存在时设置 key 的值。(值存在，设置不成功Data=0，不存在设置成功Data=1
func (m RedisGoYouKi) SetNx(key string,data interface{},expireTimeSecond int)redisenum.RedisResponse{
	response,cli := m.Execute(redisenum.Str_SetNx,true,key,data)
	if cli != nil {
		defer cli.Close()
		if  response.Result {
			var result,ok = convertTypeOk(response.Data)
			if ok && result == 1 && expireTimeSecond > 0 {
				response.Data, response.Err = setSingleKeyExpire(cli,key,expireTimeSecond)
			}
		}
	}
	return response
}


// MSet 同时设置一个或多个 key-value 对。
func (m RedisGoYouKi) MSet(expireTimeSecond int,params ...interface{}) redisenum.RedisResponse{
	response,cli := m.Execute(redisenum.Str_MSet,true,params...)
	if cli != nil {
		defer cli.Close()
		if  response.Result {
			response.Data, response.Err = setMoreKeysExpire(cli,expireTimeSecond,params)
		}
	}
	return response
}

// MSetNx 同时设置一个或多个 key-value 对，当且仅当所有给定 key 都不存在。(数据存在设置的时候Data返回0，
func (m RedisGoYouKi) MSetNx(expireTimeSecond int,params ...interface{}) redisenum.RedisResponse{
	response,cli := m.Execute(redisenum.Str_MSetNx,true,params...)
	if cli != nil {
		defer cli.Close()
		if  response.Result {
			var result,ok = convertTypeOk(response.Data)
			if ok && result == 1 {
				response.Data, response.Err = setMoreKeysExpire(cli, expireTimeSecond, params)
			}
		}
	}
	return response
}

// PSetEx 这个命令和 SETEX 命令相似，但它以毫秒为单位设置 key 的生存时间，而不是像 SETEX 命令那样，以秒为单位。
func (m RedisGoYouKi) PSetEx(key string,data interface{},milliseconds int)redisenum.RedisResponse{
	response,_ := m.Execute(redisenum.Str_PSetEx,false,key,milliseconds,data)
	return response
}

// Incr 将 key 中储存的数字值增一
func (m RedisGoYouKi) Incr(key string) redisenum.RedisResponse{
	response,_ := m.Execute(redisenum.Str_Incr,false,key)
	return response
}

// IncrBy 将 key 所储存的值加上给定的增量值（increment） 。
func (m RedisGoYouKi) IncrBy(key string,num int) redisenum.RedisResponse{
	response,_ := m.Execute(redisenum.Str_Incrby,false,key,num)
	return response
}

// IncrByFloat 将 key 所储存的值加上给定的浮点增量值（increment）
func (m RedisGoYouKi) IncrByFloat(key string,num float64) redisenum.RedisResponse{
	response,_ := m.Execute(redisenum.Str_IncrbyFloat,false,key,num)
	return response
}

// Decr 将 key 中储存的数字值减一
func (m RedisGoYouKi) Decr(key string) redisenum.RedisResponse{
	response,_ := m.Execute(redisenum.Str_Decr,false,key)
	return response
}

// DecrBy 将 key 所储存的值加上给定的增量值（increment） 。
func (m RedisGoYouKi) DecrBy(key string,num int) redisenum.RedisResponse{
	response,_ := m.Execute(redisenum.Str_Decrby,false,key,num)
	return response
}

// Append
//如果 key 已经存在并且是一个字符串， APPEND 命令将指定的 value 追加到该 key 原来值（value）的末尾,返回的result是字符长度
func (m RedisGoYouKi) Append(key string,data interface{}) redisenum.RedisResponse{
	response,_ := m.Execute(redisenum.Str_Append,false,key,data)
	return response
}

