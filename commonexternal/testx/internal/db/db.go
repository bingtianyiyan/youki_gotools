/*
Author:ydy
Date:
Desc:
*/
package db

import (
	"database/sql"
	"github.com/bingtianyiyan/youki_gotools/commonexternal/testx/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

//var Provider = wire.NewSet(New) // 同理
var Provider = wire.NewSet(New, NewDao, wire.Bind(new(Dao), new(*dao)))

func New(cfg *config.Config) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", cfg.Database.Dsn)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	return db, nil
}
