// 获取小程序全局唯一后台接口调用凭据（access_token）
// 参考文档： https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html

package auth

import (
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/conf"
	"time"
	"fmt"
	"os"
	"encoding/json"
)

type AccessToken struct {
	cfg *conf.WxmpConfT
	accessToken string
	expireTime int64
}

func NewAccessToken(name string) *AccessToken {
	cfg := conf.GetWxmpConf(name)
	if cfg == nil {
		return nil
	}
	token := &AccessToken{cfg: cfg}
	token.loadFromStore()
	return token
}

func (token *AccessToken) Get() (string, error) {
	if token.expired() {
		err := token.get_access_token()
		if err != nil {
			return "", err
		}
	}
	return token.accessToken, nil
}

func (token *AccessToken) expired() bool {
	return token.expireTime < time.Now().Unix()
}

func (token *AccessToken) get_access_token() error {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		token.cfg.AppId,
		token.cfg.AppSecret,
	)
	var res struct {
		callwxmp.BaseResult
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}
	if err := callwxmp.CallWxmp(url, "GET", nil, nil, callwxmp.HttpCall, &res); err != nil {
		return err
	}
	token.accessToken = res.AccessToken
	token.expireTime =  res.ExpiresIn + time.Now().Unix() - 10

	return token.saveToStore()
}

func (token *AccessToken) savePath() string {
	return fmt.Sprintf("%s/%s", conf.TokenStorePath, token.cfg.AppId)
}

func (token *AccessToken) saveToStore() error {
	if _, err := os.Stat(conf.TokenStorePath); os.IsNotExist(err) {
		if err = os.MkdirAll(conf.TokenStorePath, 0755); err != nil {
			return err
		}
	}
	savePath := token.savePath()
	fp, err := os.OpenFile(savePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()

	j, _ := json.Marshal(map[string]interface{} {
		"token": token.accessToken,
		"expire": token.expireTime,
	})
	fp.Write(j)
	return nil
}

func (token *AccessToken) loadFromStore() {
	savePath := token.savePath()
	fp, err := os.Open(savePath)
	if err != nil {
		return
	}
	defer fp.Close()

	var j struct {
		Token  string `json:"token"`
		Expire int64  `json:"expire"`
	}
	dec := json.NewDecoder(fp)
	if err = dec.Decode(&j); err != nil {
		return
	}

	token.accessToken, token.expireTime = j.Token, j.Expire
}
