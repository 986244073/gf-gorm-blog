package db

import (
	"gf-app/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/jinzhu/gorm"
	"os/user"
	"time"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

var db *gorm.DB
var err error

func InitDb() {
	link := g.Cfg().Get("database.gorm")
	//fmt.Println(g.Cfg().Get("database.link"))
	//	"user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	//url:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",)
	db, err = gorm.Open("mysql", link)
	if err != nil {
		panic("连接失败!")
	}
	db.SingularTable(true)
	db.AutoMigrate(&user.User{},&model.Article{},&model.Category{})
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(10 * time.Second)
	db.DB().SetMaxIdleConns(10)

}
