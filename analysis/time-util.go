package analysis

import (
	"github.com/rosbit/go-wxmp-api/conf"
	"time"
)

const (
	_dateLayout = "20060102"
)

type DaysSpanType uint8 // 日期跨度
const (
	DST_1Day DaysSpanType = iota  // 1日
	DST_7Days                     // 7日
	DST_30Days                    // 30日
)

func getDailyRange(date *time.Time) (beginDate, endDate string) {
		// 约束条件: end_date: 结束日期，限定查询1天数据，允许设置的最大值为昨日。格式为 yyyymmdd
		if date == nil {
			now := time.Now().AddDate(0, 0, -1) // 取昨日
			date = &now
		}

		beginDate = date.Format(_dateLayout)
		endDate = beginDate
		return
}

func getMonthlyRange(date *time.Time) (beginDate, endDate string) {
		if date == nil {
			now := time.Now()
			date = &now
		}
		year, month, _ := date.Date()
		bd := time.Date(year, month, 1, 0, 0, 0, 0, conf.Loc)
		ed := bd.AddDate(0, 1, 0).AddDate(0, 0, -1) // beginDate -> 下月1日 -> 前一日(本月最后1日)
		beginDate = bd.Format(_dateLayout)
		endDate = ed.Format(_dateLayout)
		return
}

func getWeeklyRange(date *time.Time) (beginDate, endDate string) {
		if date == nil {
			now := time.Now()
			date = &now
		}

		adjustDay := 0
		switch weekDay := date.Weekday(); weekDay {
		case time.Sunday:
			adjustDay = -6
		default:
			adjustDay = int(weekDay - 1)
		}
		bd := date.AddDate(0, 0, adjustDay)
		ed := bd.AddDate(0, 0, 6)
		beginDate = bd.Format(_dateLayout)
		endDate = ed.Format(_dateLayout)
		return
}

func getDaysRange(beginDate time.Time, days DaysSpanType) (map [string]interface{}) {
		var endDate time.Time
		switch days {
		case DST_7Days:
			endDate = beginDate.AddDate(0, 0, 6)
		case DST_30Days:
			endDate = beginDate.AddDate(0, 0, 29)
		// case DST_1Day:
		default:
			endDate = beginDate
		}
		return map[string]interface{}{
			"begin_date": beginDate.Format(_dateLayout),
			"end_date": endDate.Format(_dateLayout),
		}
}

func genRangeBody(date *time.Time, getRange func(*time.Time)(string, string)) map[string]interface{} {
	beginDate, endDate := getRange(date)
	return map[string]interface{}{
		"begin_date": beginDate,
		"end_date": endDate,
	}
}
