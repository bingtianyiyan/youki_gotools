/*
Author:ydy
Date:
Desc: rescue use with defer
*/
package rescue

import "fmt"

func Recover(fs ...func()){
	for _,f := range fs {
		f()
	}

	if p := recover();p != nil{
		fmt.Println("rcover-->",p)
	}
}
