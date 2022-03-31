/*
Author:ydy
Date:
Desc:
*/
package timex

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTimerDemo(t *testing.T){
	timer := time.NewTimer(2 * time.Second)

	select {
	case <- timer.C:
		log.Println("Delayed 5s, start to do something.")
	case <- time.After(time.Second * 2):
		fmt.Println("time after")
	}

	timer.Stop()
	timer.Reset(time.Second)
}
