package main

import (
	"fmt"
	"unicode/utf8"
)

/**
map学习  day2
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

	//寻找最长不含有重复字符的字串  LeetCode find the longest substring without repeating characters

	//思路
	//对于每一个字母x
	//1.lastOccured[x]不存在 或者< start  无需操作
	//2.lastOccured[x]>=start 更新start
	//3.更新lastOccured[x] 更新maxLength

	fmt.Println(lengthOfNonRepeatingSubStr("aaaccbdd"))
	fmt.Println(lengthOfNonRepeatingSubStr("慕课网")) //这里还暂时不支持中文
	fmt.Println("<<==================================================>>")
	fmt.Println(lengthOfNonRepeatingSubStrEN("你好！")) //支持中英文

	//rune相当于go 里面的char

	fmt.Println("==================================================")

	li := "yes我爱慕课网!"
	fmt.Println(len(li))           //长度19 3+3*5+1=19个字节
	fmt.Printf("%s\n", []byte(li)) //中文 UTF-8
	fmt.Printf("%X\n", []byte(li)) //16进制
	fmt.Println("==================================================")
	for _, b := range []byte(li) {
		//ASCII码  79--y  65--e  73--s  ....  每个中文是三字节
		fmt.Printf("%X ", b)
		fmt.Println()
	}

	fmt.Println()
	//
	for i, ch := range li { //ch is a rune (char)   int32  4字节
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count ", utf8.RuneCountInString(li))

	bytes := []byte(li)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	for i, ch := range []rune(li) {
		fmt.Printf("(%d,%c)", i, ch)
	}
	fmt.Println()

	//使用range遍历pos rune对
	//使用utf8.RuneCountInString获得字符数量
	//使用len获得字节长度
	//使用[]byte获得字节

	//字符串操作
	//Fields Split Join
	//Contains Index
	//ToLower ToUpper
	//Trim TrimRight TirmLeft
	//......

}

//不支持中文
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

//支持中英文  国际化操作
func lengthOfNonRepeatingSubStrEN(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
