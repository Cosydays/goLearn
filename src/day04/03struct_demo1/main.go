package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p person
	p.name = "caijiahui"
	p.age = 18
	p.gender = "man"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p)
}
