// 本接口提供基于小程序的图片高清化能力
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.superresolution.html

package img

import (
	"github.com/rosbit/multipart-creator"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/img-util"
	"github.com/rosbit/go-wxmp-api/auth"
	"io"
)

func XResolutionImgUrl(cfgName, imgUrl string) (mediaId string, err error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = imgutil.GenerateUrl("https://api.weixin.qq.com/cv/img/superresolution", accessToken, imgUrl)
		return
	}

	return xResolution(cfgName, genParams)
}

func XResolutionImg(cfgName, fileName string, fp io.Reader) (mediaId string, err error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		return imgutil.GenerateImgMultipartParams(
			"https://api.weixin.qq.com/cv/img/superresolution",
			accessToken,
			[]multipart.Param{
				multipart.Param{"img", fileName, fp},
			},
		)
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
