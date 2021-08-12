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

/*
获取两个日期段内的所有的周数据
dayType = 1
in startTime :  2021-08-03 endTime : 2021-08-16
rteutn [{[2021-08-02 00:00:00 +0000 UTC 2021-08-03 00:00:00 +0000 UTC 2021-08-04 00:00:00 +0000 UTC 2021-08-05 00:00:00 +0000 UTC 2021-08-06 00:00:00 +0000 UTC 2021-08-07 00:00:00 +0000 UTC 2021-08-08 00:00:00 +0000 UTC]} {[2021-08-09 00:00:00 +0000 UTC 2021-08-10 00:00:00 +0000 UTC 2021-08-11 00:00:00 +0000 UTC 2021-08-12 00:00:00 +0000 UTC 2021-08-13 00:00:00 +0000 UTC 2021-08-14 00:00:00 +0000 UTC 2021-08-15 00:00:00 +0000 UTC]} {[2021-08-16 00:00:00 +0000 UTC 2021-08-17 00:00:00 +0000 UTC 2021-08-18 00:00:00 +0000 UTC 2021-08-19 00:00:00 +0000 UTC 2021-08-20 00:00:00 +0000 UTC 2021-08-21 00:00:00 +0000 UTC 2021-08-22 00:00:00 +0000 UTC]}]


dayType = 0
in startTime : 2021-08-03 endTime : 2021-08-16
rteutn [{[2021-08-03 00:00:00 +0000 UTC 2021-08-04 00:00:00 +0000 UTC 2021-08-05 00:00:00 +0000 UTC 2021-08-06 00:00:00 +0000 UTC 2021-08-07 00:00:00 +0000 UTC 2021-08-08 00:00:00 +0000 UTC]} {[2021-08-09 00:00:00 +0000 UTC 2021-08-10 00:00:00 +0000 UTC 2021-08-11 00:00:00 +0000 UTC 2021-08-12 00:00:00 +0000 UTC 2021-08-13 00:00:00 +0000 UTC 2021-08-14 00:00:00 +0000 UTC 2021-08-15 00:00:00 +0000 UTC]} {[2021-08-16 00:00:00 +0000 UTC]}]

*/

type WeekDayInfo struct {
	DateList []time.Time
}

func (t *TimeUtil) IntervalWeekDay(startTime, endTime time.Time, dayType int) []WeekDayInfo {
	//获取开始时间的周一 和 结束时间的 周末
	strDate := startTime
	if dayType == 1 {
		d := int(startTime.Weekday())
		if d == 0 {
			d = 7
		}
		if d > 1 {
			//获取周一
			strDate = startTime.AddDate(0, 0, -(d - 1))

		}
	}

	endDate := endTime
	//获取 endTime 所在的周末
	if dayType == 1 {
		ed := int(endTime.Weekday())
		if ed == 0 {
			ed = 7
		}
		ced := 7 - ed
		endDate = endTime.AddDate(0, 0, ced)
	}

	days := t.BetweenDayStr(strDate, endDate)

	var resp []WeekDayInfo
	//组装一个 map 已周一 为 key

	weekDict := make(map[string][]time.Time)
	for _, v := range days {
		today, _ := time.Parse(localDayStr, v)
		//获取周一的时间
		td := int(today.Weekday())
		one := today
		if td == 0 {
			td = 7
		}
		if td > 1 {
			one = today.AddDate(0, 0, -(td - 1))
		}

		oneStr := one.Format(localDayStr)
		weekDict[oneStr] = append(weekDict[oneStr], today)
	}

	for _, v := range weekDict {
		info := WeekDayInfo{
			DateList: v,
		}

		resp = append(resp, info)
	}

	return resp
}

/*
获取某个时间的周几
dayType: 1:获取周一时间 2:获取周二时间......0:获取周日时间
*/
func (t *TimeUtil) WithDay(tm time.Time, dayType int) time.Time {
	if dayType < 0 || dayType > 6 {
		return tm
	}
	if dayType == 0 {
		dayType = 7
	}

	wd := int(tm.Weekday())
	if wd == 0 {
		wd = 7
	}

	d := tm
	if wd > 1 {
		d = tm.AddDate(0, 0, -(wd - 1))
	}

	d = d.AddDate(0, 0, dayType-1)

	return d
}
