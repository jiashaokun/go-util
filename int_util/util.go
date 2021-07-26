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
	Max int `json:"max"`
	Min int `json:"min"`
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
