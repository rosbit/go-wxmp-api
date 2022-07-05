package callwxmp

import (
	"github.com/rosbit/gnet"
	"fmt"
	"io"
)

type FnCallWxmp func(url string, method string, params interface{}, headers map[string]string) ([]byte, error)

func HttpCall(url string, method string, postData interface{}, headers map[string]string) ([]byte, error) {
	return callGnet(url, method, postData, headers, gnet.Http)
}

func JsonCall(url string, method string, jsonData interface{}, headers map[string]string) ([]byte, error) {
	return callGnet(url, method, jsonData, headers, gnet.JSON)
}

func callGnet(url string, method string, postData interface{}, headers map[string]string, fnCall gnet.HttpFunc) ([]byte, error) {
	status, content, _, err := fnCall(url, gnet.M(method), gnet.Params(postData), gnet.Headers(headers))
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
	return callToReturnBody(url, method, params, headers, gnet.JSON)
}

func JsonCallToReturnBody(url string, method string, params interface{}, headers map[string]string) (contentType string, body io.ReadCloser, err error) {
	return callToReturnBody(url, method, params, headers, gnet.Http)
}

func callToReturnBody(url string, method string, params interface{}, headers map[string]string, call gnet.HttpFunc) (contentType string, body io.ReadCloser, err error) {
	status, _, resp, e := call(url, gnet.M(method), gnet.Params(params), gnet.Headers(headers), gnet.DontReadRespBody())
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
