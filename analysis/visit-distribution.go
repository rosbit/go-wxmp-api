// 获取用户小程序访问分布数据
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitDistribution.html

package analysis

import (
	"github.com/rosbit/go-wxmp-api/auth"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"time"
	"fmt"
)

type VisitDistributeItem struct {
	Key     uint `json:"key"`                    // 场景 id，定义在各个 index 下不同
	VisitPV uint `json:"value"`                  // 该场景 id 访问 pv
	VisitUV uint `json:"access_source_visit_uv"` // 该场景 id 访问 uv
}

type VisitDistribution struct {
	RefDate string `json:"ref_date"`  // 日期，格式为 yyyymmdd
	List []struct {
		Index string `json:"index"`   // 分布类型, access_source_session_cnt/access_staytime_info/access_depth_info
		Items []VisitDistributeItem `json"item_list"` // 分布数据列表
	} `json:"list"`
}

// days为“DST_1Day”时，beginDate必须大于今日
func GetVisitDistribution(cfgName string, beginDate time.Time, days DaysSpanType) (*VisitDistribution, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution?access_token=%s", accessToken)
		body = getDaysRange(beginDate, days)
		return
	}

	var res struct {
		callwxmp.BaseResult
		VisitDistribution
	}
	if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.JsonCall, &res); err != nil {
		return nil, err
	}
	return &res.VisitDistribution, nil
}
