// 本接口用于获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景。目前仅针对国内非个人主体的小程序开放
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html

package urllink

import (
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"fmt"
)

type ExpireType = uint8
const (
	ExpireTypeTime ExpireType = iota
	ExpireTypeDayInterval
)

func GenerateURL(cfgName string, path, query string, expireType ExpireType, expireTime int64) (urlLink string, err error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/wxa/generate_urllink?access_token=%s", accessToken)
		b := map[string]interface{}{
			"path": path,
			"query": query,
			"is_expire": true,
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
		UrlLink string `json:"url_link"`
	}
	if err = auth.CallWxmp(cfgName, genParams, "POST", callwxmp.JsonCall, &res); err != nil {
		return
	}
	urlLink = res.UrlLink
	return
}
