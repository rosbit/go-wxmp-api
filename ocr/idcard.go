// 本接口提供基于小程序的身份证 OCR 识别
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.idcard.html

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

type IdCard interface {
	GetType() string  // Front: 正面, Back: 反面
}

type idcardType struct {
	Type string `json:"type"`
}

func (id *idcardType) GetType() string {
	return id.Type
}

type FrontIdCard struct {
	idcardType
	Name string `json:"name"`
	Id   string `json:"id"`
	Addr string `json:"addr"`
	Gender string `json:"gender"`
	Nationality string `json:"nationality"`
}

type BackIdCard struct {
	idcardType
	ValidDate string `json:"valid_date"`
}


func OcrIdCardImgUrl(cfgName, imgUrl string, isFront bool) (IdCard, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		v := u.Values{}
		v.Set("img_url", imgUrl)
		v.Set("access_token", accessToken)
		v.Set("type", "photo")
		url = fmt.Sprintf("https://api.weixin.qq.com/cv/ocr/idcard?%s", v.Encode())
		return
	}

	return ocrIdCard(cfgName, genParams, isFront)
}

func OcrIdCardImg(cfgName, fileName string, fp io.Reader, isFront bool) (IdCard, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/cv/ocr/idcard?type=photo&access_token=%s", accessToken)
		params := []multipart.Param{
			multipart.Param{"img", fileName, fp},
		}

		b := &bytes.Buffer{}
		contentType, _ := multipart.Create(b, "", params)
		body = b.Bytes()
		headers = map[string]string{"Content-Type": contentType}
		return
	}

	return ocrIdCard(cfgName, genParams, isFront)
}

func ocrIdCard(cfgName string, genParams auth.FnGeneParams, isFront bool) (IdCard, error) {
	if isFront {
		var res struct {
			callwxmp.BaseResult
			FrontIdCard
		}
		if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.HttpCall, &res); err != nil {
			return nil, err
		}
		return &res.FrontIdCard, nil
	} else {
		var res struct {
			callwxmp.BaseResult
			BackIdCard
		}
		if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.HttpCall, &res); err != nil {
			return nil, err
		}
		return &res.BackIdCard, nil
	}
}
