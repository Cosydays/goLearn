package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Atoi() 字符串=》整型
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%v\n", i1, i1)
	}

	//Itoa() 整型=》字符串
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2)

	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Printf("%T, %v\n", b, b)
	f, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Printf("%T, %v\n", f, f)
	i, err := strconv.ParseInt("-2", 10, 64)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Printf("%T, %v\n", i, i)
	u, err := strconv.ParseUint("2", 10, 64)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Printf("%T, %v\n", u, u)
}
