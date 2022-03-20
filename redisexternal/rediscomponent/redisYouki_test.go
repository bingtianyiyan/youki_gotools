package rediscomponent

import (
	"fmt"
	"testing"
)

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

func TestRedisGoYouKi_KeyExists(t *testing.T) {
   var resp = redisGoCli.KeyExists("key1")
   if resp.Result{
   	fmt.Println(resp.Result)
   }
}
