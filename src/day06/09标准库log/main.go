package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Open log file failed, err:", err)
		return
	}
	//配置日志输入位置
	log.SetOutput(logFile)
	//flag选项
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	//配置日志前缀
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志")
}
