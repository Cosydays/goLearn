package main

import "fmt"

func main() {
	s := "Hello Caspar"
	//len()求的是byte字节的数量
	n := len(s)
	fmt.Println("s中字节的数量为：", n)

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		fmt.Printf("%c\n", s[i])
	}

	//打印字符串中的每个字符
	for _, c := range s {
		fmt.Printf("%c \n", c)
	}

	//"hello"
	//字符串修改
	s2 := "白萝卜"      // [白 萝 卜]
	s3 := []rune(s2) //把字符串强制转成了一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3)) //把rune切片强制转换成字符串

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := "H" //string
	c4 := 'H' //int32
	fmt.Printf("c1:%T c2:%T\n", c3, c4)
	fmt.Printf("%d\n", c4)

	//类型转换
	n1 := 10 //int类型
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
