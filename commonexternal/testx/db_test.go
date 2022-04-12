package testx

import (
	"fmt"
	"github.com/iris-contrib/jade/testdata/imp/model"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"testing"
)

func TestGormDemo(t *testing.T){
	// gorm配置
	gormConf := &gorm.Config{}
	// 连接数据库
	if err := simple.OpenDB("root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local", gormConf, 10, 20, Models...); err != nil {
		logrus.Error(err)
	}
	var data = Get(simple.DB(),38)
	fmt.Println(data)

	//新增
	//var insertData = &User{Name: "gorm",Age: 20}
	//Insert(simple.DB(),insertData)

	//更新
    //data.Name="gorm-up"
    //Update(simple.DB(),data)

    //根据列更新
    //var mapCl  map[string]interface{} = make(map[string]interface{})
    //mapCl["Name"] = "toms-up"
    //Updates(simple.DB(),38,mapCl)

    //分页
	var cnd = simple.NewSqlCnd()
	cnd.Like("name","jack")
	cnd.Page(0,2)
	cnd.Desc("id")
	list, paging := FindPageByCnd(simple.DB(),cnd)
	fmt.Println(list)
	fmt.Println(paging)
}

func  Get(db *gorm.DB, id int64) *User {
	ret := &User{}
	if err := db.Table("user").First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func Insert(db *gorm.DB,u *User){
	var err = db.Table("user").Create(u).Error
	if err != nil{
		fmt.Println(err)
	}
}

func Update(db *gorm.DB,u *User){
	var err =db.Table("user").Save(u).Error
	if err != nil{
		fmt.Println(err)
	}
}

func  Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Table("user").Model(&User{}).Where("id = ?", id).Updates(columns).Error
	return
}

func UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Table("user").Model(&model.User{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []User, paging *simple.Paging) {
	return FindPageByCnd(db, &params.SqlCnd)
}

func  FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []User, paging *simple.Paging) {
	cnd.Find(db.Table("user"), &list)
	count := cnd.Count(db.Table("user"), &User{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

type UserController struct {
	Ctx iris.Context
}

var Models = []interface{}{
	&User{},
}

type User struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Name         string `gorm:"type:text;size:20;" json:"name" form:"name"`
	Age            int   `json:"age" form:"age"`            // 年龄
}

