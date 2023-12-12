package connect

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: time.Second * 2,
}

func Get(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		logx.Errorf("connect.Get error: %v", err)
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
