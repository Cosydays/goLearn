package main

import "fmt"

//panic 和 recover

func funcA() {
	fmt.Println("a")
}

//1.recover()必须搭配defer使用
//2.defer一定要在可能引发panic的语句之前定义
func funcB() {
	//常用的例子
	//刚刚打开数据库连接
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("Error!!!") //程序崩溃退出
	fmt.Println("b")
}

func funcC() {

	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
