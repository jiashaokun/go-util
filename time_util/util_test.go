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
