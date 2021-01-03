package main

import "fmt"

//接口的实现
//接口保存的分为两部分：动态类型+动态值
type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫动！")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡动！")
}

func (c chicken) eat(food string) {
	fmt.Printf("鸡吃%s...\n\n", food)
}

func main() {
	var a1 animal //定义一个接口类型的变量
	fmt.Printf("%T\n", a1)

	bc := cat{ //定义一个cat类型的变量bc
		name: "淘气",
		feet: 4,
	}
	a1 = bc
	a1.eat("小黄鱼")
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)

	kfc := chicken{ //定义一个chicken类型的变量kfc
		feet: 2,
	}
	a1 = kfc
	fmt.Printf("%T\n", a1)
}
