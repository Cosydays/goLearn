package main

import (
	"fmt"
)

//运算符

func main() {
	var (
		a = 5
		b = 2
	)

	//算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句，不能放在等号的右边赋值 ==》 a=a+1
	b-- //单独的语句，不能放在等号的右边赋值 ==》 b=b-1
	fmt.Println(a)

	//关系运算符
	fmt.Println(a == b) //Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b)
	fmt.Println(a >= b)

	//逻辑运算符
	//如果年龄大于18岁并且年龄小于60岁
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("上班！")
	} else {
		fmt.Println("不上班！")
	}
	//如果年龄小于18岁 或者 年龄大于60岁
	if age < 18 || age > 60 {
		fmt.Println("不用上班！")
	} else {
		fmt.Println("上班！")
	}
	//位运算符：针对的是二进制数
	//&：按位与（两位均为1才为1）
	fmt.Println(5 & 2) //0
	//|：按位或（两位有一个为1就为1）
	fmt.Println(5 | 2) //7
	//^：按位异或（两位不一样则为1）
	fmt.Println(5 ^ 2) //7
	//<<：将二进制位左移指定位数
	fmt.Println(5 << 1)  //1010=>10
	fmt.Println(1 << 10) //10000000000=>1024
	//>>：将二进制位右移指定位数
	fmt.Println(5 >> 1) //10=>2
	var m = int8(1)     //只能存八位
	fmt.Println(m << 10)

	//赋值运算符，用来给变量赋值的
	x := 10
	x += 1  //x = x + 1
	x -= 1  //x = x - 1
	x *= 2  //x = x * 2
	x /= 2  //x = x / 2
	x %= 2  //x = x % 2
	x <<= 2 //x = x << 2
	x >>= 2 //x = x >> 2
	x &= 2  //x = x & 2
	x |= 2  //x = x | 2
	x ^= 2  //x = x ^ 2
	fmt.Println(x)
}
