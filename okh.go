package main

import (
	"fmt"
	"ghttp/ghttp"
)

func goGet()  {
	http := ghttp.NewHttp()
	s := http.Headers(map[string]string{
		"auth": "1234",
	}).Get().Req("https://wwww.baidu.com", map[string]string{
		"keyword":"golang",
	})
	fmt.Println(s)
}

func doPost(){
	http := ghttp.NewHttp()
	s :=http.Headers(map[string]string{
		"auth":"ssss",
	}).Post().Json().Req("https://127.0.0.1:8080", map[string]string{
		"keyword":"golang",
	})
	fmt.Println(s)
}
func main() {
	goGet()
	//doPost()
}
