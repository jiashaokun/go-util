package str_util

import "fmt"

type StrUtil struct{}

func (s *StrUtil) InArray(str string, data []string) bool {
	if str == "" || len(data) == 0 {
		return false
	}

	for _, v := range data {
		if str == v {
			return true
		}
	}

	return false
}

func (s *StrUtil) Unique(data []string) []string {
	if len(data) <= 1 {
		return data
	}

	var res []string
	for _, v := range data {
		if ok := s.InArray(v, res); ok {
			continue
		}
		res = append(res, v)
	}

	return res
}

//interface 排重 in []interface{1, "1", 2, "2", "a", "a", 1.12, "1.12"} return []interface{"1", "2", "a"}
func (s *StrUtil) UniqueAny(data []interface{}) []string {
	var resp []string
	if len(data) == 0 {
		return resp
	}

	var middle []string
	for _, v := range data {
		val := fmt.Sprintf("%v", v)
		middle = append(middle, val)
	}

	return s.Unique(middle)
}
