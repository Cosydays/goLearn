package main

import (
	"encoding/json"
	"fmt"
)

//复习结构体

type tmp struct {
	x int
	y int
}

type person struct {
	name string
	age  int
}

func sum(x int, y int) int {
	return x + y
}

//构造（结构体变量的）函数，返回值是对应结构体类型
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//方法
//接收者使用对应类型的首字母小写
//指定了接收者之后，只有接收者这个类型的变量才能调用这个方法
func (p person) dream(str string) {
	fmt.Printf("%s的梦想是%s\n", p.name, str)
}

func (p person) guonian0() {
	p.age++ //此处的p是p1的副本，所以修改的是副本
}

//指针接收者
//1.需要修改结构体变量的值时要使用指针接收者
//2.结构体本身比较大，拷贝的内存开销比较大时也要使用指针接收者
//3.保持一致性：如果有一个方法使用了指针接收者，其他的方法为了统一也要使用指针接收者
func (p *person) guonian1() {
	p.age++
}

func main() {
	var b = tmp{
		10,
		20,
	}
	fmt.Println(b)

	var a = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(a)

	var x int
	y := int8(10)
	fmt.Println(x, y)

	//结构体实例化1.
	var p1 person
	p1.name = "caijiahui"
	p1.age = 28

	//结构体实例化2.
	p2 := person{"caspar", 18}
	p3 := person{"jihui", 20}

	//结构体实例化3.
	//调用构造函数生成person类型变量
	p4 := newPerson("cai", 18)
	fmt.Println(p1, p2, p3, p4)

	p1.dream("挣钱")
	p2.dream("学习")

	fmt.Println(p1.age)
	p1.guonian0()
	fmt.Println(p1.age)

	p1.guonian1() //实际上是(&p1).guonian1()
	fmt.Println(p1.age)

	//结构体嵌套
	type addr struct {
		province string
		city     string
	}

	type student struct {
		name string
		addr //匿名嵌套别的结构体，就使用类型做名称
	}

	type point struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	po1 := point{100, 200}

	//序列化
	b1, err := json.Marshal(po1) //Go语言中把错误当成值返回，错误通常作为第二个返回值
	//如果出错了
	if err != nil {
		fmt.Printf("marshal failed, err%v\n", err)
	}
	fmt.Println(string(b1))

	//反序列化：由字符串 --> Go中的结构体变量
	str1 := `{"x":10, "y":20}`
	var po2 point //造一个结构体变量，准备接受反序列化的值
	//反序列化时要传递指针
	err = json.Unmarshal([]byte(str1), &po2)
	if err != nil {
		fmt.Printf("unmarshal failed, err%v\n", err)
	}

	//map
	m1 := map[int64]string{
		10001: "hhh",
		10010: "heiheihei",
		10086: "hehehe",
	}
	name0 := m1[10010]
	fmt.Println(name0)     //取不到就返回value类型的零值
	name1, ok := m1[20000] //ok=true表示有这个key，ok=false表示没有这个key
	fmt.Println(name1, ok)

	for k, v := range m1 {
		fmt.Println(k, v)
	}
	for k := range m1 {
		fmt.Println(k)
	}

}
