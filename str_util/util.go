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

//info 为空 则 过滤 空值  input []string{"", ""} return []string{}
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

//返回某个数据的 key 值
func (s *StrUtil) Index(find interface{}, data []interface{}) int {
	if len(data) == 0 {
		return -1
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
			fmt.Println(k)
			return k
		}
	}

	return 0
}
