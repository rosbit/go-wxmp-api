// 本接口提供基于小程序的身份证 OCR 识别
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.idcard.html

package ocr

import (
	"github.com/rosbit/multipart-creator"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/img-util"
	"github.com/rosbit/go-wxmp-api/auth"
	"io"
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
		url = imgutil.GenerateUrl("https://api.weixin.qq.com/cv/ocr/idcard?type=photo", accessToken, imgUrl)
		return
	}

	return ocrIdCard(cfgName, genParams, isFront)
}

func OcrIdCardImg(cfgName, fileName string, fp io.Reader, isFront bool) (IdCard, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		return imgutil.GenerateImgMultipartParams(
			"https://api.weixin.qq.com/cv/ocr/idcard?type=photo",
			accessToken,
			[]multipart.Param{
				multipart.Param{"img", fileName, fp},
			},
		)
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
