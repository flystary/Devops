package main

import (
	"fmt"
	"strings"
)

// Compare 函数，用于比较两个字符串的大小，如果两个字符串相等，返回为 0。如果 a 小于 b ，返回 -1 ，反之返回 1 。不推荐使用这个函数，直接使用 == != > < >= <= 等一系列运算符更加直观。
//func Compare(a, b string) int

//EqualFold 函数，计算 s 与 t 忽略字母大小写后是否相等。
//func EqualFold(s, t string) bool

func main() {

	a := "gopher"
	b := "hello world"

	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Compare(a, a))
	fmt.Println(strings.Compare(b, a))

	fmt.Println(strings.EqualFold("GO", "go"))
	fmt.Println(strings.EqualFold("壹", "一"))

	//是否存在某个字符或子串

	// 子串 substr 在 s 中，返回 true
	//func Contains(s, substr string) bool
	// chars 中任何一个 Unicode 代码点在 s 中，返回 true
	//func ContainsAny(s, chars string) bool
	// Unicode 代码点 r 在 s 中，返回 true
	//func ContainsRune(s string, r rune) bool

	fmt.Println("----------------------")
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("in failure", "s g"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))

}
