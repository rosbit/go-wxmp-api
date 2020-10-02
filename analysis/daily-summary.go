// 获取用户访问小程序数据概况
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getDailySummary.html

package analysis

import (
	"github.com/rosbit/go-wxmp-api/auth"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"time"
	"fmt"
)

type DailySummary struct {
	RefDate  string `json:"ref_date"`    // 日期，格式为 yyyymmdd
	VisitTotal uint `json:"visit_total"` // 累计用户数
	SharePV    uint `json:"share_pv"`    // 转发次数
	ShareUV    uint `json:"share_uv"`    // 转发人数
}

// date为nil表示取昨日
func GetDailySummary(cfgName string, date *time.Time) ([]DailySummary, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend?access_token=%s", accessToken)
		body = genRangeBody(date, getWeeklyRange)
		return
	}

	var res struct {
		callwxmp.BaseResult
		List []DailySummary `json:"list"`
	}
	if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.JsonCall, &res); err != nil {
		return nil, err
	}
	return res.List, nil
}
