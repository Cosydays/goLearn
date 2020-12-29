package main

import (
	"fmt"
)

func demo0(name string) {
	fmt.Println("Hello, ", name)
}

func demo1(f func(string), name string) {
	f(name)
}

func demo2() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

//闭包
//通过闭包的打包，
//把一个需要传参的函数打包成一个不需要传参的函数
func bi(f func(string), name string) func() {
	return func() {
		f(name)
	}
}

func main() {
	demo1(demo0, "caijiahui")
	ret := demo2()
	fmt.Printf("%T\n", ret)
	sum := ret(10, 20)
	fmt.Println(sum)

	//通过闭包的打包，
	//把一个需要传参的函数打包成一个不需要传参的函数
	demo3 := bi(demo0, "caijiahui")
	demo3()
}
