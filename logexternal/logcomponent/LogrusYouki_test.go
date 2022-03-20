package logcomponent

import (
	"testing"
)

func inti(){
	InitFileLog()
}

func TestPrint(t *testing.T){
	var log = new(LogrusYouki)
	log.Print("hello print")
}
