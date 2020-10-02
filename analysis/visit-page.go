// 访问页面。目前只提供按 page_visit_pv 排序的 top200。
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitPage.html

package analysis

import (
	"github.com/rosbit/go-wxmp-api/auth"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"time"
	"fmt"
)

type VisitPage struct {
	RefDate string `json:"ref_date"`
	List []struct {
		PagePath    string `json:"page_path"`      // 页面路径
		PageVisitPV uint   `json:"page_visit_pv"`  // 访问次数
		PageVisitUV uint   `json:"page_visit_uv"`  // 访问人数
		PageStaytimePV float64 `json:"page_staytime_pv"` // 次均停留时长
		EntrypagePV uint   `json:"entrypage_pv"`   // 进入页次数
		ExitpagePV  uint   `json:"exitpage_pv"`    // 退出页次数
		PageSharePV uint   `json:"page_share_pv"`  // 转发次数
		PageShareUV uint   `json:"page_share_uv"`  // 转发人数
	} `json:"list"`
}

// days为“DST_1Day”时，beginDate必须大于今日
func GetVisitPage(cfgName string, beginDate time.Time, days DaysSpanType) (*VisitPage, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage?access_token=%s", accessToken)
		body = getDaysRange(beginDate, days)
		return
	}

	var res struct {
		callwxmp.BaseResult
		VisitPage
	}
	if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.JsonCall, &res); err != nil {
		return nil, err
	}
	return &res.VisitPage, nil
}
