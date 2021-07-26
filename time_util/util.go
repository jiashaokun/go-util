package time_util

import (
	"time"
)

type TimeUtil struct{}

const localDayStr = "2006-01-02"
const localTimeStr = "2006-01-02 15:04:05"

//返回两段时间的日期 in 2021-01-01 11:00:00  2021-01-02 11:00:00  return []string{"2021-01-01", "2021-01-02"}
func (t *TimeUtil) BetweenDayStr(str, end time.Time) []string {
	var strS, endS string
	if str.IsZero() {
		strS = time.Now().Format(localDayStr)
	} else {
		strS = str.Format(localDayStr)
	}

	if end.IsZero() {
		endS = time.Now().Format(localDayStr)
	} else {
		endS = end.Format(localDayStr)
	}

	if strS == endS {
		return []string{endS}
	}

	var days []string
	days = append(days, strS)
	listDay, _ := time.Parse(localDayStr, strS)
	for {
		timeNextDay := listDay.AddDate(0, 0, 1)
		day := timeNextDay.Format(localDayStr)
		days = append(days, day)

		listDay, _ = time.Parse(localDayStr, day)
		if day == endS {
			break
		}
	}

	return days
}

//返回某个月的第一天的零点 in 2021-01-05 return 2021-01-01
func (t *TimeUtil) MonthOneDay(today time.Time) time.Time {
	var todayStr string
	if today.IsZero() {
		todayStr = time.Now().Format(localDayStr)
	} else {
		todayStr = today.Format(localDayStr)
	}

	td, _ := time.Parse(localDayStr, todayStr)

	d := td.Day()
	mc := d - 1
	if mc == 0 {
		return td
	}

	return td.AddDate(0, 0, -mc)

}

//返回 一个月的最后一天的 零点
func (t *TimeUtil) MonthEndDay(today time.Time) time.Time {
	var todayStr string
	if today.IsZero() {
		todayStr = time.Now().Format(localDayStr)
	} else {
		todayStr = today.Format(localDayStr)
	}

	td, _ := time.Parse(localDayStr, todayStr)

	nextMonth := td.AddDate(0, 1, 0)

	nextOneDay := t.MonthOneDay(nextMonth)

	resDay := nextOneDay.AddDate(0, 0, -1)

	return resDay
}
