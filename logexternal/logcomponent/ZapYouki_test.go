package logcomponent

import "testing"

func TestZapLog(t *testing.T){
	InitZapLog()
	var log = new(ZapYouki)
	log.Error("zapLogerror")
}
