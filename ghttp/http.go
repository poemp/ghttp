package ghttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DefaultHttp struct {
	//发送的请求
	url string
	//请求头
	header map[string]string
	//需要发送的数据
	data map[string]string
}

const (
	// POST 请求
	POST = "POST"
	// GET 请求
	GET = "GET"
)

// NewHttp create ghttp clinet
func NewHttp() *DefaultHttp {
	return &DefaultHttp{
		url:    "",
		header: map[string]string{},
		data:   nil,
	}
}



// Headers 设置请求头
func (t *DefaultHttp) Headers(header map[string]string) *DefaultHttp {
	for s := range header {
		t.header[s] = header[s]
	}
	return t
}

// Get 转化为Post请求
func (t *DefaultHttp) Get() *Get {
	return &Get{
		DefaultHttp: t,
		method: GET,
	}
}

// Post 转化为Post请求
func (t *DefaultHttp) Post() *Post {
	return &Post{
		DefaultHttp: t,
		method: POST,
	}
}



type Get struct {
	*DefaultHttp
	//使用的方法 get 或者是 post
	method string
}

// Req 请求
func (t *Get) Req(url string, data map[string]string) string {
	t.url = url
	t.method = GET
	t.data = data
	return t.doGet()
}

// 发送get 请求
func (t *Get) doGet() string {
	params := url.Values{}
	if len(t.data) > 0 {
		data := t.data
		for s := range data {
			params.Set(s, data[s])
		}
	}
	client := &http.Client{}
	Url, _ := url.Parse(t.url)
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	req, e := http.NewRequest(t.method, urlPath, bytes.NewReader(nil))
	if e != nil {
		panic(e.Error())
	}
	//如果header存在
	if len(t.header) > 0 {
		header := t.header
		for k := range header {
			req.Header.Add(k, header[k])
		}
	}
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

// post 請求

type Post struct {
	*DefaultHttp
	//使用的方法 get 或者是 post
	method string
}

// Req post请求
func (t *Post) Req(url string, data map[string]string) string {
	t.url = url
	t.method = POST
	t.data = data
	return t.doPost()
}

// Json 发送json数据
//
// Post 请求才有效果
//
//Get 请求忽略
func (t *Post) Json() *Post {
	t.header["Content-Type"] = "application/json; charset=utf-8"
	return t
}

// Form form表单
//
// Post 请求才有效果
//
//Get 请求忽略
func (t *Post) Form() *Post {
	t.header["Content-Type"] = "application/x-www-form-urlencoded"
	return t
}

//对 post 请求进行封装
func (t *Post) doPost() string {
	client := &http.Client{}

	byteData, _ := json.Marshal(t.data)

	req, e := http.NewRequest(t.method, t.url, bytes.NewReader(byteData))
	if e != nil {
		panic(e.Error())
	}
	//如果header存在
	if len(t.header) > 0 {
		header := t.header
		for k := range header {
			req.Header.Add(k, header[k])
		}
	}
	res, ee := client.Do(req)
	if ee != nil {
		panic(ee.Error())
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
