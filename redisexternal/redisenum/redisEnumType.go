package redisenum


const (
	//字符串相关
	set = "Set"
	get = "Get"
	getSet = "GETSET"
	mGet = "MGET"
	setEx = "SETEX"
	setNx ="SETNX"
	mSet = "MSet"
	mSetNx ="MSETNX"
	pSetEx ="PSETEX"
	incr = "INCR"
	incrby = "INCRBY"
	incrbyFloat = "INCRBYFLOAT"
	decr ="DECR"
	decrby ="DECRBY"
	append = "APPEND"
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
	delKey = "Del"
	dump = "DUMP"
	expire ="Expire"
	Key_Exists ="EXISTS"
	pttl = "PTTL"
)