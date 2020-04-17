package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// go语言model就是一个结构体  struct
type Page struct {
	//mvc  model
	Id int
	Website string
	Email string
}

//初始化
func init()  {
	//aliasName, driverName, dataSource string, params ...int  注册一下数据库
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:123456@tcp(118.190.58.175:13306)/LICSLAN?charset=utf8")
	orm.RegisterModel(new(Page))
}

//查询数据
func GetPage() Page  {
	/*result:=Page{Website:"www.licslan.com",Email:"licslan@sina.com"}
	return result*/
	//可以先更新一下
	//
	res := orm.NewOrm()
	p:=Page{Id:1}
	err := res.Read(&p)
	if err!=nil{
		fmt.Println("the result has some problem",err)
		panic(err)
	}
	return p
}
//更新数据
func UpdatePage()  {
	p:=Page{Id:1,Email:"new Email for update"}
	res := orm.NewOrm()
	//只改Email  cols set the columns those want to update.
	res.Update(&p,"Email")
	//res.Insert()  新增操作
	//res.Delete()  删除操作
}