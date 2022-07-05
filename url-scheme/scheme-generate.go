// 本接口用于获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景。目前仅针对国内非个人主体的小程序开放
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html

package urlscheme

import (
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"fmt"
)

func GenerateScheme(cfgName string, jumpWxa JumpWxa, expireType ExpireType, expireTime int64) (openlink string, err error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/wxa/generatescheme?access_token=%s", accessToken)
		b := map[string]interface{}{
			"jump_wxa": jumpWxa,
			"expire_type": expireType,
		}
		if expireType == ExpireTypeTime {
			b["expire_time"] = expireTime
		} else {
			b["expire_interval"] = expireTime
		}
		body = b
		return
	}

	var res struct {
		callwxmp.BaseResult
		Openlink string `json:"openlink"`
	}
	if err = auth.CallWxmp(cfgName, genParams, "POST", callwxmp.JsonCall, &res); err != nil {
		return
	}
	openlink = res.Openlink
	return
}
