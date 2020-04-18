package controllers

import (
	"github.com/astaxie/beego/orm"
	"web/consts"
	"web/models"
)

type FormatController struct {
	//继承
	BaseController
}

func (c *FormatController) Edit() {

	midvalue, _ := c.GetInt("mid")
	menu := models.MenuModel{Mid: midvalue}
	orm.NewOrm().Read(&menu)

	c.Data["Mid"] = midvalue
	c.Data["Format"] = menu.Format

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "format/footerjs_edit.html"

	c.setTpl("format/edit.html", "common/layout_edit.html")

}

func (c *FormatController) EditDo() {
	mid, _ := c.GetInt("mid")
	f := c.GetString("formatstr") //提交的format信息

	if 0 != mid {
		menu := models.MenuModel{Mid: mid, Format: f}
		mid, _ := orm.NewOrm().Update(&menu, "format")
		c.jsonResult(consts.JRCodeSucc, "ok", mid)
	}
	c.jsonResult(consts.JRCodeFailed, "", 0)

}
