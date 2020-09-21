// 下发小程序和公众号统一的服务消息
// 参考：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html

package msg

import (
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"fmt"
)

// 发送小程序模板消息
func SendMiniProgramMsg(cfgName, toUser, page, tmplId, formId string, data map[string]interface{}, keyword string) error {
	params := map[string]interface{}{
		"touser": toUser,
		"weapp_template_msg": map[string]interface{}{
			"template_id": tmplId,
			"page": page,
			"form_id": formId,
			"data": func()map[string]interface{}{
				res := map[string]interface{}{}
				for k, v := range data {
					res[k] = map[string]string{
						"value": fmt.Sprintf("%v", v),
					}
				}
				return res
			}(),
			"emphasis_keyword": keyword,
		},
	}

	return uniformSend(cfgName, params)
}

// 发送公众号模板消息
type MsgData struct {
	Val interface{}
	Color string
}
func SendOfficialAccountMsg(cfgName, toUser, tmplId, appId, url, miniProgram string, data map[string]MsgData) error {
	params := map[string]interface{}{
		"touser": toUser,
		"mp_template_msg": map[string]interface{}{
			"appid": appId,
			"template_id": tmplId,
			"url": url,
			"miniprogram": miniProgram,
			"data": func()map[string]interface{}{
				res := map[string]interface{}{}
				for k, v := range data {
					res[k] = map[string]string{
						"value": fmt.Sprintf("%v", v.Val),
						"color": v.Color,
					}
				}
				return res
			}(),
		},
	}

	return uniformSend(cfgName, params)
}

func uniformSend(cfgName string, params map[string]interface{}) error {
	genParams := func(accessToken string)(url string, body interface{}, headers map[string]string)  {
		url = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=%s", accessToken)
		body = params
		return
	}

	var res callwxmp.BaseResult
	return auth.CallWxmp(cfgName, genParams, "POST", callwxmp.JsonCall, &res)
}
