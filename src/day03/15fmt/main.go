package main

import "fmt"

//fmt

func main()  {
	fmt.Print("shahe")
	fmt.Print("nazha")
	fmt.Println("------")
	fmt.Println("shahe")
	fmt.Println("nazha")
	//Printf("格式化字符串", 值)
	//%T:查看类型
	//%d:十进制数
	//%b:二进制数
	//%o:八进制数
	//%x:十六进制数
	//%c:字符
	//%s:字符串
	//%p:指针
	//%v:值
	//%f:浮点数
	//%t:布尔值

	var m1 = make(map[string]int, 1)
	m1["lixiang"] = 100
	fmt.Printf("%v\n", m1) //map[lixiang:100]
	fmt.Printf("%#v\n", m1) //map[string]int{"lixiang":100}

	//打印百分号
	printBaifenbi(90)

	fmt.Printf("%v\n", "100")
	//整数->字符
	fmt.Printf("%q\n", 65)
	//浮点数和复数
	fmt.Printf("%b\n", 3.1415926)
	//字符串
	fmt.Printf("%q\n", "李想")
	fmt.Printf("%7.3s\n", "李想有理想")

	//获取用户输入
	var s string
	fmt.Scan(&s)
	fmt.Println("用户输入的内容是：", s)

	var (
		name string
		age int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println(name, age, class)

	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)
}

func printBaifenbi(num int) {
	//用两个%，表示一个真正意义上的%
	fmt.Printf("%d%%\n", num)
}