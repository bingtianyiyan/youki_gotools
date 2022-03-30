/*
Author:ydy
Date:
Desc:
*/
package syncx

import "testing"

func TestDoneChan_Close(t *testing.T) {
	var o = NewDoneChan()
	for i := 0;i <10;i++{
		o.Close()
	}
}
