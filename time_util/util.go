package time_util

import (
	"time"
)

type localTime string

type TimeUtil struct {
	Local localTime
}

const localDayStr = "2006-01-02"
const localTimeStr = "2006-01-02 15:04:05"

var dateWeekMap = map[time.Weekday]string{
	time.Monday:    "周一",
	time.Tuesday:   "周二",
	time.Wednesday: "周三",
	time.Thursday:  "周四",
	time.Friday:    "周五",
	time.Saturday:  "周六",
	time.Sunday:    "周日",
}

// BetweenDayStr 返回两段时间的日期 in 2021-01-01 11:00:00  2021-01-02 11:00:00  return []string{"2021-01-01", "2021-01-02"}
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

// MonthOneDay 返回某个月的第一天的零点 in 2021-01-05 return 2021-01-01
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

// MonthEndDay 返回 一个月的最后一天的 零点
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

// IntervalWeekDay 已周为单位,把相同周的日期放在同一个组，分割时间
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

// WithDay 获取某个时间的周几。dayType: 1:获取周一时间 2:获取周二时间......0:获取周日时间
func (t *TimeUtil) WithDay(tm time.Time, dayType time.Weekday) time.Time {
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

	d = d.AddDate(0, 0, int(dayType)-1)

	return d
}

// WithDayListDay 获取一段时间内的周几：0：星期日。输入：startDate = 2022-09-01, endDate = 2022-09-20; day = 0(星期日) return [2022-01-07, 2022-01-14]
func (t *TimeUtil) WithDayListDay(startDate, endDate time.Time, day time.Weekday) []time.Time {
	var resp []time.Time
	dayList := t.BetweenDayStr(startDate, endDate)
	for _, tmStr := range dayList {
		tm, _ := time.ParseInLocation(localDayStr, tmStr, time.Local)
		if tm.Weekday() == day {
			resp = append(resp, tm)
		}
	}

	return resp
}

// TimeIntersectionInfo 返回两端时间是否有交集
type TimeIntersectionInfo struct {
	FirstTimeStart  time.Time
	FirstTimeEnd    time.Time
	SecondTimeStart time.Time
	SecondTimeEnd   time.Time
}

func (t *TimeUtil) TimeIntersection(info TimeIntersectionInfo) bool {
	//1. A B A B
	if info.SecondTimeStart.After(info.FirstTimeStart) && info.SecondTimeStart.Before(info.SecondTimeEnd) {
		return true
	}
	//2. A B B A
	if info.FirstTimeStart.Before(info.SecondTimeStart) && info.FirstTimeEnd.After(info.SecondTimeEnd) {
		return true
	}
	//3. B A B A
	if info.FirstTimeStart.After(info.SecondTimeStart) && info.FirstTimeStart.Before(info.SecondTimeEnd) {
		return true
	}
	//4. B A A B
	if info.SecondTimeStart.Before(info.FirstTimeStart) && info.SecondTimeEnd.After(info.FirstTimeEnd) {
		return true
	}
	//相等
	if info.FirstTimeStart == info.SecondTimeStart || info.FirstTimeStart == info.SecondTimeEnd || info.FirstTimeEnd == info.SecondTimeStart || info.FirstTimeEnd == info.SecondTimeEnd {
		return true
	}
	return false
}

// IntersectionTimeInfo 两个时间段的交集时间
type IntersectionTimeInfo struct {
	TimeStart time.Time
	TimeEnd   time.Time
}

func (t *TimeUtil) IntersectionTime(info TimeIntersectionInfo) IntersectionTimeInfo {
	resp := IntersectionTimeInfo{}

	if info.SecondTimeEnd == info.FirstTimeEnd {
		resp.TimeEnd = info.SecondTimeEnd
	}

	if info.SecondTimeStart == info.FirstTimeStart {
		resp.TimeStart = info.FirstTimeStart
	}

	if info.FirstTimeEnd.After(info.SecondTimeStart) && info.FirstTimeEnd.Before(info.SecondTimeEnd) {
		resp.TimeEnd = info.FirstTimeEnd
	}

	if info.FirstTimeStart.After(info.SecondTimeStart) && info.FirstTimeStart.Before(info.SecondTimeEnd) {
		resp.TimeStart = info.FirstTimeStart
	}

	if info.SecondTimeEnd.After(info.FirstTimeStart) && info.SecondTimeEnd.Before(info.FirstTimeEnd) || info.SecondTimeEnd == info.FirstTimeEnd {
		resp.TimeEnd = info.SecondTimeEnd

		if info.FirstTimeStart.Before(info.SecondTimeStart) {
			resp.TimeStart = info.SecondTimeStart
		}

		if info.FirstTimeStart.After(info.SecondTimeStart) {
			resp.TimeStart = info.FirstTimeStart
		}

		return resp
	}

	if info.SecondTimeStart.After(info.FirstTimeStart) && info.SecondTimeStart.Before(info.FirstTimeEnd) || info.SecondTimeStart == info.FirstTimeStart {
		resp.TimeStart = info.SecondTimeStart

		if info.SecondTimeEnd.After(info.FirstTimeEnd) {
			resp.TimeEnd = info.FirstTimeEnd
		}

		if info.SecondTimeEnd.Before(info.FirstTimeEnd) {
			resp.TimeEnd = info.SecondTimeEnd
		}

		return resp
	}

	if info.SecondTimeEnd == info.FirstTimeStart {
		resp.TimeStart = info.SecondTimeEnd
		resp.TimeEnd = info.SecondTimeEnd
	}

	if info.FirstTimeEnd == info.SecondTimeStart {
		resp.TimeStart = info.FirstTimeEnd
		resp.TimeEnd = info.FirstTimeEnd
	}

	return resp
}

type DateWeeks struct {
	Date     time.Time //传入的日期
	ThisWeek bool      //是否是本周 true 是 false 不是
	Week     string    //周几
}

// DateWeekThis 获取某一天是不是本周，是本周的周几
func (t *TimeUtil) DateWeekThis(date *time.Time) DateWeeks {
	var resp DateWeeks
	if date.IsZero() {
		return resp
	}
	resp.Date = *date

	_, d := date.ISOWeek()
	_, n := time.Now().ISOWeek()

	if d == n {
		resp.ThisWeek = true
	}

	resp.Week = dateWeekMap[date.Weekday()]
	return resp
}

// MothDays 返回某个月固定的某几天 input []int{1,2} return ["2021-01-01", "2021-01-02"]
func (t *TimeUtil) MothDays(days []int) []time.Time {
	var resp []time.Time
	zeroDay := t.MonthOneDay(time.Now())
	for _, v := range days {
		if v <= 0 {
			continue
		}
		day := zeroDay.AddDate(0, 0, v-1)
		//如果超过本月则不加
		if day.After(t.MonthEndDay(time.Now())) {
			continue
		}
		resp = append(resp, day)
	}
	return resp
}

type DayTwice struct {
	ZeroTime time.Time
	EndTime  time.Time
}

// DayTwiceTime 返回某个时间的 0 点 和 24 点
func (t *TimeUtil) DayTwiceTime(in time.Time) DayTwice {
	var resp DayTwice
	zero := in.Hour()
	resp.ZeroTime = in.Truncate(time.Duration(zero))
	ux := time.Unix(resp.ZeroTime.Unix()+86399, 64)
	us := ux.Format(time.DateTime)
	resp.EndTime, _ = time.ParseInLocation(time.DateTime, us, time.Local)

	return resp
}
