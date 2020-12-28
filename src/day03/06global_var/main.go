package main

import "fmt"

//变量的作用域

var x = 100 //定义一个全局变量

//定义一个函数
func f1() {
	//函数中查找变量的顺序
	//1.现在函数内部查找
	//2.找不到往函数外卖呢查找，一直找到全局

	//函数内部作用域
	name := "caijiahui"
	fmt.Println(name)
	fmt.Println(x)
}

func main() {
	f1()

	//语句块作用域
	if i := 10; i < 18 {
		fmt.Println("go to school!")
	}
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
}
