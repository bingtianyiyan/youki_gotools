/*
Author:ydy
Date:
Desc:
*/
package testx

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	var result = Add(1,2)
	var excepted = 3
	if result != excepted{
		fmt.Println("err not actual result")
		return
	}
	fmt.Println("success excepted result")
	assert.Equal(t,result,excepted)
}

func TestAdd2(t *testing.T) {
	var result = Add(3,2)
	var excepted = 3
	assert.Equal(t,result,excepted)
}

/*
子测试
 */

func TestSubAdd(t *testing.T) {
	// setup
	t.Logf("Setup")

	t.Run("group", func(t *testing.T) {
		t.Run("Test1", TestAdd)
		t.Run("Test2", TestAdd2)
	})
}

/*
benchmark
 */
func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		 MakeSliceWithoutAlloc()
	}
}

func BenchmarkMakeSliceWithPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeSliceWithPreAlloc()
	}
}
