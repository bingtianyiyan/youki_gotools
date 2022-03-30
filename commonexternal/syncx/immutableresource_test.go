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

func TestImmutableResource_Get(t *testing.T) {
	var rs = NewImmutableResource(func()(interface{},error){
		fmt.Println("a")
		return 1,nil
	}, func(resource *ImmutableResource) {

	})

	var data,err = rs.Get()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(data)
}
