package int_util

import (
	"github.com/jiashaokun/go-util"
	"testing"
)

func TestInArray(t *testing.T) {
	util := new(go_util.Util)
	if ok := util.IntUtil.InArray(5, []int{1, 2, 3, 5}); !ok {
		t.Fatalf("InArray was err")
	}
}

func TestUnique(t *testing.T) {
	util := new(go_util.Util)

	res := util.IntUtil.Unique([]int{1, 2, 1})
	if len(res) != 2 {
		t.Fatalf("int Unique was err")
	}
}
