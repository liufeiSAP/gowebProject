package main

import (
	_ "gowebProject/routers"

	"github.com/astaxie/beego"

	//"fmt"
	"gowebProject/models"

	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

//引入数据模型
func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
	// 运行时
	beego.Run()
}
