/*
Author:ydy
Date:
Desc:
*/
package testx

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPoolConn(t *testing.T){
	var strPool = sync.Pool{
		New: func() interface {}{
			 var conn,_ = newConn()
			 return conn
		},
	}

	conn := (strPool.Get()).(*sql.DB)
	if conn != nil {
		err := conn.Ping()
		fmt.Println("d0-->",err)
		strPool.Put(conn)
	}

	runtime.GC()
	conn = (strPool.Get()).(*sql.DB)
	if conn != nil {
		err := conn.Ping()
		fmt.Println("Gc-->",err)
		 strPool.Put(conn)
	}

	_conn,_ := newConn()
	strPool.Put(_conn)
	conn = (strPool.Get()).(*sql.DB)
	if conn != nil {
		err := conn.Ping()
		fmt.Println("d1--->",err)
		strPool.Put(conn)
	}
}

func newConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	return db, nil
}
