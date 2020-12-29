package main

import "fmt"

//if条件判断
func main() {
	age := 19

	//if age > 18 {
	//	fmt.Println("大于18岁！")
	//} else {
	//	fmt.Println("小于18岁！")
	//}

	//if age > 35 {
	//	fmt.Println("人到中年！")
	//} else if age > 18 {
	//	fmt.Println("人到青年！")
	//} else {
	//	fmt.Println("小孩子！")
	//}

	if age := 19; age > 18 {
		fmt.Println("大于18岁")
	} else {
		fmt.Println("小于18岁！")
	}
	fmt.Println(age)

}
