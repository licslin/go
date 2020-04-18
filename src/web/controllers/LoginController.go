package controllers

import "C"
import (
	"github.com/astaxie/beego"
	"strings"
	"web/models"
	"web/utils"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Index() {
	if c.Ctx.Request.Method == "POST" {
		userKey := strings.TrimSpace(c.GetString("username"))
		password := strings.TrimSpace(c.GetString("password"))
		if len(userKey) > 0 && len(password) > 0 {
			password := utils.Md5([]byte(password))
			user := models.GetUserByName(userKey)
			if password == user.PassWord {
				//放到session里面去
				c.SetSession("xcmsuser", user)
				c.Redirect("/", 302)
				c.StopRun()
			}
		}
	}
	C.TplName = "login/index.html"
}
