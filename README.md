# 微信小程序开发SDK

go-wxmp-api是对微信小程序API的封装，可以当作SDK使用，特点
 - 支持多个小程序

## 使用方法

```go
packge main

import (
	"github.com/rosbit/go-wxmp-api"
	"github.com/rosbit/go-wxmp-api/auth"
)

func main() {
	wxmpapi.SetTokenStorePath("/path/to/store/token")
	wxmpapi.SetWxmpConf("cfg1", "appId1", "appKey1")
	wxmpapi.SetWxmpConf("cfg2", "appId2", "appKey2")

	session, err := auth.Code2Session("cfg1", "code-from-front-end")
	// usage of session
}
```
