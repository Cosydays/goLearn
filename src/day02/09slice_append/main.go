package main

import (
	"fmt"
)

//append() 为切片追加元素

func main() {
	s1 := []string{"beijing", "shanghai", "shenzhen"}
	//s1[3] = "guangzhou" //错误的写法 会导致编译错误：索引越界
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	//注意：调用append函数必须用原来的切片变量接收返回值
	//append追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
	s1 = append(s1, "guangzhou")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "hangzhou", "chengdu")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := []string{"wuhan", "xian", "suzhou"}
	s1 = append(s1, ss...) //...表示拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

}
