/*
Author:ydy
Date:
Desc:
*/
package timex

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTicker(t *testing.T){
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Ticker tick.")
	}
}

func TestTickerDemo(t *testing.T){
	ticker := NewTicker(time.Second * 2)
	defer ticker.Stop()
	for v := range ticker.Chan(){
		fmt.Println("vvvvv---",v)
	}
}
