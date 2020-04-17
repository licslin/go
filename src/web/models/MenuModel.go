package models

import "github.com/astaxie/beego/orm"

type MenuMmodel struct {

	Mid int `orm:"pk;auto"`
	Parent int
	Name string `orm:size(45)`
	Format string `orm:"size(2048);default({})"`
}

//树状结构
type MenuTree struct {
	MenuMmodel
	Child []MenuMmodel
}
//表名
func (m *MenuMmodel)TableName() string  {
	return "xcms_menu"
}

func MenuStruct() map[int]MenuTree {
	query:=orm.NewOrm().QueryTable("xcms_menu")
	data:=make([]*MenuMmodel,0)
	//排序
	query.OrderBy("parent","-seq").All(&data)

	var menu=make(map[int]MenuTree)
	if len(data)>0{
		for _,v:=range data{
			if 0==v.Parent{
				var tree=new(MenuTree)
				tree.MenuMmodel=*v
				//父节点
				menu[v.Mid]=*tree
			}else {
				//判断一下有没有这个父节点
				if tmp,ok:=menu[v.Parent];ok{
					//如果有 将子节点放到父节点里面来
					tmp.Child=append(tmp.Child,*v)
					menu[v.Parent]=tmp
				}
			}
		}
	}
	//返回这个树状结构
	return menu
}

//获取列表方法
func MenuList()([]*MenuMmodel,int64)  {
	query := orm.NewOrm().QueryTable("xcms_menu")
	total,_:=query.Count()
	data := make([]*MenuMmodel,0)
	query.OrderBy("parent","-seq").All(&data)
	return data,total

}
