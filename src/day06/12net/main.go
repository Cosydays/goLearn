package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("err!")
		return
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}
