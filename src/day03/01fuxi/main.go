package main

import "fmt"

//day03复习

func main() {
	var name string
	name = "caijiahui"
	fmt.Println(name)

	//数组的初始化方式1
	var ages1 [30]int //声明了一个变量ages，它是[30]int类型
	ages1 = [30]int{1, 2, 3, 4, 6}
	fmt.Println(ages1)
	//数组的初始化方式2
	var ages2 = [...]int{1, 2, 3, 4, 6}
	fmt.Println(ages2)
	//数组的初始化方式3
	var ages3 = [...]int{1: 100, 99: 200} //索引为1的地方值为100，索引为99的地方值为200
	fmt.Println(ages3)

	//二维数组
	var a1 [3][2]int
	a1 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a1)
	//多维数组只有最外层可以使用...
	var a2 = [...][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a2)

	//数组是值类型
	x := [3]int{1, 2, 3}
	y := x     //把x的值拷贝了一份给了y
	y[1] = 200 //修改的是副本y，并不影响x

	fmt.Println(x) //[1, 2, 3]
	f1(x)
	fmt.Println(x) //[1, 2, 3]

	//切片slice
	var s1 []int //没有分配内存，==nil
	fmt.Println(s1)
	fmt.Println(s1 == nil)
	//切片初始化方式1
	s1 = []int{1, 2, 3}
	fmt.Println(s1)
	//切片初始化方式2
	//make初始化，分配内存
	s2 := make([]bool, 2, 4)
	fmt.Println(s2) //[false, false]
	s3 := make([]int, 0, 4)
	fmt.Println(s3 == nil)

	//切片不存值，它就像一个框，去底层数组框值
	//切片的扩容策略
	//1.如果申请的容量大于原来的2倍，那就直接扩容至新申请的容量
	//2.如果小于1024，那么就直接两倍
	//3.如果小于1024，就按照1.25倍去扩容
	//4.具体存储的值类型不同，扩容策略也有不同
	s4 := []int{1, 2, 3}
	s5 := s4
	fmt.Println(s5)
	s5[1] = 200
	fmt.Println(s4) //[1, 200, 3]
	fmt.Println(s5) //[1, 200, 3]

	//切片必须要初始化后才能使用=》make()函数
	var s6 []int
	s6 = make([]int, 1)
	s6[0] = 100
	fmt.Println(s6)

	//append会自动初始化切片
	var s7 []int
	s7 = append(s7, 1)
	fmt.Println(s7)

	//copy()函数
	s8 := []int{1, 2, 3}
	s9 := s8
	var s10 = make([]int, 3, 3)
	copy(s10, s8)
	s9[1] = 200
	fmt.Println(s8)  //[1, 200, 3]
	fmt.Println(s9)  //[1, 200, 3]
	fmt.Println(s10) //[1, 2 ,3]

	//指针
	//Go里面的指针只能读不能修改，不能修改指针变量指向的地址
	myName := "caijiahui"
	namePointer := &myName
	fmt.Println(namePointer)
	fmt.Printf("%T\n", namePointer) //*string
	nameV := *namePointer           //根据内存地址找值
	fmt.Println(nameV)

	//map
	var m1 map[string]int
	fmt.Println(m1  == nil) //true
	m1 = make(map[string]int, 10)
	m1["caijiahui"] = 24
	fmt.Println(m1)
	fmt.Println(m1["cai"]) //如果key不存在，返回的是对应类型的零值
	score, ok := m1["cai"]
	if !ok {
		fmt.Println("没有cai这个元素")
	} else {
		fmt.Println("cai的值为", score)
	}
	delete(m1, "jiahui") //删除的key不存在就什么都不干
	delete(m1, "caijiahui")
	fmt.Println(m1) //map[]
	fmt.Println(m1 == nil) //false，已经开辟了内存

}

func f1(a [3]int) {
	//Go语言中的函数传递的都是值
	a[1] = 100 //此处修改的是副本的值
}
