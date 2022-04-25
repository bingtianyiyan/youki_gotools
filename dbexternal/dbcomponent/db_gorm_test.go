package dbcomponent

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

var (
	myDb *gorm.DB
)

func init(){
	myDb = gormMysql()
}
func gormMysql() *gorm.DB {
	var m = Mysql{
		Path: "127.0.0.1",
		Port: "3306",
		Username: "root",
		Password: "root",
		Dbname: "test1",
		MaxIdleConns: 100,
		MaxOpenConns: 100,
		Config: "parseTime=true",
	}//这边要修改成配置文件初始化后的对象
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), Gorm.Config()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

type User struct {
	gorm.Model
	Name string
}

// `Profile` 属于 `User`， `UserID` 是外键
type Profile struct {
	gorm.Model
	UserID  uint
	User   User `gorm:"FOREIGNKEY:id;ASSOCIATION_FOREIGNKEY:ID"`
	Name   string
}

func TestMigrate(t *testing.T){
	allModels := []interface{}{&User{}, &Profile{}}
	err := myDb.AutoMigrate(allModels...)
	fmt.Println(err)
}

func TestCreate(t *testing.T){
    var user = User{
    	Name: "jack",
	}
	var profile = Profile{
		Name: "jack-pr",
		User: user,
	}
	myDb.Begin()
	//批量插入用数组也可以
	var result = myDb.Create(&user)
	if result.Error != nil{
		fmt.Println(result.Error)
		return
	}
	profile.UserID = user.ID
	result = myDb.Create(&profile)
	if result.Error != nil{
		fmt.Println(result.Error)
		return
	}
	myDb.Commit()
}

func TestBelongsAssociate(t *testing.T){
	var user User
	var profile Profile
	myDb.Find(&profile,"id = ?",1)
	fmt.Println(profile)
	pointerOfUser := &profile
	//查询关联数据根据Profile的User
	if err := myDb.Model(&pointerOfUser).Association("User").Find(&user); err != nil {
		t.Errorf("failed to query users, got error %#v", err)
	}
	//追加关联
	var user2 User
	user2.Name="jack-appebd"
	if err := myDb.Model(&pointerOfUser).Association("User").Append(&user2); err != nil {
		t.Errorf("failed to query users, got error %#v", err)
	}

}

func TestQuery(t *testing.T){
 var result User
 //单查询
 myDb.Where("id = ?",2).Select("name").Find(&result)
 //多条以及分页(limit+offset
 var manyResult []User
 myDb.Order("id desc").Limit(2).Offset(1).Find(&manyResult)

 //sql多表查询
	var manyResult2 []User
 myDb.Raw("select a.* from users a inner join profiles b on a.id=b.user_id").Find(&manyResult2)
}

func TestUpdate(t *testing.T){
	var user User
	myDb.Find(&user).Where("id = ?",1)
	myDb.Table("users").Model(&user).Update("name","tom-up2")

	//方法2
	myDb.Table("users").Where("id = ?",1).Update("name","tom-up4")
	//原生sql  exec
}

//把相关联的信息也加载出来
func TestPreload(t *testing.T){
	var user = new(User)
	myDb.Table("profiles").Preload("User").Where("id=1").Find(&user)
}

func TestTransaction(t *testing.T){
	//方法1
	myDb.Transaction(func(tx *gorm.DB) error {
		return nil
	})

	//方法2
	myDb.Begin()
	myDb.Commit()
	myDb.Rollback()
}

func TestSession(t *testing.T){
	var result User
	//DryRun不执行Sql
	//Session是创建一个新DB
	myDb.Session(&gorm.Session{DryRun: true}).Where("id = ?",2).Select("name").Find(&result)
 
}
