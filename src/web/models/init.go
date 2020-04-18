package models

import (
	"github.com/astaxie/beego/orm"
)

//数据库连接初始化
func init() {
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:123456@tcp(118.190.58.175:13306)/LICSLAN?charset=utf8")

	//初始化menumodel
	orm.RegisterModel(new(MenuModel))
}
