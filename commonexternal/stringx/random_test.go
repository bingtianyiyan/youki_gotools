/*
Author:ydy
Date:
Desc:
*/
package stringx

import (
	"fmt"
	"math/rand"
	"testing"
)

func  TestRandId(t *testing.T) {
	for i:=0;i<10;i++ {
		var rd = Randn(10)
		fmt.Println(rd)
	}
for j :=0;j<10;j++ {
	fmt.Println(rand.Int())
}
}
