package str_util

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"strings"
)

type StrUtil struct{}

func (s *StrUtil) InArray(str string, data []string) bool {

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

// interface 排重 in []interface{1, "1", 2, "2", "a", "a", 1.12, "1.12"} return []interface{"1", "2", "a"}
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

// ArrayFilter info 为空 则 过滤 空值  input []string{"", ""} return []string{}
func (s *StrUtil) ArrayFilter(data []string, info string) []string {
	var resp []string
	for _, v := range data {
		val := strings.Trim(v, " ")
		if info != val {
			resp = append(resp, val)
		}
	}

	return resp
}

// ArrayFilterAny 过滤多个值
func (s *StrUtil) ArrayFilterAny(data []string, info []string) []string {
	var resp []string
	for _, v := range data {
		if !s.InArray(v, info) {
			resp = append(resp, v)
		}
	}

	return resp
}

// Index 返回某个数据的 key 值
func (s *StrUtil) Index(find interface{}, data []interface{}) []int {
	var res []int
	if len(data) == 0 {
		return res
	}

	var fdBuf bytes.Buffer
	fEnc := gob.NewEncoder(&fdBuf)
	fEnc.Encode(find)
	fd := fdBuf.Bytes()
	fes := fmt.Sprintf("%x", md5.Sum(fd))
	//把值 md5 如果完全相等 那就是相等
	for k, v := range data {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		enc.Encode(v)
		d := buf.Bytes()
		es := fmt.Sprintf("%x", md5.Sum(d))
		if fes == es {
			res = append(res, k)
		}
	}

	return res
}

// 返回一个map的 keys in map[string]string{} return []string{}
func (s *StrUtil) ArrayKeys(data map[string]string) []string {
	var resp []string
	if len(data) == 0 {
		return resp
	}

	for k, _ := range data {
		resp = append(resp, k)
	}

	return resp
}
