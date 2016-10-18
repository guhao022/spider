package analyzer

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"

	"net/http"
	"bufio"
	"io"

	"spider/utils/context"
)

type Page struct {

	isfail   bool
	errormsg string

	req *context.Request

	item *Item

	body io.ReadCloser
	header  http.Header
	cookies []*http.Cookie

	docParser *goquery.Document

	jsonMap *simplejson.Json

	targetRequests []*context.Request
}

func NewPage(req *context.Request) *Page {
	return &Page{req: req}
}

func (this *Page) Errormsg() string {
	return this.errormsg
}

func (this *Page) SetRequest(r *context.Request) *Page {
	this.req = r
	return this
}

func (this *Page) GetRequest() *context.Request {
	return this.req
}

func (this *Page) SetHeader(header http.Header) {
	this.header = header
}

func (this *Page) GetHeader() http.Header {
	return this.header
}

func (this *Page) SetCookies(cookies []*http.Cookie) {
	this.cookies = cookies
}

func (this *Page) GetCookies() []*http.Cookie {
	return this.cookies
}

func (this *Page) AddTargetRequest(url string, respType string) *Page {
	this.targetRequests = append(this.targetRequests, context.NewRequest(url, respType, "", "GET", "", nil, nil, nil, nil))
	return this
}

func (this *Page) AddTargetRequests(urls []string, respType string) *Page {
	for _, url := range urls {
		this.AddTargetRequest(url, respType)
	}
	return this
}

func (this *Page) AddTargetRequestWithParams(req *context.Request) *Page {
	this.targetRequests = append(this.targetRequests, req)
	return this
}

func (this *Page) AddTargetRequestsWithParams(reqs []*context.Request) *Page {
	for _, req := range reqs {
		this.AddTargetRequestWithParams(req)
	}
	return this
}

func (this *Page) GetTargetRequests() []*context.Request {
	return this.targetRequests
}

func (this *Page) SetHtmlParser(doc *goquery.Document) *Page {
	this.docParser = doc
	return this
}

func (this *Page) GetHtmlParser() *goquery.Document {
	return this.docParser
}

func (this *Page) ResetHtmlParser() *goquery.Document {
	r := bufio.NewReader(this.body)
	var err error
	this.docParser, err = goquery.NewDocumentFromReader(r)
	if err != nil {
		panic(err.Error())
	}
	return this.docParser
}

func (this *Page) SetJson(js *simplejson.Json) *Page {
	this.jsonMap = js
	return this
}

func (this *Page) GetJson() *simplejson.Json {
	return this.jsonMap
}





