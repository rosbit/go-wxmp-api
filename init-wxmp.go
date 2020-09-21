package wxmpapi

import (
	"github.com/rosbit/go-wxmp-api/conf"
)

func SetWxmpConf(cfgName, appId, appKey string) {
	conf.SetWxmpConf(cfgName, appId, appKey)
}

func SetTokenStorePath(tokenStorePath string) {
	conf.SetTokenStorePath(tokenStorePath)
}
