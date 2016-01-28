package core

import "net/http"

type Response struct {
	//抓取请求相关信息
	req *Request

	resp *http.Response

	err error
}

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

