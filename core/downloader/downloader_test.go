package downloader

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
	"spider/utils/context"
	"spider/utils/page"
)

func TestDownloadHtml(t *testing.T) {
	var req *context.Request
	req = context.NewRequest("http://live.sina.com.cn/zt/l/v/finance/globalnews1/", "html", "", "GET", "", nil, nil, nil, nil)

	var dl Downloader
	dl = NewHttpDownloader()

	var p *page.Page
	p = dl.Download(req)

	var doc *goquery.Document
	doc = p.GetHtmlParser()
	//fmt.Println(doc)
	//body := p.GetBodyStr()
	//fmt.Println(body)

	var s *goquery.Selection
	s = doc.Find("body")
	if s.Length() < 1 {
		t.Error("html parse failed!")
	}

	   /*doc, err := goquery.NewDocument("http://live.sina.com.cn/zt/l/v/finance/globalnews1/")
	   if err != nil {
		   fmt.Printf("%v",err)
	   }
	   s := doc.Find("meta");
	   fmt.Println(s.Length())

	   resp, err := http.Get("http://live.sina.com.cn/zt/l/v/finance/globalnews1/")
	   if err != nil {
		   fmt.Printf("%v",err)
	   }
	   defer resp.Body.Close()
	   doc, err = goquery.NewDocumentFromReader(resp.Body)
	   s = doc.Find("meta");
	   fmt.Println(s.Length())*/
}
