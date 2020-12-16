package main

import "fmt"

//1.求数组[1,3,5,7,8]所有元素的和
//2.找出数组中和为指定值的两个元素的下标，比如从数组[1,3,5,7,8]中找出何为8的两个元素的下标分别为(0,3)和(1,2)

func main() {
	//题1
	arr1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range arr1 {
		sum += v
	}
	fmt.Println("所有元素的和为：", sum)
	//题2
	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++{
			if arr1[i] + arr1[j] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}
