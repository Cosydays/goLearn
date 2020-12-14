package main

import "fmt"

//整型
//无符号整型：uint8、uint16、uint32、 uint64
//无符号整型：int8、int16、int32、 int64
//uint和int：具体是32位还是64位
//uintptr：表示指针

//浮点型
//float64和float32
//Go语言中浮点数默认是float64

//复数
//complex128和complex64

//布尔值
//true和flase
//不能和其他的类型做转换

//其他进制数
//Go语言中没办法直接定义二进制数
//八进制加0，十六禁止加0x

func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1)
	fmt.Printf("%o\n", i1)
	fmt.Printf("%x\n", i1)
	//八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	//十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	//查看变量的类型
	fmt.Printf("%T\n", i3)

	//声明int8类型的变量
	i4 := int8(9) //明确指定int8类型，否则就是默认为int类型
	fmt.Printf("%T\n", i4)
}

