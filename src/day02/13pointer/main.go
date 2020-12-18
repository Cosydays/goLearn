package main

import "fmt"

//pointer

func main() {

	//1.&:取地址
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Println(*p)
	fmt.Printf("%T\n", p) //*int:int类型的指针

	//2.*:根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m) //int

	//new函数申请一个内存地址
	var a1 *int
	fmt.Println(a1)
	var a2 = new(int)
	fmt.Println(a2) //地址
	fmt.Println(*a2) //0
	*a2 = 100
	fmt.Println(*a2) //100
}

