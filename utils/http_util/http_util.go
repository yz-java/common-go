package http_util

import (
	"net/http"
	"net"
	"time"
	"strings"
	"io/ioutil"
	"net/url"
)
//超时时间单位秒/s
func CreateHttpClient(timeout time.Duration) http.Client {
	//默认超时30s
	if timeout == 0 {
		timeout=30
	}
	c := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(timeout)
				c, err := net.DialTimeout(netw, addr, timeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
	return c
}

func responseToString(resp *http.Response) string  {
	respData,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(respData)
}

func Get(url string) string {
	resp,err:=http.Get(url)
	if err != nil {
		return ""
	}
	return responseToString(resp)
}

//post 字节流
func PostEntity(url, entity string) string  {
	resp,err:=http.Post(url,"",strings.NewReader(entity))
	if err != nil {
		return ""
	}
	respData,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(respData)
}
//post form 表单
func PostForm(URL string, dataMap map[string]string) string {
	values := url.Values{}
	for k,v :=range dataMap{
		values.Set(k,v)
	}
	resp,err:=http.PostForm(URL,values)
	if err != nil {
		return ""
	}
	return responseToString(resp)
}


