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

func TestGoSafe(t *testing.T) {
	GoSafe(func() {
		fmt.Println("safe")
		panic("111")
	})
}
