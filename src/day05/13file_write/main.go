package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件写内容

func writeDemo1() {
	fileObj, err := os.OpenFile("C:/Users/Cosydays/Desktop/goLearn/src/day05/13file_write/xx.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//write
	fileObj.Write([]byte("cai jiahui\n"))
	//writeString
	fileObj.WriteString("casparcai")
	defer fileObj.Close()
}

func writeDemo2() {
	fileObj, err := os.OpenFile("C:/Users/Cosydays/Desktop/goLearn/src/day05/13file_write/xx.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileObj.Close()
	//创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello world\n") //写到缓存中
	wr.Flush()                      //将缓存中的内容写入文件
}

func writeDemo3() {
	str := "hello world!"
	err := ioutil.WriteFile("C:/Users/Cosydays/Desktop/goLearn/src/day05/13file_write/xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
func main() {
	//writeDemo1
	//writeDemo2()
	writeDemo3()
}
