package basic

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestProxy(t *testing.T) {
	httpReq, err := http.NewRequest(http.MethodPost, "http://www.baidu.com", bytes.NewBufferString("abc"))
	if err != nil {
		fmt.Println(err)
		return
	}
	httpReq.Header.Add("content-type", "application/json")
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://{userName}:{password}@{proxyIp}:{proxyPort}")
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}
	response, err := client.Do(httpReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	httpResp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(httpResp))
}
