package main

import (
	"github.com/num5/web"

	"spider/core/common/log"
	"spider/core"

	"net/http"
	"io/ioutil"
	"bufio"
)

func main() {

	r := web.New()

	r.Get("/", fc)


	log.CLog("[TRAC] Server start listen on # %d # ", 9900)

	err := http.ListenAndServe(":9900", r)

	if err != nil {
		panic(err)
	}

}

func fc(w http.ResponseWriter, r *http.Request) {
	req := core.NewRequest()
	req.SetUrl("http://www.golune.com").SetMethod("get").SetDepth(0)

	downclient := core.NewPageDownloader(nil)

	resp, err := downclient.Download(req)
	if err != nil {
		log.CLog("[ERRO] %s", err)
	}
	ra := resp.HttpResp()
	s, err := ioutil.ReadAll(bufio.NewReader(ra.Body))

	w.Write([]byte(s))
}
