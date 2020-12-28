package main

import "fmt"

//递归：自己调用自己

//计算n的阶乘
func f(n uint64) uint64 {
	if (n <= 1) {
		return 1
	}
	return n * f(n - 1)
}
func main() {
	ret := f(5)
	fmt.Println(ret)
}
