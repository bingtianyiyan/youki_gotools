/*
Author:ydy
Date:
Desc:
*/
package filex

import "os"

//create file
func CreateFile(name string) (*os.File,error){
	return  os.Create(name)
}

//OpenFile
func OpenFile(path string) (*os.File,error){
	return os.Open(path)
}

func OpenFileWithMode(path string,flag int,mode os.FileMode)(*os.File,error){
	return os.OpenFile(path,flag,mode)
}
