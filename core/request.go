package core

import (
	"net/http"
	"strings"
	"net/url"
)

// 请求
type Request struct {
	url string
	method string
	postBody url.Values
	header http.Header
	depth int
}

func NewRequest(url string) *Request {
	return &Request{url: url}
}

// 设置url
func (req *Request) SetUrl(url string) *Request {
	req.url = url
	return req
}

// 设置访问方式
func (req *Request) SetMethod(method string) *Request {
	req.method = strings.ToUpper(method)
	return req
}

// 设置参数
func (req *Request) SetPostBody(b url.Values) *Request {
	req.postBody = b
	return req
}

// 设置header
func (req *Request) SetHeader(key, val string) *Request {
	req.header.Set(key, val)
	return req
}

// 添加header
func (req *Request) AddHeader(key, val string) *Request {
	req.header.Add(key, val)
	return req
}

// 设置cookie
func (req *Request) SetCookies(cookie string) *Request {
	req.header.Set("Cookie", cookie)
	return req
}

// 设置深度
func (req *Request) SetDepth(depth int) *Request {
	req.depth = depth
	return req
}

// 获取url
func (req *Request) GetUrl() string {
	return req.url
}

// 获取method
func (req *Request) GetMethod() string {
	return req.method
}

// 获取参数信息
func (req *Request) GetPostBody() url.Values {
	return req.postBody
}

// 获取header信息
func (req *Request) GetHeader() http.Header {
	return req.header
}

// 获取cookie
func (req *Request) GetCookie() string {
	return req.header.Get("Cookie")
}

// 获取深度值。
func (req *Request) GetDepth() int {
	return req.depth
}


