// 登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html

package auth

import (
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/conf"
	"fmt"
)

type Session struct {
	Openid string
	SessionKey string `json:"session_key"`
	Unionid string
}

func Code2Session(cfgName string, code string) (*Session, error) {
	cfg := conf.GetWxmpConf(cfgName)
	if cfg == nil {
		return nil, fmt.Errorf("no config found for %s", cfgName)
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		cfg.AppId, cfg.AppSecret, code,
	)
	var res struct {
		callwxmp.BaseResult
		Session
	}
	if err := callwxmp.CallWxmp(url, "GET", nil, nil, callwxmp.HttpCall, &res); err != nil {
		return nil, err
	}
	return &res.Session, nil
}

