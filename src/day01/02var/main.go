package main

import (
	"fmt"
)

/*变量*/

//Go语言中推荐使用驼峰式命名
//var studentName string

//三种声明变量的方式
//1.var name1 string
//2.var name2 = "jiaHui"
//3.函数内部专属：name3 := "jiaHui"

//匿名变量（哑元变量）
//当有些数据必须用变量接受但是又不使用它时，就可以用_来接收这个值

//批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "caijiahui"
	age = 18
	isOk = true
	// Go语言中非全局变量声明必须使用，不使用就编译不过去
	// var nouse string
	fmt.Print(isOk) // 在终端中输出要打印的内容
	fmt.Println()
	fmt.Printf("name:%s\n", name) // %s：占位符，使用name这个变量的值去替换占位符
	fmt.Println(age)              // 打印完指定的内容之后会在后面自动加一个换行符
	//nouse = "caspar"
	var s1 string = "caijiahui"
	fmt.Println(s1)
	//类型推导（根据值判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)
	//短变量声明：在函数内部，可以使用更简略的:=方式声明并初始化变量
	s3 := "hhhhh"
	fmt.Println(s3)
	//同一个作用域中不能重复声明同名的变量
	//s1 := "10"
	//匿名变量是一个特殊的变量，用来忽略返回值
}
