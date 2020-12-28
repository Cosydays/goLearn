package main

import "fmt"

//panic & recover
func f1() {
	defer func() {
		err := recover() //收集当前的错误
		fmt.Println("放手去爱")
		fmt.Println(err)
	}() //加括号代表立即执行
	panic("犯了不可原谅的错误")
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}


func main() {
	f1()
	f2()
}
