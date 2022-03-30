/*
Author:ydy
Date:
Desc:
*/
package syncx

import (
	"fmt"
	"testing"
)

func TestOnceGuard_Take(t *testing.T) {
    var ob = new(OnceGuard)
    var rs = ob.Take()
    fmt.Println(rs)
    fmt.Println(ob.CheckTake())
}
