/**
 * 微信小程序配置信息
 */
package conf

import (
	"fmt"
	"os"
	"time"
)

type WxmpConfT struct {
	AppId     string
	AppSecret string
}

var (
	WxmpConfs = map[string]*WxmpConfT{} // <name> => <configure>
	TokenStorePath string
	Loc = time.FixedZone("UTC+8", 8*60*60)
)

func SetWxmpConf(cfgName, appId, appSecret string) {
	WxmpConfs[cfgName] = &WxmpConfT{
		AppId: appId,
		AppSecret: appSecret,
	}
}

func GetWxmpConf(cfgName string) *WxmpConfT {
	if c, ok := WxmpConfs[cfgName]; ok {
		return c
	}
	return nil
}

func SetTokenStorePath(storePath string) {
	TokenStorePath = storePath
	setTZ()
}

func getEnv(name string, result *string, must bool) error {
	s := os.Getenv(name)
	if s == "" {
		if must {
			return fmt.Errorf("env \"%s\" not set", name)
		}
	}
	*result = s
	return nil
}

func setTZ() {
	var p string
	getEnv("TZ", &p, false)
	if p != "" {
		if loc, err := time.LoadLocation(p); err == nil {
			Loc = loc
		}
	}
}

