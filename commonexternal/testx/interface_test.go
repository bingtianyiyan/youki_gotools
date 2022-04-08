/*
Author:ydy
Date:
Desc:
*/
package testx

import (
	"fmt"
	"testing"
	"unsafe"
)

type iface struct {
	itab, data uintptr
}

func TestInterfaceDynamic(t *testing.T) {
	var a interface{} = nil
	var b interface{} = (*int)(nil)

	x := 5
	var c interface{} = (*int)(&x)

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))

	fmt.Println(ia, ib, ic)
	fmt.Println(*(*int)(unsafe.Pointer(ic.data)))
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("[Name: %s], [Age: %d]", s.Name, s.Age)
}

func TestInterfacePrintToString(t *testing.T) {
	var s = Student{
		Name: "qcrao",
		Age:  18,
	}

	fmt.Println(s)
}
