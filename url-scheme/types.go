package urlscheme

type ExpireType = uint8
const (
	ExpireTypeTime ExpireType = iota
	ExpireTypeDayInterval
)

type JumpWxa struct {
	Path string  `json:"path"`  // 通过 scheme 码进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带 query。path 为空时会跳转小程序主页。
	Query string `json:"query"` // 通过 scheme 码进入小程序时的 query，最大1024个字符，只支持数字，大小写英文以及部分特殊字符：`!#$&'()*+,/:;=?@-._~%``
	EnvVersion string `json:"env_version"` // 要打开的小程序版本。正式版为"release"，体验版为"trial"，开发版为"develop"，仅在微信外打开时生效。
}

type SchemeInfo struct {
	AppId string `json"appid"` // 小程序 appid。
	Path  string `json:"path"` // 小程序页面路径。
	Query string `json:"query"` // 小程序页面query。
	CreateTime int64 `json:"create_time"` // 创建时间，为 Unix 时间戳。
	ExpireTime int64 `json:"expire_time"` // 到期失效时间，为 Unix 时间戳，0 表示永久生效
	EnvVersion string `json:"env_version"` // 要打开的小程序版本。正式版为"release"，体验版为"trial"，开发版为"develop"。
}
