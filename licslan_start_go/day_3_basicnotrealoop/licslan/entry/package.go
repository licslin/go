package main

import ".."
/**
封装
*/
func main() {

	//名字一般使用CamelCase
	//首字母大写 public  针对包来说的
	//首字母小写 private 针对包来说的
	//每个目录一个包
	//main包包含一个可执行的入口
	//为结构定义的方法必须放在同一个包内
	//可以是不同的文件

	//要扩展一个已有的方法 取消了继承
 	//如何扩充系统类型或者别人的类型   ---定义别名  使用组合

}


type myTreeNode struct {
	//如何引入licslan包下面的objectnew.go结构体呢？
	node *licslan.TreeNode
}
