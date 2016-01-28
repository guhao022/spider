package core

import (
	"net/http"
	"strings"
	"net/url"
	
	"spider/core/common/log"
)

// 请求
type Request struct {
	req *http.Request
	depth int
}

func NewRequest() *Request {
	return new(Request)
}

// 设置url
func (req *Request) SetUrl(u string) *Request {
	url, err := url.Parse(u)
	if err != nil {
		log.Error()
	}
	req.req.URL = url
	return req
}

// 设置访问方式
func (req *Request) SetMethod(method string) *Request {
	req.method = strings.ToUpper(method)
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

// 获取header信息
func (req *Request) GetHeader() http.Header {
	return req.header
}

// 获取cookie
func (req *Request) GetCookie() string {
	return req.header.Get("Cookie")
}

// 获取深度值。
func (req *Request) GetDepth() uint32 {
	return req.depth
}


