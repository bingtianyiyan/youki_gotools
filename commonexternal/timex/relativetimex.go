/*
Author:ydy
Date:
Desc:
*/
package timex

import "time"

// Use the long enough past timex as start timex, in case timex.Now() - lastTime equals 0.
var initTime = time.Now().AddDate(-1, -1, -1)

// Now returns a relative timex duration since initTime, which is not important.
// The caller only needs to care about the relative value.
func Now() time.Duration {
	return time.Since(initTime)
}

// Since returns a diff since given d.
func Since(d time.Duration) time.Duration {
	return time.Since(initTime) - d
}

// Time returns current timex, the same as time.Now().
func Time() time.Time {
	return initTime.Add(Now())
}
