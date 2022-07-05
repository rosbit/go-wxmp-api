// 本接口用于获取用于 NFC 的小程序 scheme 码，适用于 NFC 拉起小程序的业务场景。目前仅针对国内非个人主体的小程序开放
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generateNFC.html

package urlscheme

import (
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"fmt"
)

func GenerateSchemeNFC(cfgName string, jumpWxa JumpWxa, sn string, modelId string) (openlink string, err error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/wxa/generatenfcscheme?access_token=%s", accessToken)
		body = map[string]interface{}{
			"jump_wxa": jumpWxa,
			"sn": sn,
			"model_id": modelId,
		}
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
