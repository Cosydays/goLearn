package main

import "fmt"

//切片slice

func main() {
	//1.切片的定义
	var s1 []int    //定义一个存放int类型元素的切片
	var s2 []string //定义一个存放string类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s1 == nil)
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"beijing", "shanghai", "guangzhou"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s1 == nil)
	//切片的长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s2))
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s2), cap(s2))

	//2.由数组得到切片
	//1）切片指向了一个底层的数组
	//2）切片的长度就是它元素的个数
	//3）切片的容量是底层数组从切片的第一个元素到最后一个元素的数量
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] //基于一个数组切割，左包含右不包含，左闭右开
	fmt.Println(s3)
	s5 := a1[:4] //=>[0:4]
	fmt.Println(s5)
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	s6 := a1[3:] //=>[3:len(a1)] [7, 9, 11, 13]
	fmt.Println(s6)
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))
	s7 := a1[:] //=>[0:len(a1)]
	fmt.Println(s7)
	fmt.Printf("len(s7):%d cap(s7):%d\n", len(s7), cap(s7))
	//切片再切割
	s8 := s6[3:]
	fmt.Println(s8)
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))

	//注意：切片是引用类型，都指向了底层的一个数组。
	a1[6] = 1300 //修改了底层数组的值
	fmt.Println(s6)
	fmt.Println(s8)
}
