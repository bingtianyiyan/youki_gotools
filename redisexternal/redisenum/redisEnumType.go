package redisenum


const (
	//字符串相关
	Str_Set = "Set"
	Str_Get = "Get"
	Str_GetSet = "GETSET"
	Str_MGet = "MGET"
	Str_SetEx = "SETEX"
	Str_SetNx ="SETNX"
	Str_MSet = "MSet"
	Str_MSetNx ="MSETNX"
	Str_PSetEx ="PSETEX"
	Str_Incr = "INCR"
	Str_Incrby = "INCRBY"
	Str_IncrbyFloat = "INCRBYFLOAT"
	Str_Decr ="DECR"
	Str_Decrby ="DECRBY"
	Str_Append = "APPEND"
	//Hash 哈希
	hdel = "HDEL"
	hexists = "HEXISTS"
	hget = "HGET"
	hgetall = "HGETALL"
	hincrby = "HINCRBY"
	hincrbyFloat = "HINCRBYFLOAT"
	hKeys = "HKEYS"
	hLen = "HLEN"
	hmget = "HMGET"
	hmset = "HMSET"
	hset = "HSet"
	hsetnx = "HSETNX"
	hvals = "HVALS"
	//key相关
	Key_DelKey = "Del"
	Key_Dump = "DUMP"
	Key_Expire ="Expire"
	Key_Exists ="EXISTS"
	Key_Pttl = "PTTL"
)