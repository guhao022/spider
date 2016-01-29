package core

import (
	"testing"
	"spider/core/common/log"
	"bufio"
	"io/ioutil"
)

func TestDownload(t *testing.T) {
	req := NewRequest()
	req.SetUrl("http://www.golune.com").SetMethod("get").SetDepth(0)

	downclient := NewPageDownloader(nil)

	resp, err := downclient.Download(req)
	if err != nil {
		log.CLog("[ERRO] %s", err)
	}

	r := resp.HttpResp()
	s, err := ioutil.ReadAll(bufio.NewReader(r.Body))

	log.CLog("[SUCC] %s", string(s))
}
