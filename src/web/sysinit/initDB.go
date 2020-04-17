package sysinit

import (
	_ "../models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func initDB()  {
	dbAlias:=beego.AppConfig.String("db_alias")
	dbName:=beego.AppConfig.String("db_name")
	dbUser:=beego.AppConfig.String("db_user")
	dbPwd:=beego.AppConfig.String("db_password")
	dbHost:=beego.AppConfig.String("db_host")
	dbPort:=beego.AppConfig.String("db_port")
	dbCharset:=beego.AppConfig.String("db_charset")

	//注册数据库信息
	orm.RegisterDataBase(
		dbAlias,
		"mysql",
		dbUser+":"+
		dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+
		"?charset="+dbCharset,30)
	//如果是开发模式 则显示m命令信息
	isDEv:= beego.AppConfig.String("runmode")=="dev"

	//自动建表  根据struct自动创建表   orm.RegisterModel(new(Page))  好比这段代码的功能  根据"对象"创建相应的表
	orm.RunSyncdb("default",false,isDEv)

	if isDEv{
		orm.Debug=isDEv
	}
}
