package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// ResponseResult 响应结果
type ResponseResult struct {
	StatusCode int
	Header     http.Header
	Cookies    []*http.Cookie
	Body       []byte
}

// HTTPPostJSON 模拟请求POST方法 使用json
func HTTPPostJSON(reqURL string, headers map[string]string, paramMap map[string]interface{}, cookies []*http.Cookie, proxyIP string) (ResponseResult, error) {
	client := &http.Client{}
	var bytesData []byte
	if paramMap != nil {
		var err error
		bytesData, err = json.Marshal(paramMap)
		if err != nil {
			return ResponseResult{}, err
		}
	}
	params := bytes.NewReader(bytesData)

	//是否使用代理IP
	if proxyIP != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyIP)
		}
		transport := &http.Transport{Proxy: proxy}
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}
	req, err := http.NewRequest("POST", reqURL, params)
	if err != nil {
		return ResponseResult{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	if cookies != nil {
		for _, c := range cookies {
			req.AddCookie(c)
		}
	}
	rr := ResponseResult{}
	resp, err := client.Do(req)
	if err != nil {
		return rr, err
	}
	rr.StatusCode = resp.StatusCode
	rr.Header = resp.Header
	rr.Cookies = resp.Cookies()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rr, err
	}
	rr.Body = body
	return rr, nil
}

// HTTPPostForm 模拟请求POST方法 使用form
func HTTPPostForm(reqURL string, headers map[string]string, paramMap map[string]interface{}, cookies []*http.Cookie, proxyIP string) (ResponseResult, error) {

	client := &http.Client{}
	urlmap := url.Values{}
	if paramMap != nil {
		for k, v := range paramMap {
			urlmap.Add(k, fmt.Sprint(v))
		}
	}
	parmas := strings.NewReader(urlmap.Encode())

	//是否使用代理IP
	if proxyIP != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyIP)
		}
		transport := &http.Transport{Proxy: proxy}
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}
	req, err := http.NewRequest("POST", reqURL, parmas)
	if err != nil {
		return ResponseResult{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	if cookies != nil {
		for _, c := range cookies {
			req.AddCookie(c)
		}
	}
	rr := ResponseResult{}
	resp, err := client.Do(req)
	if err != nil {
		return rr, err
	}
	rr.StatusCode = resp.StatusCode
	rr.Header = resp.Header
	rr.Cookies = resp.Cookies()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rr, err
	}
	rr.Body = body
	return rr, nil
}

// HTTPPutForm 模拟请求PUT方法 使用form
func HTTPPutForm(reqURL string, headers map[string]string, paramMap map[string]interface{}, cookies []*http.Cookie, proxyIP string) (ResponseResult, error) {
	client := &http.Client{}
	urlmap := url.Values{}
	if paramMap != nil {
		for k, v := range paramMap {
			urlmap.Add(k, fmt.Sprint(v))
		}
	}
	params := strings.NewReader(urlmap.Encode())
	//是否使用代理IP
	if proxyIP != "" {
		//
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyIP)
		}
		transport := &http.Transport{Proxy: proxy}
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}
	req, err := http.NewRequest("PUT", reqURL, params)
	if err != nil {
		return ResponseResult{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	if cookies != nil {
		for _, c := range cookies {
			req.AddCookie(c)
		}
	}
	rr := ResponseResult{}
	resp, err := client.Do(req)
	if err != nil {
		return rr, err
	}
	rr.StatusCode = resp.StatusCode
	rr.Header = resp.Header
	rr.Cookies = resp.Cookies()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rr, err
	}
	rr.Body = body
	return rr, nil
}

// HTTPGet 模拟请求GET方法
func HTTPGet(reqURL string, headers map[string]string, paramMap map[string]interface{}, cookies []*http.Cookie, proxyIP string) (ResponseResult, error) {
	client := &http.Client{}
	urlmap := url.Values{}
	if paramMap != nil {
		for k, v := range paramMap {
			urlmap.Add(k, fmt.Sprint(v))
		}
	}
	parmas := strings.NewReader(urlmap.Encode())

	//是否使用代理IP
	if proxyIP != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyIP)
		}
		transport := &http.Transport{Proxy: proxy}
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}
	req, err := http.NewRequest("GET", reqURL, parmas)
	if err != nil {
		return ResponseResult{}, err
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	if cookies != nil {
		for _, c := range cookies {
			req.AddCookie(c)
		}
	}
	rr := ResponseResult{}
	resp, err := client.Do(req)
	if err != nil {
		return rr, err
	}
	rr.StatusCode = resp.StatusCode
	rr.Header = resp.Header
	rr.Cookies = resp.Cookies()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rr, err
	}
	rr.Body = body
	return rr, nil
}

// HTTPPostJSON 模拟请求POST方法 使用json
func HttpPostClass(reqURL string, headers map[string]string, class interface{}, cookies []*http.Cookie, proxyIP string) (ResponseResult, error) {

	client := &http.Client{}
	var bytesData []byte
	if class != nil {
		var err error
		bytesData, err = json.Marshal(class)
		if err != nil {
			return ResponseResult{}, err
		}
	}
	params := bytes.NewReader(bytesData)

	//是否使用代理IP
	if proxyIP != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyIP)
		}
		transport := &http.Transport{Proxy: proxy}
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}
	req, err := http.NewRequest("POST", reqURL, params)
	if err != nil {
		return ResponseResult{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	if cookies != nil {
		for _, c := range cookies {
			req.AddCookie(c)
		}
	}
	rr := ResponseResult{}
	resp, err := client.Do(req)
	if err != nil {
		return rr, err
	}
	rr.StatusCode = resp.StatusCode
	rr.Header = resp.Header
	rr.Cookies = resp.Cookies()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rr, err
	}
	rr.Body = body
	return rr, nil
}
