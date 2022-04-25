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

func TestWorkerGroup_Start(t *testing.T) {
	var pool = NewWorkerGroup(func() {
		fmt.Println("wk")
	},10)
	 pool.Start()
}
