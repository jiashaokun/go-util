package go_util

import "testing"

func TestStrUtilInArray(t *testing.T) {
	util := new(Util)
	if ok := util.StrUtil.InArray("a", []string{"a", "b"}); !ok {
		t.Fatalf("str InArray was err")
	}
}
