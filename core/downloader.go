package core

import "net/http"

type Downloader interface {
	Download(*Request) *Response
}

type downloader struct {
	http.Client
}

func (self *downloader) Download(req *Request) *Response {
}