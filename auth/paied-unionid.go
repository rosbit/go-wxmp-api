// 用户支付完成后，获取该用户的 UnionId，无需用户授权
//  [NOTE]调用前需要用户完成支付，且在支付后的五分钟内有效
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html

package auth

import (
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"fmt"
)

// get paied union_id by open_id & transaction_id
func GetPaidUnionIdByTxId(cfgName string, openId string, txId string) (unionId string, err error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/wxa/getpaidunionid?access_token=%s&openid=%s&transaction_id=%s",
			accessToken, openId, txId,
		)
		return
	}

	return getPaiedUnionId(cfgName, genParams)
}

// get paied union_id by open_id, mch_id & out_trade_no
func GetPaidUnionIdByTradeNo(cfgName string, openId string, mchId, outTradeNo string) (unionId string, err error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/wxa/getpaidunionid?access_token=%s&openid=%s&mch_id=%si&out_trade_no=%s",
			accessToken, openId, mchId, outTradeNo,
		)
		return
	}

	return getPaiedUnionId(cfgName, genParams)
}

func getPaiedUnionId(cfgName string, genParams FnGeneParams) (string, error) {
	var res struct {
		callwxmp.BaseResult
		Unionid string
	}
	if err := CallWxmp(cfgName, genParams, "GET", callwxmp.HttpCall, &res); err != nil {
		return "", err
	}
	return res.Unionid, nil
}
