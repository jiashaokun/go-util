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
