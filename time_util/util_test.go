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
