package main

import "fmt"

//结构体模拟是心啊其他语言中的”继承“

type animal struct {
	name string
}

//给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动！", a.name)
}

func main() {

}
