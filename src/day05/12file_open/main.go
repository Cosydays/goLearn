package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//打开文件1.
func readFromFile1() {
	fileObj, err := os.Open("C:/Users/Cosydays/Desktop/goLearn/src/day05/12file_open/main.go")
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	//读文件
	//var tmp = make([]byte, 128) //指定读的长度
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

//打开文件2.
//利用bufio这个包读文件
func readFromFileByBufio() {
	fileObj, err := os.Open("C:/Users/Cosydays/Desktop/goLearn/src/day05/12file_open/main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	//创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err:%v", err)
			return
		}
		fmt.Print(line)
	}
}

//打开文件3.
func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("C:/Users/Cosydays/Desktop/goLearn/src/day05/12file_open/main.go")
	if err != nil {
		fmt.Printf("read line failed, err:%v", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	//readFromFile1()
	//readFromFileByBufio()
	readFromFileByIoutil()
}
