package remote

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"
)

type HttpClient struct {
	MaxConnectCount int
	MaxRedirect     int
	Timeout         time.Duration
	Insecure        bool
}

type Error struct {
	timeout bool
	Err     error
}

func newResponse(res *http.Response) *ResponseStruct {
	return &ResponseStruct{EchoStatusCode: res.StatusCode, EchoContentLength: res.ContentLength, EchoHeader: res.Header, EchoBody: EchoBodyStruct{res.Body}}
}

func (e *Error) Timeout() bool {
	return e.timeout
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func GetDefaultClient() *HttpClient {
	return &HttpClient{MaxConnectCount: 3, MaxRedirect: 2, Timeout: time.Second * 30}
}

var defaultDialer = &net.Dialer{Timeout: 3000 * time.Millisecond}
var defaultTransport = &http.Transport{Dial: defaultDialer.Dial}
var defaultClient = &http.Client{Transport: defaultTransport}

func (hc *HttpClient) Do(req *http.Request) (*ResponseStruct, error) {
	if hc.Insecure {
		defaultTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	} else if defaultTransport.TLSClientConfig != nil {
		// the default TLS client (when transport.TLSClientConfig==nil) is
		// already set to verify, so do nothing in that case
		defaultTransport.TLSClientConfig.InsecureSkipVerify = false
	}
	timeout := false
	var timer *time.Timer
	if hc.Timeout > 0 {
		timer = time.AfterFunc(hc.Timeout, func() {
			defaultTransport.CancelRequest(req)
			timeout = true
		})
	}
	res, err := defaultClient.Do(req)
	if timer != nil {
		timer.Stop()
	}

	if err != nil {
		if op, ok := err.(*net.OpError); !timeout && ok {
			timeout = op.Timeout()
		}
		return nil, &Error{timeout: timeout, Err: err}
	}

	if isRedirect(res.StatusCode) && hc.MaxConnectCount > 0 {
		loc, _ := res.Location()
		hc.MaxConnectCount--
		u, err := url.Parse(loc.String())
		if err != nil {
			return nil, err
		}
		req.URL = u
		return hc.Do(req)
	}
	return newResponse(res), nil
}

func isRedirect(status int) bool {
	switch status {
	case http.StatusMovedPermanently:
		return true
	case http.StatusFound:
		return true
	case http.StatusSeeOther:
		return true
	case http.StatusTemporaryRedirect:
		return true
	default:
		return false
	}
}
