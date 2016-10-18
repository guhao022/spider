package context

import "net/http"

type ResourceManage interface {
	GetOne()
	FreeOne()
	Has() uint
	Left() uint
}

type Response struct {
	//抓取请求相关信息
	req *Request

	// 响应
	resp *http.Response

	// 响应错误信息
	err error
}

// 创建新的响应
func NewResponse(req *Request, resp *http.Response) *Response {
	return &Response{req: req, resp: resp}
}

// 设置错误信息
func (resp *Response) SetError(err error) {
	resp.err = err
}

// 获取错误信息
func (resp *Response) GetError() error {
	return resp.err
}

// 获取http相应
func (resp *Response) HttpResp() *http.Response {
	return resp.resp
}

// 获取原始请求
func (resp *Response) HttpReq() *Request {
	return resp.req
}

