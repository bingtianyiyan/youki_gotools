/*
Author:ydy
Date:
Desc:
*/
package syncx

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeoutLimit_Borrow(t *testing.T) {
	var rl = NewTimeoutLimit(2)
	fmt.Println("r1-->",rl.Borrow(time.Second))
	fmt.Println("r2-->",rl.Borrow(time.Second))
	rl.Return()
	fmt.Println("r3-->",rl.Borrow(time.Second))
	rl.Return()
	fmt.Println("r4-->",rl.Borrow(time.Second))
	fmt.Println("r5-->",rl.Borrow(time.Second))
}
