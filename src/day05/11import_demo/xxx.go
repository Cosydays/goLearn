package main

//包的路径从GOPATH/src后面的路径开始写起，路径分隔符用"/"
//想被别的包调用的标识符首字母大写
//单行导入和多行导入
//导入包的时候可以指定别名
//导入包不想使用包内部的标识符，需要使用匿名导入
//每个包导入的时候会自动执行一个名为init()的函数，它没有参数也没有返回值也不能手动调用

import (
	calc "day05/10calc"
	"fmt"
)

//1.全局变量和常量先声明
var x = 100

const pi = 3.14

//2.然后执行init()函数
func init() {
	fmt.Println("自动执行！")
	fmt.Println(x, pi)
}

//3.最后执行main函数
func main() {
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}
