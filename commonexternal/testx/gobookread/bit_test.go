package gobookread

import (
	"bytes"
	"flag"
	"fmt"
	"testing"
	"time"
	"unicode/utf8"
)

type Flags uint
const (
	FlagUp Flags = 1 << iota // is up
	FlagBroadcast // supports broadcast access capability
	FlagLoopback // is a loopback interface
	FlagPointToPoint // belongs to a point-to-point link
	FlagMulticast // supports multicast access capability
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func TestBitYunsuan(t *testing.T){
	//^   0001  0010     0011
	fmt.Println(1^2)
	//&^  0001 0010   0001
    fmt.Println(1&^2)
	//   0001 0011   0000 先全部置为0，然后以前面数据为准，如果两者对比不一样，则以前面数据为准，否则为0
	fmt.Println(1&^3)
	// 0010 0101   0010
	fmt.Println(2&^5)

	n := utf8.RuneCountInString("hr你好啊!")
	fmt.Println(n)
	fmt.Println("==========================")
	s := "hr你好啊!"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}


	fmt.Println(intsToString([]int{2,3,4}))

	fmt.Println(FlagLoopback)

	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
}

func intsToString(values []int) string{
	var buf bytes.Buffer
	buf.WriteString("[")
	for i,v := range values{
		if i >0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf,"%d",v)
	}
	buf.WriteString("]")
	return buf.String()
}


func revertArr(values []int){
	for i,j :=0,len(values) - 1;i<j;i,j=i+1,j-1{
		values[i],values[j] = values[j],values[i]
	}
}