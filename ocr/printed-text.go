// 本接口提供基于小程序的通用印刷体 OCR 识别
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.printedText.html

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

type PrintedText struct {
	Items []struct{
		Text string `json:"text"`
		Pos  PosT   `json:"pos"`
	} `json:"code_results"`
	ImgSize ImgSizeT `json:"img_size"`
}

func ScanPrintedTextImgUrl(cfgName, imgUrl string) (*PrintedText, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		v := u.Values{}
		v.Set("img_url", imgUrl)
		v.Set("access_token", accessToken)
		url = fmt.Sprintf("http://api.weixin.qq.com/cv/ocr/comm?%s", v.Encode())
		return
	}

	return scanPrintedText(cfgName, genParams)
}

func ScanPrintedTextImg(cfgName, fileName string, fp io.Reader) (*PrintedText, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("http://api.weixin.qq.com/cv/ocr/comm?access_token=%s", accessToken)
		params := []multipart.Param{
			multipart.Param{"img", fileName, fp},
		}

		b := &bytes.Buffer{}
		contentType, _ := multipart.Create(b, "", params)
		body = b.Bytes()
		headers = map[string]string{"Content-Type": contentType}
		return
	}

	return scanPrintedText(cfgName, genParams)
}

func scanPrintedText(cfgName string, genParams auth.FnGeneParams) (*PrintedText, error) {
	var res struct {
		callwxmp.BaseResult
		PrintedText
	}
	if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.HttpCall, &res); err != nil {
		return nil, err
	}
	return &res.PrintedText, nil
}
