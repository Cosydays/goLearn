package main

import "fmt"

//结构体嵌套

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address //匿名嵌套结构体
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "caspar",
		age:  9000,
		address: address{
			province: "shandong",
			city:     "weihai",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.address.city)
	fmt.Println(p1.address.city) //先在自己几条狗体找这个字段，找不到就去匿名嵌套的结构体中查找
}
