package qrcode

import (
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"fmt"
	"io"
)

// 获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html
func CreateQRCode(cfgName, path string, width int) (string, io.ReadCloser, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=%s", accessToken)
		body = map[string]interface{}{
			"path": path,
			"width": width,
		}
		return
	}

	return auth.FetchBody(cfgName, genParams, "POST", callwxmp.JsonCallToReturnBody)
}

// 获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func GetWxaCode(cfgName, path string, width int, autoColor bool, lineColor map[string]int, isHyaline bool) (string, io.ReadCloser, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacode?access_token=%s", accessToken)
		body = map[string]interface{}{
			"path": path,
			"width": width,
			"auto_color": autoColor,
			"line_color": lineColor,
			"is_hyaline": isHyaline,
		}
		return
	}

	return auth.FetchBody(cfgName, genParams, "POST", callwxmp.JsonCallToReturnBody)
}

// 获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制
// 参考文档： https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func GetWxaCodeUnlimit(cfgName, scene, page string, width int, autoColor bool, lineColor map[string]int, isHyaline bool) (string, io.ReadCloser, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", accessToken)
		body = map[string]interface{}{
			"scene": scene,
			"page": page,
			"width": width,
			"auto_color": autoColor,
			"line_color": lineColor,
			"is_hyaline": isHyaline,
		}
		return
	}

	return auth.FetchBody(cfgName, genParams, "POST", callwxmp.JsonCallToReturnBody)
}
