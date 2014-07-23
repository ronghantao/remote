package remote

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type ResponseStruct struct {
	EchoStatusCode    int
	EchoContentLength int64
	EchoBody          EchoBodyStruct
	EchoHeader        http.Header
}

type EchoBodyStruct struct {
	io.ReadCloser
}

func (eb *EchoBodyStruct) GetBytes() []byte {
	body, err := ioutil.ReadAll(eb)
	if err != nil {
		return nil
	}
	return body
}

func (eb *EchoBodyStruct) ToString() (string, error) {
	body, err := ioutil.ReadAll(eb)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (eb *EchoBodyStruct) ToJson(o interface{}) error {
	if body, err := ioutil.ReadAll(eb); err != nil {
		return err
	} else if err := json.Unmarshal(body, o); err != nil {
		return err
	}

	return nil
}
