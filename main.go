package main

import (
	_ "app/routers"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

// 判断所给路径文件/文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func init() {
	orm.Debug = true

	dbpath := "/data/data.db?charset=utf8&loc=Asia%2FShanghai"
	if FileExists("=is_debug.tmp") {
		println("debug version.")
		dbpath = "data.db?charset=utf8&loc=Asia%2FShanghai"
	}

	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", dbpath)

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
