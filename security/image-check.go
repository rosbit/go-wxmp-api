// 校验一张图片是否含有违法违规内容。
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.imgSecCheck.html

package security

import (
	"github.com/rosbit/multipart-creator"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"bytes"
	"fmt"
	"io"
)

func ImgSecCheck(cfgName, fileName string, fp io.Reader) error {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/wxa/img_sec_check?access_token=%s", accessToken)
		params := []multipart.Param{
			multipart.Param{"media", fileName, fp},
		}

		b := &bytes.Buffer{}
		contentType, _ := multipart.Create(b, "", params)
		body = b.Bytes()
		headers = map[string]string{"Content-Type": contentType}
		return
	}

	var res callwxmp.BaseResult
	return auth.CallWxmp(cfgName, genParams, "POST", callwxmp.HttpCall, &res)
}
