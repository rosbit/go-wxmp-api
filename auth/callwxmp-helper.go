// 所有需要access_token调用的统一入口

package auth

import (
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"io"
	"fmt"
)

type FnGeneParams func(accessToken string)(url string, body interface{}, headers map[string]string)

func FetchBody(cfgName string, genParams FnGeneParams, method string, getBody callwxmp.FnGetBody) (string, io.ReadCloser, error) {
	token := NewAccessToken(cfgName)
	if token == nil {
		return "", nil, fmt.Errorf("no config found for %s", cfgName)
	}
	accessToken, err := token.Get()
	if err != nil {
		return "", nil, err
	}
	url, body, headers := genParams(accessToken)
	return callwxmp.CallwxmpToReturnBody(url, method, body, headers, getBody)
}

func CallWxmp(cfgName string, genParams FnGeneParams, method string, call callwxmp.FnCallWxmp, res callwxmp.WxmpResult) error {
	token := NewAccessToken(cfgName)
	if token == nil {
		return fmt.Errorf("no config found for %s", cfgName)
	}
	accessToken, err := token.Get()
	if err != nil {
		return err
	}
	url, body, headers := genParams(accessToken)
	return callwxmp.CallWxmp(url, method, body, headers, call, res)
}
