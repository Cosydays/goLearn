package main

import "fmt"

//单行注释

/*
多行注释
*/

////Go语言函数外的语句必须以关键字开头
//var name = "家辉"
//var age int
//age = 10

//const (
//	Num = 100
//)

//main函数是入口函数
//他没有参数也没有返回值

func main() {
	//函数内部定义的变量必须使用
	//var isOk = true
	//fmt.Println(isOk)

	//if循环
	//var age = 19
	//if age > 18 {
	//	fmt.Println("成年了！")
	//} else if age > 7{
	//	fmt.Println("上小学了！")
	//} else {
	//	fmt.Println("幼儿园！")
	//}

	//for.1
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	//for.2
	//var i = 0
	//for ; i < 10; i++ {
	//	fmt.Println(i)
	//}

	//for.3
	//var j = 0
	//for j < 10 {
	//	fmt.Println(j)
	//	j++
	//}

	//for.4
	//for {
	//	fmt.Println("无限循环")
	//}

	//for.5
	//s := "我叫蔡家辉"
	//fmt.Println(len(s))

	//for i, v := range s {
	//	fmt.Printf("%d %c\n", i, v)
	//}

	//哑元变量，不想用到的都直接扔个它
	//for _, v := range s {
	//	fmt.Println(v)
	//}

	//打印9*9乘法表
	//for i :=1; i < 10; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d*%d=%d\t", j, i, j*i)
	//	}
	//	fmt.Println()
	//}
	//var f1 = 1.234 //float64
	//fmt.Printf("%T", f1)
	//var f2 = float32(1.234)
	//fmt.Printf("%T", f2)

	//byte和rune
	//s1 := "Hello"
	//s2 := "你好"
	//for _, v := range s1 {
	//	fmt.Printf("%c\n", v)
	//}
	//for _, v := range s2 {
	//	fmt.Printf("%T\n", v)
	//}

	//八进制数表示方法（加0）
	var n1 = 0777
	//十六进制数表示方法（加0x）
	var n2 = 0xff
	fmt.Println(n1, n2)
	fmt.Printf("%o\n", n1)
	fmt.Printf("%x\n", n2)
}
