package main

import (
	"spider/core/common/log"
	"spider/core"

	"github.com/num5/web"
	"net/http"
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
	i := resp.HttpResp()

	w.Write()
}
