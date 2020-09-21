// 检查一段文本是否含有违法违规内容。
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.msgSecCheck.html

package security

import (
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"fmt"
)

func MsgSecCheck(cfgName, content string) error {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s", accessToken)
		body = map[string]interface{}{
			"content": content,
		}
		return
	}

	var res callwxmp.BaseResult
	return auth.CallWxmp(cfgName, genParams, "POST", callwxmp.JsonCall, &res)
}
