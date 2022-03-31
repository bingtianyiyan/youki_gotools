/*
Author:ydy
Date:
Desc:
*/
package testx

////文件生成命令：mockgen -source=mock.go -destination=mock_mock.go -package=testx

// db.go
type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}