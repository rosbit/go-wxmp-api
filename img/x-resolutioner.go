// 本接口提供基于小程序的图片高清化能力
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.superresolution.html

package img

import (
	"github.com/rosbit/multipart-creator"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"bytes"
	"fmt"
	"io"
	u "net/url"
)

func XResolutionImgUrl(cfgName, imgUrl string) (mediaId string, err error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		v := u.Values{}
		v.Set("img_url", imgUrl)
		v.Set("access_token", accessToken)
		url = fmt.Sprintf("https://api.weixin.qq.com/cv/img/superresolution?%s", v.Encode())
		return
	}

	return xResolution(cfgName, genParams)
}

func XResolutionImg(cfgName, fileName string, fp io.Reader) (mediaId string, err error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/cv/img/superresolution?access_token=%s", accessToken)
		params := []multipart.Param{
			multipart.Param{"img", fileName, fp},
		}

		b := &bytes.Buffer{}
		contentType, _ := multipart.Create(b, "", params)
		body = b.Bytes()
		headers = map[string]string{"Content-Type": contentType}
		return
	}

	return xResolution(cfgName, genParams)
}

func xResolution(cfgName string, genParams auth.FnGeneParams) (string, error) {
	var res struct {
		callwxmp.BaseResult
		MediaId string `json:"media_id"`
	}
	if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.HttpCall, &res); err != nil {
		return "", err
	}
	return res.MediaId, nil
}
