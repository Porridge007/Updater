package main

import (
	"Updater/models"
	_ "Updater/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.1.95:3306)/updater_db?charset=utf8&loc=Asia%2FShanghai", 30)

	// register model
	orm.RegisterModel(new(models.File))
	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.UserToken))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
