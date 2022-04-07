/*
Author:ydy
Date:
Desc:
*/
package config

import (
	"encoding/json"
	"github.com/google/wire"
	"os"
)

var Provider = wire.NewSet(New) // 将New办法申明为Provider，示意New办法能够创立一个被他人依赖的对象,也就是Config对象

type Config struct {
	Database database `json:"database"`
}

type database struct {
	Dsn string `json:"dsn"`
}

func New() (*Config, error) {
	//file, _ := exec.LookPath(os.Args[0])
	//path, _ := filepath.Abs(file)
	//fmt.Println(path)
	//index := strings.LastIndex(path, string(os.PathSeparator))
	//path = path[:index]
	//fp, err := os.Open("C:/Code/youki_gotools/commonexternal/testx/config/app.json")
	fp, err := os.Open("../config/app.json")
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	var cfg Config
	if err := json.NewDecoder(fp).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
