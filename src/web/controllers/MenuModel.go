package controllers

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
