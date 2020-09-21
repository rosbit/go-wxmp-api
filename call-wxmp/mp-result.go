package callwxmp

import (
	"strings"
	"fmt"
	"io"
	"encoding/json"
)

type WxmpResult interface {
	GetErrcode() int
	GetErrmsg() string
}

type BaseResult struct {
	Errcode int
	Errmsg  string
}
func (b *BaseResult) GetErrcode() int {
	return b.Errcode
}
func (b *BaseResult) GetErrmsg() string {
	return b.Errmsg
}

func CallWxmp(url string, method string, params interface{}, headers map[string]string, call FnCallWxmp, res WxmpResult) error {
	resp, err := call(url, method, params, headers)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(resp, res); err != nil {
		return err
	}
	if res.GetErrcode() != 0 {
		return fmt.Errorf("errcode: %d, errmsg: %s", res.GetErrcode(), res.GetErrmsg())
	}
	return nil
}

func CallwxmpToReturnBody(url string, method string, params interface{}, headers map[string]string, getBody FnGetBody) (string, io.ReadCloser, error) {
	contentType, body, err := getBody(url, method, params, headers)
	if err != nil {
		return "", nil, err
	}

	if strings.HasPrefix(contentType, "application/json") {
		defer body.Close()

		var res BaseResult
		if err = json.NewDecoder(body).Decode(&res); err != nil {
			return "", nil, err
		}
		return "", nil, fmt.Errorf("errcode: %d, errmsg: %s", res.GetErrcode(), res.GetErrmsg())
	}
	return contentType, body, nil
}
