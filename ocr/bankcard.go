// 本接口提供基于小程序的银行卡 OCR 识别
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.bankcard.html

package ocr

import (
	"github.com/rosbit/multipart-creator"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"bytes"
	"fmt"
	"io"
	u "net/url"
)

func OcrBankcardImgUrl(cfgName, imgUrl string) (id string, err error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		v := u.Values{}
		v.Set("img_url", imgUrl)
		v.Set("access_token", accessToken)
		url = fmt.Sprintf("https://api.weixin.qq.com/cv/ocr/bankcard?%s", v.Encode())
		return
	}

	return ocrBankcard(cfgName, genParams)
}

func OcrBankcardImg(cfgName, fileName string, fp io.Reader) (id string, err error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/cv/ocr/bankcard?access_token=%s", accessToken)
		params := []multipart.Param{
			multipart.Param{"img", fileName, fp},
		}

		b := &bytes.Buffer{}
		contentType, _ := multipart.Create(b, "", params)
		body = b.Bytes()
		headers = map[string]string{"Content-Type": contentType}
		return
	}

	return ocrBankcard(cfgName, genParams)
}

func ocrBankcard(cfgName string, genParams auth.FnGeneParams) (string, error) {
	var res struct {
		callwxmp.BaseResult
		Id string `json:"id"`
	}
	if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.HttpCall, &res); err != nil {
		return "", err
	}
	return res.Id, nil
}
