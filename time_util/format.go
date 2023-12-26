package time_util

import (
	"time"
)

// UnixToTimeString 将unix时间戳改成时间字符串
func (t *TimeUtil) UnixToTimeString(tm int64) string {
	ts := time.Unix(tm, 0)
	if t.Local == "" {
		t.Local = localTimeStr
	}
	if ts.IsZero() {
		return ""
	}
	return ts.Format(string(t.Local))
}

// TimeStringToTime 字符串时间转时间格式
func (t *TimeUtil) TimeStringToTime(tm string) *time.Time {
	ti, err := time.ParseInLocation(time.DateTime, tm, time.Local)
	if err == nil {
		return &ti
	}
	ti, err = time.ParseInLocation(time.DateOnly, tm, time.Local)
	if err == nil {
		return &ti
	}
	return nil
}
