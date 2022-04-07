/*
Author:ydy
Date:
Desc:
*/
package db

import "database/sql"

type Dao interface { // 接口申明
	Version() (string, error)
}

type dao struct { // 默认实现
	db *sql.DB
}

func (d dao) Version() (string, error) {
	var version string
	row := d.db.QueryRow("SELECT VERSION()")
	if err := row.Scan(&version); err != nil {
		return "", err
	}
	return version, nil
}

func NewDao(db *sql.DB) *dao { // 生成dao对象的办法
	return &dao{db: db}
}
