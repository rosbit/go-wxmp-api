// 校验一张图片是否含有违法违规内容。
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.imgSecCheck.html

package security

import (
	"github.com/rosbit/multipart-creator"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"github.com/rosbit/go-wxmp-api/img-util"
	"io"
)

func ImgSecCheck(cfgName, fileName string, fp io.Reader) error {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		return imgutil.GenerateImgMultipartParams(
			"https://api.weixin.qq.com/wxa/img_sec_check",
			accessToken,
			[]multipart.Param{
				multipart.Param{"media", fileName, fp},
			},
		)
	}

	var res callwxmp.BaseResult
	return auth.CallWxmp(cfgName, genParams, "POST", callwxmp.HttpCall, &res)
}
