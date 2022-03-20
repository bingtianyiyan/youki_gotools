package redisenum
/*
redis返回结果
 */

type RedisResponse struct {
	//成功/失败
	Result bool
	//返回错误
	Err error
	//具体结果数据
	Data interface{}
}