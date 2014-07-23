package remote

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

type RequestStruct struct {
	Headers        map[string]string
	QueryStringMap map[string]string
	Uri            string
	ContentType    string
	Accept         string
	UserAgent      string
}

const (
	HttpMethod_post = "POST"
	HttpMethod_put  = "PUT"
	HttpMethod_get  = "GET"
	HttpMethod_del  = "DELETE"
)

func (r *RequestStruct) initQueryString() {
	values := url.Values{}
	for k, v := range r.QueryStringMap {
		values.Add(k, v)
	}
	r.Uri = r.Uri + "?" + values.Encode()
}

func (r *RequestStruct) initHeader(req *http.Request) {
	req.Header.Set("User-Agent", r.UserAgent)
	req.Header.Set("Content-Type", r.ContentType)
	req.Header.Set("Accept", r.Accept)
	if r.Headers != nil {
		for k, v := range r.Headers {
			req.Header.Set(k, v)
		}
	}
}

func (r *RequestStruct) GetMethodRequest() (*http.Request, error) {
	if r.QueryStringMap != nil {
		r.initQueryString()
	}
	req, er := http.NewRequest(HttpMethod_get, r.Uri, nil)
	if er != nil {
		return nil, er
	}
	r.initHeader(req)
	return req, nil
}

func (r *RequestStruct) PostFileRequest(params map[string]string, fpath string, fpathName string) (*http.Request, error) {
	if r.QueryStringMap != nil {
		r.initQueryString()
	}
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	fileWriter, err := bodyWriter.CreateFormFile(fpathName, fpath)
	if err != nil {
		return nil, err
	}
	fh, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	if _, err = io.Copy(fileWriter, fh); err != nil {
		return nil, err
	}
	if params != nil {
		for k, v := range params {
			_ = bodyWriter.WriteField(k, v)
		}
	}
	r.ContentType = bodyWriter.FormDataContentType()
	bodyWriter.Close()
	req, er := http.NewRequest(HttpMethod_post, r.Uri, bodyBuffer)
	if er != nil {
		return nil, er
	}
	r.initHeader(req)
	return req, nil
}

func (r *RequestStruct) PostRequest(params map[string]string) (*http.Request, error) {
	if r.QueryStringMap != nil {
		r.initQueryString()
	}
	data := url.Values{}
	//处理参数列表
	if params != nil {
		for k, v := range params {
			data.Add(k, v)
		}
	}
	req, er := http.NewRequest(HttpMethod_post, r.Uri, bytes.NewBufferString(data.Encode()))
	if er != nil {
		return nil, er
	}
	r.initHeader(req)
	if len(r.ContentType) == 0 {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	return req, nil
}

func (r *RequestStruct) PutMethodRequest(params map[string]string) (*http.Request, error) {
	if r.QueryStringMap != nil {
		r.initQueryString()
	}
	data := url.Values{}
	//处理参数列表
	if params != nil {
		for k, v := range params {
			data.Add(k, v)
		}
	}

	req, er := http.NewRequest(HttpMethod_put, r.Uri, bytes.NewBufferString(data.Encode()))
	if er != nil {
		return nil, er
	}
	r.initHeader(req)
	if len(r.ContentType) == 0 {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	return req, nil
}
