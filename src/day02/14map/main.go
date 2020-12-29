package main

import "fmt"

//map

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        //还没有初始化（没有在内存中开辟空间）
	m1 = make(map[string]int, 10) //要估算好该map容量，避免在程序运行期间再动态扩容
	m1["caijiahui"] = 18
	m1["jiahui"] = 24
	fmt.Println(m1)
	fmt.Println(m1["jiahui"])
	val, ok := m1["cai"]
	//约定成俗用ok来接收返回的布尔值
	if !ok {
		fmt.Println("Not find this key")
	} else {
		fmt.Println(val)
	}

	//map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//删除
	delete(m1, "caijiahui")
	fmt.Println(m1)
	delete(m1, "cai") //删除不存在的key
}
