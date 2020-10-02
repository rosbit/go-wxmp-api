package analysis

import (
	"github.com/rosbit/go-wxmp-api/auth"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"time"
	"fmt"
)

type VisitTrend struct {
	RefDate       string `json:"ref_date"`          // 日期，格式为 yyyymmdd
	SessionCount    uint `json:"session_cnt"`       // 打开次数
	VisitPV         uint `json:"visit_pv"`          // 访问次数
	VisitUV         uint `json:"visit_uv"`          // 访问人数
	VisitUVNew      uint `json:"visit_uv_new"`      // 新用户数
	StayTimeUV      float64 `json:"stay_time_uv"`      // 人均停留时长 (浮点型，单位：秒)
	StayTimeSession float64 `json:"stay_time_session"` // 次均停留时长 (浮点型，单位：秒)
	VisitDepth      float64 `json:"visit_depth"`       // 平均访问深度 (浮点型)
}

// 获取用户访问小程序数据日趋势
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getDailyVisitTrend.html
// date 为nil表示取昨日
func GetDailyVisitTrend(cfgName string, date *time.Time) ([]VisitTrend, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend?access_token=%s", accessToken)
		body = genRangeBody(date, getDailyRange)
		return
	}

	return getVisitTrend(cfgName, genParams)
}

// 获取用户访问小程序数据月趋势(能查询到的最新数据为上一个自然月的数据)
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getMonthlyVisitTrend.html
// date 如果为nil，表示当日所在的月；如果为非nil，表示日期所在的月
func GetMonthlyVisitTrend(cfgName string, date *time.Time) ([]VisitTrend, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyvisittrend?access_token=%s", accessToken)
		body = genRangeBody(date, getMonthlyRange)
		return
	}

	return getVisitTrend(cfgName, genParams)
}

// 获取用户访问小程序数据周趋势
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getWeeklyVisitTrend.html
// date 如果为nil，表示当日所在的周；如果为非nil，表示日期所在的周
func GetWeeklyVisitTrend(cfgName string, date *time.Time) ([]VisitTrend, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/datacube/getweanalysisappidweeklyvisittrend?access_token=%s", accessToken)
		body = genRangeBody(date, getWeeklyRange)
		return
	}

	return getVisitTrend(cfgName, genParams)
}

func getVisitTrend(cfgName string, genParams auth.FnGeneParams) ([]VisitTrend, error) {
	var res struct {
		callwxmp.BaseResult
		List []VisitTrend `json:"list"`
	}
	if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.JsonCall, &res); err != nil {
		return nil, err
	}
	return res.List, nil
}
