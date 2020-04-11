package licslan

import (
	"fmt"
)

/**
"面向对象"  结构体学习  day3  包名由main改为了liclsan  小写私有 private  大写字母开头public
*/
func main() {

	//go语言仅支持封装
	//不支持继承和多态
	//面向接口编程 go语言没有class  只有struct  结构体定义
	//go语言没有构造函数
	//结构创建在堆上还是栈上
	//c++ 局部变量栈上  出栈就没有了  堆上手动维护
	//java 对象基本上实在堆上  垃圾回收就行
	//go 不用知道在哪分配

	//申明一个结构体变量
	var root TreeNode
	root = TreeNode{value: 3}
	root.left = &TreeNode{}
	root.right = &TreeNode{5, nil, nil}
	root.right.left = new(TreeNode)
	root.right.left = CreateNode(2)

	//定义一个slice
	nodes := []TreeNode{
		{value: 3},
		{},
		{6, nil, &root}, //指针指向root
	}
	fmt.Println(nodes)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	root.print()
	fmt.Println()
	root.setValue(5)
	root.print()
	fmt.Println("~~~~~~~~~~~~vvvvvvvvvvvvv~~~~~~~~~~~~~~~~")

	var proot *TreeNode
	proot.setValue(200)
	proot = &root
	proot.setValue(300)
	proot.print()
	fmt.Println("~~~~~~~~~~~~mmmmmmmmm~~~~~~~~~~~~~~~~")
	proot.traverse()

}

//工厂函数  返回了局部变量的地址
func CreateNode(value int) *TreeNode {
	//返回局部变量地址   在C++里面就会有问题
	return &TreeNode{value: value}
}

//建立一棵树
type TreeNode struct {
	value int
	//定义属性  变量名在前类型在后
	left, right *TreeNode //指针  地址
}

//为结构体定义方法 和其他语言有很大的不同  node TreeNode 这里作为接收者
//go语言所有参数都是传值的  值接收者  值传递是不可以结构体里面元素属性的内容的
func (node TreeNode) print() {
	fmt.Print(node.value, " ")
}

//结构体方法设置  要修改元素的值  需要加指针的 *  指针接收者
func (node *TreeNode) setValue(value int) {
	//指针接收者根据场景判断需不需要判断nil  如果是值接收者 就不可能是nil了
	if node == nil {
		fmt.Println("setting value to nil node. Ignored")
		return
	}
	node.value = value
}

//遍历这颗树  中序遍历
func (node *TreeNode) traverse() {

	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

//显示定义和命名方法接收者  只有使用指针才可以改变结构内容  nil指针也可以调用方法

//编译器很聪明  知道你是要指针地址   还是值

//值接收者 VS 指针接收者

//要改变内容的话必须使用指针接收者
//值接收者是一个拷贝  要考虑一个性能问题  结构过大也考虑使用指针  小的结构就无所谓了  拷贝就拷贝
//一致性：如有指针接收者  最好都是指针接收者  不要有的是指针有的是值
//值接收者是go语言特有的  指针接收者其他语言都有  c++ java <this指针>  python self
//值/指针接收者均可接收值/指针
