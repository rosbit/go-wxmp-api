// 数据签名校验
// 接口会同时返回 rawData、signature，其中 signature = sha1( rawData + session_key )
// 参考文档：https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html

package decrypt

import (
	"crypto/sha1"
	"io"
	"fmt"
)

func CheckSignature(sessionKey string, rawData []byte, signature string) error {
	h := sha1.New()
	h.Write(rawData)
	io.WriteString(h, sessionKey)
	signature2 := fmt.Sprintf("%x", h.Sum(nil))
	if signature == signature2 {
		return nil
	}
	return fmt.Errorf("not matched, signature(%s) != signature2(%s)", signature, signature2)
}
