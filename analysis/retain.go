package analysis

import (
	"github.com/rosbit/go-wxmp-api/auth"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"time"
	"fmt"
)

type RetainItem struct {
	Key   uint // 标识，0开始，表示当天(月、周)，1表示1天(月、周)后。依此类推，key取值分别是：0,1,2,3,4,5,6,7,14,30
	Value uint // key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

type RetainRes struct {
	RefData    string       `json:"ref_date"`     // 日期
	VisitUVNew []RetainItem `json:"visit_uv_new"` // 新增用户留存
	VisitUV    []RetainItem `json:"visit_uv"`     // 活跃用户留存
}

// 获取用户访问小程序日留存
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html
// date 为nil表示昨日
func GetDailyRetain(cfgName string, date *time.Time) (*RetainRes, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token=%s", accessToken)
		body = genRangeBody(date, getDailyRange)
		return
	}

	return getRetain(cfgName, genParams)
}

// 获取用户访问小程序月留存
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html
// date 如果为nil，表示当日所在的月；如果为非nil，表示日期所在的月
func GetMonthlyRetain(cfgName string, date *time.Time) (*RetainRes, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo?access_token=%s", accessToken)
		body = genRangeBody(date, getMonthlyRange)
		return
	}

	return getRetain(cfgName, genParams)
}

// 获取用户访问小程序周留存
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html
// date 如果为nil，表示当日所在的周；如果为非nil，表示日期所在的周
func GetWeeklyRetain(cfgName string, date *time.Time) (*RetainRes, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo?access_token=%s", accessToken)
		body = genRangeBody(date, getWeeklyRange)
		return
	}

	return getRetain(cfgName, genParams)
}

func getRetain(cfgName string, genParams auth.FnGeneParams) (*RetainRes, error) {
	var res struct {
		callwxmp.BaseResult
		RetainRes
	}
	if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.JsonCall, &res); err != nil {
		return nil, err
	}
	return &res.RetainRes, nil
}
