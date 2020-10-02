// 获取小程序新增或活跃用户的画像分布数据。时间范围支持昨天、最近7天、最近30天。其中，新增用户数为时间范围内首次访问小程序的去重用户数，活跃用户数为时间范围内访问过小程序的去重用户数。
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html

package analysis

import (
	"github.com/rosbit/go-wxmp-api/auth"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"time"
	"fmt"
)

type PortraitItem struct {
	Id    uint   `json:"id"`    // 属性值id
	Name  string `json:"name"`  // 属性值名称，与id对应。如属性为 province 时，返回的属性值名称包括「广东」等。
	Value uint   `json:"access_source_visit_uv"` // 该场景访问uv
}

type VisitUVT struct {
	Index     int            `json:"index"`     // 分布类型
	Province  []PortraitItem `json:"province"`  // 省份，如北京、广东等
	City      []PortraitItem `json:"city"`      // 城市，如北京、广州等
	Genders   []PortraitItem `json:"genders"`   // 性别，包括男、女、未知
	Platforms []PortraitItem `json:"platforms"` // 终端类型，包括 iPhone，android，其他
	Devices   []PortraitItem `json:"devices"`   // 机型，如苹果 iPhone 6，OPPO R9 等
	Ages      []PortraitItem `json:"ages"`      // 年龄，包括17岁以下、18-24岁等区间
}

type UserPortrait struct {
	RefDate    string   `json:"ref_date"`     // 时间范围，如："20170611-20170617"
	VisitUVNew VisitUVT `json:"visit_uv_new"` // 新用户画像
	VisitUV    VisitUVT `json:"visit_uv"`     // 活跃用户画像
}

// days为“DST_1Day”时，beginDate必须大于今日
func GetUserPortrait(cfgName string, beginDate time.Time, days DaysSpanType) (*UserPortrait, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/datacube/getweanalysisappiduserportrait?access_token=%s", accessToken)
		body = getDaysRange(beginDate, days)
		return
	}

	var res struct {
		callwxmp.BaseResult
		UserPortrait
	}
	if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.JsonCall, &res); err != nil {
		return nil, err
	}
	return &res.UserPortrait, nil
}
