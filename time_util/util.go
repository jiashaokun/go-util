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
