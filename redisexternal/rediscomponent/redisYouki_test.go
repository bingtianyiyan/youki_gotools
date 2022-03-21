package rediscomponent

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
)

type Book struct {
	BookName  string
	Author    string
	PageCount string
	Press     string
}
var top1 = Book{
	BookName:  "Crazy golang",
	Author:    "Moon",
	PageCount: "600",
	Press:     "GoodBook",
}

var (
	redisGoCli = new(RedisGoYouKi)
)

func init(){
    redisGoCli.redisCli = GetRedisGoPoolInstance(RedisGoConfig{
       redisAddress: "127.0.0.1:6379",
       maxIdle: 10,
       maxActive: 100,
       maxIdleTimeout:300,
       maxTimeout: 3000,
       lazyLimit: false,
       maxSize: 100,
	})
}

/*
Key相关
 */

func TestRedisGoYouKi_KeyExists(t *testing.T) {
   var resp = redisGoCli.KeyExists("key")
   if resp.Result{
   	fmt.Println(resp.Result)
   }
}

func TestRedisGoYouKi_SetExpire(t *testing.T) {
	var resp = redisGoCli.SetExpire("key",1000000)
	if resp.Result{
		fmt.Println(resp.Result)
	}
}

func TestRedisGoYouKi_DelKeys(t *testing.T) {
	var resp = redisGoCli.DelKeys("key")
	if resp.Result{
		fmt.Println(resp.Result)
	}
}

func TestRedisGoYouKi_Dump(t *testing.T) {
	var resp = redisGoCli.Dump("key")
	if resp.Result{
		str,err := redis.String(resp.Data,resp.Err)
		fmt.Println(err,str)
	}
}

/*
String相关
 */

func TestRedisGoYouKi_Set(t *testing.T) {
	var resp = redisGoCli.Set("key2","234",100000)//设置过期时间
	var resp2 = redisGoCli.Set("key3","234",0)//不设置过期时间
	fmt.Println(resp,resp2)
}

func TestRedisGoYouKi_Get(t *testing.T) {
	var resp = redisGoCli.Get("key2")
	fmt.Println(redis.String(resp.Data,resp.Err))
}

func TestRedisGoYouKi_GetSet(t *testing.T) {
	var resp = redisGoCli.GetSet("key2","567")
	fmt.Println(redis.String(resp.Data,resp.Err))
}

func TestRedisGoYouKi_MGet(t *testing.T) {
	var resp = redisGoCli.MGet("key2","key3")
	var byteData,err = redis.ByteSlices(resp.Data,resp.Err)
	if err != nil{
		fmt.Println(err)
		return
	}
	for _,v := range byteData{
		fmt.Println(string(v))
	}
}

func TestRedisGoYouKi_SetEx(t *testing.T) {
	var resp = redisGoCli.SetEx("key2","999",1000000)
	fmt.Println(resp)
}

func TestRedisGoYouKi_SetNx(t *testing.T) {
	var resp = redisGoCli.SetNx("key2","10000",1000000)
	fmt.Println(redis.String(resp.Data,resp.Err))
	var resp2 = redisGoCli.SetNx("key4","10001",1000000)//设置不存在数据
	fmt.Println(redis.String(resp2.Data,resp2.Err))
}

func TestRedisGoYouKi_MSet(t *testing.T) {
	var resp = redisGoCli.MSet(100000,"key05","444","key06","333","key07","你哈M")
	fmt.Println(redis.String(resp.Data,resp.Err))
}

func TestRedisGoYouKi_MSetNx(t *testing.T) {
	var resp = redisGoCli.MSetNx(100000,"key05","444","key06","333")
	fmt.Println(redis.String(resp.Data,resp.Err))
	var resp2 = redisGoCli.MSetNx(100000,"key08","444","key09","333")//设置不存在的
	fmt.Println(redis.String(resp2.Data,resp2.Err))
}

func TestRedisGoYouKi_PSetEx(t *testing.T) {
	var resp = redisGoCli.PSetEx("key10","tttt",1000000)
	fmt.Println(redis.String(resp.Data,resp.Err))
}

func TestRedisGoYouKi_Incr(t *testing.T) {
	var resp = redisGoCli.Incr("key4")
	fmt.Println(redis.String(resp.Data,resp.Err))
}

func TestRedisGoYouKi_Append(t *testing.T) {
	var resp = redisGoCli.Append("key09","yyyyy")
	fmt.Println(redis.Int64(resp.Data,resp.Err))
}

/*
Hash
 */

func TestRedisGoYouKi_HSet(t *testing.T) {
	var resp = redisGoCli.HSetNx("HKey01","name","youki",100000)
	fmt.Println(redis.Int64(resp.Data,resp.Err))
}

func TestRedisGoYouKi_HMSetMany(t *testing.T) {
	var resp = redisGoCli.HMSetMany("HKey01",&top1,0)
	fmt.Println(redis.Int64(resp.Data,resp.Err))
}

func TestRedisGoYouKi_HSetNx(t *testing.T) {
	var resp = redisGoCli.HSetNx("HKey01","name","youki1",0)
	fmt.Println(redis.Int64(resp.Data,resp.Err))
}

func TestRedisGoYouKi_HExists(t *testing.T) {
	var resp = redisGoCli.HExists("HKey01","name")
	fmt.Println(redis.Int64(resp.Data,resp.Err))
}

func TestRedisGoYouKi_HGet(t *testing.T) {
	var resp = redisGoCli.HGet("HKey01","name")
	fmt.Println(redis.String(resp.Data,resp.Err))
}

func TestRedisGoYouKi_HGetAll(t *testing.T) {
	var resp = redisGoCli.HGetAll("HKey01")
	fmt.Println(redis.Strings(resp.Data,resp.Err))
}

func TestRedisGoYouKi_HMGet(t *testing.T) {
	var resp = redisGoCli.HMGet("HKey01","name","BookName")
	fmt.Println(redis.Strings(resp.Data,resp.Err))
}

func TestRedisGoYouKi_HVals(t *testing.T) {
	var resp = redisGoCli.HVals("HKey01")
	fmt.Println(redis.Strings(resp.Data,resp.Err))
}

func TestRedisGoYouKi_HIncrBy(t *testing.T) {
	var resp = redisGoCli.HIncrBy("HKey01","PageCount",20)
	fmt.Println(redis.Int64(resp.Data,resp.Err))
}