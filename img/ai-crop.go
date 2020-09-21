// 本接口提供基于小程序的图片智能裁剪能力
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.aiCrop.html

package img

import (
	"github.com/rosbit/multipart-creator"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"bytes"
	"fmt"
	"io"
	u "net/url"
	"strings"
)

type AICropRes struct {
	CropResults []struct{
		Left int `json:"crop_left"`
		Top  int `json:"crop_top"`
		Right  int `json:"crop_right"`
		Bottom int `json:"crop_bottom"`
	} `json:"results"`
	ImgSize struct {
		Width  int `json:"w"`
		Heigth int `json:"h"`
	} `json:"img_size"`
}

func AICropImgUrl(cfgName, imgUrl string, ratios []float64) (*AICropRes, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		v := u.Values{}
		v.Set("img_url", imgUrl)
		v.Set("access_token", accessToken)
		url = fmt.Sprintf("https://api.weixin.qq.com/cv/img/aicrop?%s", v.Encode())
		if len(ratios) == 0 {
			body = nil
		} else {
			r := make([]string, len(ratios))
			for i, rr := range ratios {
				r[i] = fmt.Sprintf("%v", rr)
			}
			body = fmt.Sprintf("ratios=%s", strings.Join(r, ","))
		}
		return
	}

	return aiCrop(cfgName, genParams)
}

func AICropImg(cfgName, fileName string, fp io.Reader, ratios []float64) (*AICropRes, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/cv/img/aicrop?access_token=%s", accessToken)
		var params []multipart.Param
		if len(ratios) == 0 {
			params = make([]multipart.Param, 1)
		} else {
			params = make([]multipart.Param, 2)
			r := make([]string, len(ratios))
			for i, rr := range ratios {
				r[i] = fmt.Sprintf("%v", rr)
			}
			params[1] = multipart.Param{"ratios", strings.Join(r, ","), nil}
		}
		params[0] = multipart.Param{"img", fileName, fp}

		b := &bytes.Buffer{}
		contentType, _ := multipart.Create(b, "", params)
		body = b.Bytes()
		headers = map[string]string{"Content-Type": contentType}
		return
	}

	return aiCrop(cfgName, genParams)
}

func aiCrop(cfgName string, genParams auth.FnGeneParams) (*AICropRes, error) {
	var res struct {
		callwxmp.BaseResult
		AICropRes
	}
	if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.HttpCall, &res); err != nil {
		return nil, err
	}
	return &res.AICropRes, nil
}
