package main

func main() {
	/*	resp, err := http.Get("https://www.liwenzhou.com/")
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("read from resp.Body failed, err%v\n", err)
			return
		}
		fmt.Print(string(body))*/

	//带参数的GET
	/*	apiUrl := "http://127.0.0.1:9090/get"
		//URL param
		data := url.Values{}
		data.Set("name", "小王子")
		data.Set("age", "18")
		u, err := url.ParseRequestURI(apiUrl)
		if err != nil {
			fmt.Printf("parse url requestUrl failed, err:%v\n", err)
			return
		}
		u.RawQuery = data.Encode()
		fmt.Println(u.String())
		resp, err := http.Get(u.String())
		if err != nil {
			fmt.Printf("post failed, err:%v\n", err)
			return
		}
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("ger resp failed, err:%v\n", err)
			return
		}
		fmt.Println(string(b))*/

	//Post请求
	url := "http://127.0.0.1:9090/post"

}
