/*
Author:ydy
Date:
Desc:
*/
package threading

import (
	"fmt"
	"testing"
)

func TestRoutineGroup_Run(t *testing.T) {
	var rg = NewRoutineGroup()
	var l chan struct{} = make(chan struct{},10)
	for i :=0;i<10;i++{
	   rg.Run(func() {
	   	fmt.Println("a")
			l <- struct{}{}
		})
	}
	<- l
	rg.Wait()
}
