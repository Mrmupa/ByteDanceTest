package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type ConnInfo struct {
	MyUser string
	Password string
	Host string
	Port int
	Db string
}

func DbConn(MyUser, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser,Password, Host, Port, Db )
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	return db
}

type Admin struct {
	Id int
	Name string
	Account string
	Password string
}
type User struct {
	Id int
	Name string
	Account string
	Password string
}

func main() {
	cn := ConnInfo{
		"root",
		"Zzx123456",
		"127.0.0.1",
		3306,
		"test",
	}

	db := DbConn(cn.MyUser,cn.Password,cn.Host,cn.Db,cn.Port)
	db.AutoMigrate(&Admin{})
	defer db.Close() // 关闭数据库链接，defer会在函数结束时关闭数据库连接
	db.Create(Admin{Id: 5,Name: "h",Account: "new",Password: "2312312"})
	db.Delete(&Admin{})

}