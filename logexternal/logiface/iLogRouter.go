package logiface
/*
日志接口
 */
type ILogRouter interface {
	LogExecute(ILogRequest)
}
