package time_util

import (
	"fmt"
	"testing"
	"time"
)

func TestBetweenDayStr(t *testing.T) {
	util := TimeUtil{}

	strTime := time.Now()
	endTime := time.Now().AddDate(0, 0, 10)
	ls := util.BetweenDayStr(strTime, endTime)

	fmt.Println(ls)
	if len(ls) != 11 {
		t.Fatalf("TimeUtil BetweenDayStr was err")
	}
}

func TestMonthOneDay(t *testing.T) {
	util := TimeUtil{}

	res := util.MonthOneDay(time.Now())
	fmt.Println(res)
}

func TestMonthEndDay(t *testing.T) {
	util := TimeUtil{}

	res := util.MonthEndDay(time.Now())
	fmt.Println(res)
}

func TestIntervalWeekDay(t *testing.T) {
	str := "2021-08-03"
	st, _ := time.Parse(localDayStr, str)
	end := "2021-08-16"
	ed, _ := time.Parse(localDayStr, end)

	util := TimeUtil{}
	res := util.IntervalWeekDay(st, ed, 0)
	fmt.Println(res)
	if len(res) != 3 {
		t.Fatalf("TestIntervalWeekDay was err")
	}
}

func TestWithDay(t *testing.T) {
	str := "2021-08-03"
	st, _ := time.Parse(localDayStr, str)
	util := TimeUtil{}
	res := util.WithDay(st, 0)
	fmt.Println(res)
}

func TestTimeIntersection(t *testing.T) {
	util := TimeUtil{}
	os := "2021-01-02 12:00:00"
	oe := "2021-01-02 14:00:00"

	as := "2021-01-02 12:00:00"
	ae := "2021-01-02 14:00:00"
	p1 := TimeIntersectionInfo{
		FirstTimeStart:  timeFt(os),
		FirstTimeEnd:    timeFt(oe),
		SecondTimeStart: timeFt(as),
		SecondTimeEnd:   timeFt(ae),
	}
	res1 := util.TimeIntersection(p1)
	fmt.Println(res1)

	bs := "2021-01-02 11:00:00"
	be := "2021-01-02 15:00:00"
	p2 := TimeIntersectionInfo{
		FirstTimeStart:  timeFt(os),
		FirstTimeEnd:    timeFt(oe),
		SecondTimeStart: timeFt(bs),
		SecondTimeEnd:   timeFt(be),
	}
	res2 := util.TimeIntersection(p2)
	fmt.Println(res2)

	cs := "2021-01-02 12:10:00"
	ce := "2021-01-02 13:00:00"
	p3 := TimeIntersectionInfo{
		FirstTimeStart:  timeFt(os),
		FirstTimeEnd:    timeFt(oe),
		SecondTimeStart: timeFt(cs),
		SecondTimeEnd:   timeFt(ce),
	}
	res3 := util.TimeIntersection(p3)
	fmt.Println(res3)
}

func timeFt(t string) time.Time {
	t2, _ := time.Parse(localTimeStr, t)

	return t2
}

func TestIntersectionTime(t *testing.T) {
	util := TimeUtil{}

	os := "2021-01-02 14:00:00"
	oe := "2021-01-02 14:01:00"

	as := "2021-01-02 13:00:00"
	ae := "2021-01-02 15:00:00"
	p1 := TimeIntersectionInfo{
		FirstTimeStart:  timeFt(os),
		FirstTimeEnd:    timeFt(oe),
		SecondTimeStart: timeFt(as),
		SecondTimeEnd:   timeFt(ae),
	}

	res := util.IntersectionTime(p1)
	fmt.Println(res)
}

func TestWithDayListDay(t *testing.T) {
	util := TimeUtil{}
	startDate, _ := time.ParseInLocation(localDayStr, "2022-09-01", time.Local)
	endDate, _ := time.ParseInLocation(localTimeStr, "2022-09-30 19:01:01", time.Local)
	days := util.WithDayListDay(startDate, endDate, 0)
	fmt.Println(days)
}
