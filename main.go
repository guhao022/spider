package main

import (
	"github.com/num5/web"

	"spider/utils/log"
	"spider/core"

	"net/http"
	"io/ioutil"
	//"encoding/json"
)

func main() {

	r := web.New()

	r.Get("/", fc)

	log.CLog("[TRAC] Server start listen on # %d #", 9900)

	err := http.ListenAndServe(":9900", r)

	if err != nil {
		panic(err)
	}

}

func fc(w http.ResponseWriter, r *http.Request) {
	req := core.NewRequest()
	req.SetUrl("http://www.baidu.com").SetMethod("get").SetDepth(0)

	downclient := core.NewPageDownloader(nil)

	resp, err := downclient.Download(req)
	if err != nil {
		log.CLog("[ERRO] %s", err)
	}
	ra := resp.HttpResp()

	/*j, err := json.Marshal(ra.Body)
	if err != nil {
		log.CLog(err.Error())
	}*/
	s, err := ioutil.ReadAll(ra.Body)
	defer ra.Body.Close()

	if err != nil {
		log.CLog(err.Error())
	}

	w.Write([]byte(s))
}
