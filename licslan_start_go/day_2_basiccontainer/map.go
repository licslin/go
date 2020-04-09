package main

import (
	"fmt"
)

/**
map学习
*/
func main() {
	//定义一个Map
	m := map[string]string{
		"name":   "licslan",
		"course": "go",
		"site":   "licslan.com",
		"good":   "not bad",
	}
	fmt.Println(m)
	fmt.Println(len(m))
	delete(m, "name")
	fmt.Println(m)
	fmt.Println(len(m))
	//空map
	m2 := make(map[string]int) //m2 == empty map
	fmt.Println(m2)
	var m3 map[string]int //m3==nil
	fmt.Println(m3)
	//遍历map
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k, _ := range m {
		fmt.Println(k)
	}
	for k := range m {
		fmt.Println(k)
	}
	//上面的key是无序的 hashMap

	//获取值
	courserName := m["course"]
	fmt.Println(courserName)

	courserName, ok := m["course"]
	fmt.Println(courserName, ok)

	//如果获取的值 通过一个错误的key去取  获取的值将会是一个zero value

	//看获取的结果是否存在
	courserNamexx, ok := m["xx"]
	fmt.Println(courserNamexx, ok)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	courserName, ok = m["x"]
	fmt.Println(courserName, ok)

	if courserName, ok := m["x"]; ok {
		fmt.Println(courserName)
	} else {
		fmt.Println("the key does not exist")
	}
	//key不存在时 获取value类型的初始值zero value
	//用value,ok:=m[key]来判断是否存在key

	//使用range遍历key  或者遍历key,value对
	//不保证遍历顺序 要顺序的话  需要手动对key排序

	//map使用哈希表 必须可以比较相等
	//map的key除了slice map function的内建类型都可以作为key
	//struct类型不包括上述字段 也可以作为key

	//寻找最长不含有重复字符的字串
}
