# go-util 

### 功能列表
1. 整形类型 int_util
   1. 检测一个 int 是否在 list 中
   2. 排重
   3. 返回最大 和 最小 值
   4. 返回 []int{} 中元素的 ++ 值
   5. 返回 []int{} 中元素的 -- 值 比如：5-4-3-2-1
   6. 找到某个数字在数组中重复出现的次数 in 2 []{1,2,2} return 2
2. 字符串类型 str_util
   1. InArray 字符串是否在 []string 中
   2. Unique 去重
   3. UniqueAny 多维度去重，输入 []interface{1, "1", 2, "2", "a", "a", 1.12, "1.12"} 返回 []interface{"1", "2", "a"}
   4. ArrayFilter 过滤掉某个元素
   5. TestArrayFilterAny 过滤掉多个元素
   6. ArrayKeys 返回 key 列表
3. 时间类型
   1. BetweenDayStr 返回两段时间的日期 输入 2021-01-01 11:00:00  2021-01-02 11:00:00  输出 []string{"2021-01-01", "2021-01-02"}
   2. MonthOneDay 返回某个月的第一天的零点
   3. MonthEndDay 返回 一个月的最后一天的 零点
   4. IntervalWeekDay 已周为单位,把相同周的日期放在同一个组，分割时间
   5. WithDay 获取某个时间所在周的周几
   6. WithDayListDay 获取一段时间内的周几
   7. TimeIntersectionInfo 返回两端时间是否有交集
   8. IntersectionTime 两个时间段的交集时间
   9. DateWeekThis 获取某一天是不是本周，是本周的周几 
   10. MothDays 获取某个月的某些天 input []int{0, 1, 32} return []time.Time{2022-01-01 .... 2022-01-31}
   11. TimeStringToUnix 时间字符串类型转 time.Time 类型
   12. DayTwiceTime 获取某天的两端时间 0 点和 23:59:59

### demo
```go
package go_util

import "testing"

func TestStrUtilInArray(t *testing.T) {
	util := new(Util)
	if ok := util.StrUtil.InArray("a", []string{"a", "b"}); !ok {
		t.Fatalf("str InArray was err")
	}
}
```
