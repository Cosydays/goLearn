package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//1.判断字符串中汉字的数量
	//难点是判断一个字符是汉字
	s1 := "Hello蔡家辉"
	var count int
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)

	//2.how do you do 单词出现的次数
	s2 := "how do you do"
	s3 := strings.Split(s2, " ")
	m1 := make(map[string]int, 10)
	for _, w := range(s3) {
		fmt.Println(w)
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	for key, value := range(m1) {
		fmt.Println(key, value)
	}

	//3.回文判断
	//字符串从左往右读和从右往左读是一样的，那就是回文
	s4 := "aba"
	for i := 0; i < len(s4) / 2; i++{
		if s4[i] != s4[len(s4) - i -1] {
			fmt.Println("不是回文！")
		}
	}
	fmt.Println("是回文！")
}
