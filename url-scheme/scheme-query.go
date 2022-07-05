// 本接口用于查询小程序 scheme 码
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.query.html

package urlscheme

import (
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"fmt"
)

func QueryScheme(cfgName string, scheme string) (schemeInfo *SchemeInfo, visitOpenid string, err error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/wxa/queryscheme?access_token=%s", accessToken)
		body = map[string]interface{}{
			"scheme": scheme,
		}
		return
	}

	var res struct {
		callwxmp.BaseResult
		SchemeInfo `json:"scheme_info"`
		VisitOpenid string `json:"visit_openid"`
	}
	if err = auth.CallWxmp(cfgName, genParams, "POST", callwxmp.JsonCall, &res); err != nil {
		return
	}
	schemeInfo = &res.SchemeInfo
	visitOpenid = res.VisitOpenid
	return
}
