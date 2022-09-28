package int_util

import (
	"sort"
)

type IntUtil struct {
}

//检测一个 int 是否在 list 中
func (i *IntUtil) InArray(num int, data []int) bool {
	if len(data) == 0 {
		return false
	}
	for _, v := range data {
		if num == v {
			return true
		}
	}
	return true
}

//排重
func (i *IntUtil) Unique(data []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range data {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//返回最大 和 最小 值
type MaxMin struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

//获取 最大 和 最小
func (i *IntUtil) MaxMin(data []int) MaxMin {
	resp := MaxMin{
		Max: 0,
		Min: 0,
	}
	if len(data) == 0 {
		return resp
	}
	s := data
	sort.Ints(s)
	resp.Max = s[len(s)-1]
	resp.Min = s[0]

	return resp
}

//++
func (i *IntUtil) Plus(data []int) int {
	if len(data) == 0 {
		return 0
	}

	var plus int
	for _, v := range data {
		plus = plus + v
	}

	return plus
}

// Reduce -- 举例：3,2,4,1,5 : 5-4-3-2-1
func (i *IntUtil) Reduce(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	sort.Ints(data)
	var max int
	for k, v := range data {
		if k == len(data)-1 {
			break
		}
		max = max + v
	}
	return data[len(data)-1] - max
}

//找到某个数字在数组中重复出现的次数 in 2 []{1,2,2} return 2
func (i *IntUtil) RepeatNum(num int, data []int) int {
	var repeatNum int
	if len(data) == 0 {
		return repeatNum
	}

	for _, v := range data {
		if num == v {
			repeatNum = repeatNum + 1
		}
	}

	return repeatNum
}
