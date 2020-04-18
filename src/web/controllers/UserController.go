package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"web/consts"
	"web/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) index() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs.html"
	c.setTpl()
}

func (c *UserController) list() {
	page, err := c.GetInt("page")
	if err != nil {
		page = 1
	}
	if page <= 0 {
		page = 1
	}
	size, err := c.GetInt("limit")
	if err != nil {
		size = 20
	}
	result, count := models.UserList(size, page)
	c.listJsonResult(consts.JRCodeSucc, "ok", count, result)
}

func (c *UserController) add() {

}

func (c *UserController) addDo() {

}

func (c *UserController) Edit() {
	userId, _ := c.GetInt("userid")
	o := orm.NewOrm()
	var user = models.UserModel{UserId: userId}
	o.Read(&user)

	user.PassWord = ""
	c.Data["User"] = user

	authmap := make(map[int]bool)
	if len(user.AuthStr) > 0 {
		var authobj []int
		json.Unmarshal([]byte(user.AuthStr), &authobj)
		for _, v := range authobj {
			authmap[v] = true
		}
	}

	type Menuitem struct {
		Name    string
		Ischeck bool
	}

	menu := models.ParentMenuList()
	menus := make(map[int]Menuitem)
	for _, v := range menu {
		menu[v.Mid] = Menuitem{v.Name, authmap[v.Mid]}
	}
	c.Data["Menus"] = menus

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs_edit.html"
	c.setTpl("user/edit.html", "common/layout_edit.html")

}
func (c *UserController) EditDo() {

}

func (c *UserController) DeleteDo() {

}
