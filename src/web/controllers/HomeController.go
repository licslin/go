package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"web/models"
)

//TODO
type HomeController struct {
	beego.Controller
}
func (c *HomeController) Get() {
	/*c.Data["Website"] = "beego.me"
	c.Data["Email"] = "licslan@sina.com"*/

	//测试一下session
	c.SetSession("licslan","LICSLAN")
	user := c.GetSession("licslan")
	fmt.Println("session is ",user)

	logs.Informational("user licslan loged in")

	//先更新一下再查询
	models.UpdatePage()
	//改造为从model里面获取
	page := models.GetPage()
	c.Data["Website"] = page.Website
	c.Data["Email"] = page.Email
	c.TplName = "index.tpl"
}



