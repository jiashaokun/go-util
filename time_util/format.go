package time_util

import "time"

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
