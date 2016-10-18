package download

import (
	"net/http"
	"spider/utils/context"
)

type Downloader interface {
	Download(*context.Request) (*context.Response, error)
}

type PageDownloader struct {
	client *http.Client
}

func NewPageDownloader(client *http.Client) *PageDownloader {
	if client == nil {
		client = &http.Client{}
	}
	return &PageDownloader{client}
}

func (self *PageDownloader) Download(req *context.Request) (resp *context.Response, err error) {
	var httpResp *http.Response

	url := req.GetUrl()
	b := req.GetPostBody()

	method := req.GetMethod()
	switch method {
	case "POST", "post", "Post":
		httpResp, err = self.client.PostForm(url, b)
		if err != nil {
			return nil, err
		}
	case "GET", "get", "Get":
		fallthrough
	default:
		httpResp, err = self.client.Get(url)
		if err != nil {
			return nil, err
		}
	}

	return context.NewResponse(req, httpResp), nil
}


