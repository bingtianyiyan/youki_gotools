/*
Author:ydy
Date:
Desc:
*/
package threading

import "github.com/bingtianyiyan/youki_gotools/commonexternal/rescue"

// GoSafe runs the given fn using another goroutine, recovers if fn panics.
func GoSafe(fn func()){
	go RunSafe(fn)
}

func RunSafe(fn func()){
	defer rescue.Recover()
	fn()
}
