package models

import (
	. "github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
)

type MenuModel struct {
	Mid    int `orm:"pk;auto"`
	Parent int
	Name   string `orm:size(45)`
	Format string `orm:"size(2048);default({})"`
}

//树状结构
type MenuTree struct {
	MenuModel
	Child []MenuModel
}

//表名
func (m *MenuModel) TableName() string {
	return "xcms_menu"
}

func MenuStruct() map[int]MenuTree {
	query := NewOrm().QueryTable("xcms_menu")
	data := make([]*MenuModel, 0)
	//排序
	query.OrderBy("parent", "-seq").All(&data)

	var menu = make(map[int]MenuTree)
	if len(data) > 0 {
		for _, v := range data {
			if 0 == v.Parent {
				var tree = new(MenuTree)
				tree.MenuModel = *v
				//父节点
				menu[v.Mid] = *tree
			} else {
				//判断一下有没有这个父节点
				if tmp, ok := menu[v.Parent]; ok {
					//如果有 将子节点放到父节点里面来
					tmp.Child = append(tmp.Child, *v)
					menu[v.Parent] = tmp
				}
			}
		}
	}
	//返回这个树状结构
	return menu
}

//获取列表方法
func MenuList() ([]*MenuModel, int64) {
	query := NewOrm().QueryTable("xcms_menu")
	total, _ := query.Count()
	data := make([]*MenuModel, 0)
	query.OrderBy("parent", "-seq").All(&data)
	return data, total
}

func ParentMenuList() []*MenuModel {
	query := NewOrm().QueryTable("xcms_menu").Filter("parent", 0)
	data := make([]*MenuModel, 0)
	query.OrderBy("-seq").Limit(1000).All(&data)
	return data
}

func MenuFormatStruct(mid int) *simplejson.Json {
	menu := MenuModel{Mid: mid}
	err := NewOrm().Read(&menu)
	if nil == err {
		jsonstruct, err2 := simplejson.NewJson([]byte(menu.Format))
		if nil == err2 {
			return jsonstruct
		}
	}
	return nil

}
