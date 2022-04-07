/*
Author:ydy
Date:
Desc:
*/
package main

import (
	"context"
	"fmt"
	"github.com/bingtianyiyan/youki_gotools/commonexternal/testx/internal/db"
	"log"
)

type Tracer struct {
	ch   chan string
	stop chan struct{}
}

func NewTracer() *Tracer{
	return &Tracer{
		ch : make(chan string,10),
		stop: make(chan struct{},1),
	}
}

func (self *Tracer) Event(ctx context.Context,msg string) error{
	select {
	case self.ch <- msg:
		return nil
	case <- ctx.Done():
        return ctx.Err()
	}
}

func (self *Tracer) Run(){
	for data := range self.ch{
		fmt.Println(data)
	}
	self.stop <- struct{}{}
}

func (self *Tracer) ShutDown(ctx context.Context){
	close(self.ch)
	select {
	case <- self.stop:
	case <- ctx.Done():
		fmt.Println("trace stop")
		break
	}
}

//func main()  {
//   var tracer = NewTracer()
//   go tracer.Run()
//   tracer.Event(context.Background(),"ev1")
//	tracer.Event(context.Background(),"ev2")
//   time.Sleep(time.Second *3)
//    ctx,cancel := context.WithDeadline(context.Background(),time.Now().Add(time.Second *3))
//    defer cancel()
//    tracer.ShutDown(ctx)
//
//}

//type App struct { // 最终须要的对象
//	db *sql.DB
//}
//
//func NewApp(db *sql.DB) *App {
//	return &App{db: db}
//}

//func main() {
//	app, err := InitApp() // 应用wire生成的injector办法获取app对象
//	if err != nil {
//		log.Fatal(err)
//	}
//	var version string
//	row := app.db.QueryRow("SELECT VERSION()")
//	if err := row.Scan(&version); err != nil {
//		log.Fatal(err)
//	}
//	log.Println(version)
//}

type App struct {
	dao db.Dao // 依赖Dao接口
}

func NewApp(dao db.Dao) *App { // 依赖Dao接口
	return &App{dao: dao}
}

func main() {
	app,  err := InitApp()
	if err != nil {
		log.Fatal(err)
	}
	version, err := app.dao.Version() // 调用Dao接口办法
	if err != nil {
		log.Fatal(err)
	}
	log.Println(version)
}
