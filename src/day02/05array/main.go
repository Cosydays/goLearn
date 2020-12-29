package main

import "fmt"

//数组

//存放元素的容器
//必须指定存放的元素的类型和容量（长度）
//数组的长度是数组类型的一部分

func main() {
	var a1 [3]bool //[true false true]
	var a2 [4]bool //[true true false false]
	fmt.Printf("a1:%T, a2:%T\n", a1, a2)

	//数组的初始化
	//如果不初始化，默认元素都是零值（布尔值：false，整型和浮点型都是0，字符串：”“）
	fmt.Println(a1, a2)
	//初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//初始化方式2：根据初始值自动推断数组的长度是多少
	a10 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a10)
	//初始化方式3：根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	//数组的遍历
	citys := [...]string{"beijing", "shanghai", "shenzhen"} //索引：0~2 citys[0],citys[1],citys[2]
	//1.根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//2.for range遍历
	for _, v := range citys {
		fmt.Println(v)
	}

	//多维数组
	//[[1 2] [3 4] [5 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	//多维数组的遍历1
	for _, v1 := range a11 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	//多维数组的遍历2
	for i := 0; i < len(a11); i++ {
		for j := 0; j < len(a11[0]); j++ {
			fmt.Println(a11[i][j])
		}
	}

	//注意：数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1) //[1, 2, 3]
	fmt.Println(b2) //[100, 2, 3]
}
