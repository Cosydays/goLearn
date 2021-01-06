package main

import (
	"fmt"
	"strings"
)

//字符串
//字符串不能修改

//byte和rune类型
//是类型别名

//字符串、字符、字节
//字符串：双引号包裹的是字符串
//字符：单引号包裹的是字符，单个字母、单个符号、单个文字
//字节：1byte=8bit
//go语言中字符串都是UTF8编码，UTF8编码中一个常用汉字一般占用3个字节

func main() {
	path := "\"C:\\Users\\Cosydays\\go\\src\\github.com\\day01\\08string\\xxx.go\""
	fmt.Printf("%v\n", path)
	s := "I'm ok"
	fmt.Println(s)

	//多行字符串
	s2 := `abc
			efg`
	fmt.Println(s2)
	fmt.Println(s2)
	s3 := `C:\Users\Cosydays\go\src\github.com\day01\08string\main.go`
	//字符串相关操作
	fmt.Println(len(s2))
	//字符串拼接
	name := "caijiahui"
	word := "shuai"
	fmt.Println(name + word)
	ss1 := fmt.Sprintf("%s%s", name, word)
	fmt.Println(ss1)
	//字符串分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(ss1, "caijiahui"))
	//前缀
	fmt.Println(strings.HasPrefix(ss1, "cai"))
	//后缀
	fmt.Println(strings.HasSuffix(ss1, "cai"))
	//子串出现的位置
	fmt.Println(strings.Index(ss1, "jiahui"))
	fmt.Println(strings.LastIndex(ss1, "i"))
	//拼接
	fmt.Println(strings.Join(ret, "+"))
}

//原字体大小为13
