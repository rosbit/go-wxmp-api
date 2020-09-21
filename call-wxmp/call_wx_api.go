package callwxmp

import (
	"github.com/rosbit/go-wget"
	"fmt"
	"io"
)

type FnCallWxmp func(url string, method string, params interface{}, headers map[string]string) ([]byte, error)

func HttpCall(url string, method string, postData interface{}, headers map[string]string) ([]byte, error) {
	return callWget(url, method, postData, headers, wget.Wget)
}

func JsonCall(url string, method string, jsonData interface{}, headers map[string]string) ([]byte, error) {
	return callWget(url, method, jsonData, headers, wget.PostJson)
}

func callWget(url string, method string, postData interface{}, headers map[string]string, fnCall wget.HttpFunc) ([]byte, error) {
	status, content, _, err := fnCall(url, method, postData, headers)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, fmt.Errorf("status %d", status)
	}
	return content, nil
}

type FnGetBody func(url string, method string, params interface{}, headers map[string]string) (contentType string, body io.ReadCloser, err error)

func PostToReturnBody(url string, method string, params interface{}, headers map[string]string) (contentType string, body io.ReadCloser, err error) {
	return callToReturnBody(url, method, params, headers, wget.PostJson)
}

func JsonCallToReturnBody(url string, method string, params interface{}, headers map[string]string) (contentType string, body io.ReadCloser, err error) {
	return callToReturnBody(url, method, params, headers, wget.Wget)
}

func callToReturnBody(url string, method string, params interface{}, headers map[string]string, call wget.HttpFunc) (contentType string, body io.ReadCloser, err error) {
	status, _, resp, e := call(url, method, params, headers, wget.Options{DontReadRespBody:true})
	if e != nil {
		err = e
		return
	}
	if resp == nil || resp.Body == nil {
		err = fmt.Errorf("no resp body")
		return
	}

	if status != 200 {
		err = fmt.Errorf("status code %d", status)
		return
	}
	body = resp.Body
	contentType = resp.Header.Get("Content-Type")
	return
}
