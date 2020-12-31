package main

import "fmt"

//结构体是值类型

type person struct {
	name   string
	gender string
}

//go语言中函数传参永远是拷贝
func f1(x person) {
	x.gender = "nv" //修改的是副本的gender
}

func f2(x *person) {
	//(*x).gender = "nv" //根据内存地址找到那个原变量，修改的就是原来的变量
	x.gender = "nv" //语法糖，自动根据指针找到对应的变量
}

func main() {
	var p person
	p.name = "caspar"
	p.gender = "nan"
	f1(p)
	fmt.Println(p.gender)
	f2(&p)
	fmt.Println(p.gender)

	//结构体指针1
	var p2 = new(person)
	p2.name = "caspar"
	p2.gender = "nan"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)  //p2保存的值就是一个内存地址
	fmt.Printf("%p\n", &p2) //求p2的内存地址

	//结构体指针2
	//2.1key-value初始化
	var p3 = &person{
		name:   "jiahui",
		gender: "nan",
	}
	fmt.Printf("%#v\n", p3)
	//2.2使用值列表的形式初始化，值的顺序要和结构体定义时字段的顺序一致
	p4 := &person{
		"xiaowangzi",
		"nan",
	}
	fmt.Printf("%#v\n", p4)
}
