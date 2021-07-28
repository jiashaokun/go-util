package str_util

import (
	"fmt"
	"testing"
)

func TestInArray(t *testing.T) {
	util := StrUtil{}
	if ok := util.InArray("a", []string{"a", "b"}); !ok {
		t.Fatalf("str InArray was err")
	}
}

func TestUnique(t *testing.T) {
	util := StrUtil{}

	res := util.Unique([]string{"a", "b", "a"})
	if len(res) != 2 {
		t.Fatalf("Str Unique was err")
	}
	fmt.Println(res)
}

func TestUniqueAny(t *testing.T) {
	util := StrUtil{}
	res := util.UniqueAny([]interface{}{1, "1", 2, "2", "a", "a", 1.12, "1.12"})
	fmt.Println(res)
	if len(res) != 4 {
		t.Fatalf("Str UniqueAny was err %s", res)
	}
}

func TestIndex(t *testing.T) {
	util := StrUtil{}
	if key := util.Index("5", []interface{}{1, "5"}); key != 1 {
		t.Fatalf("Str Index was err")
	}
}

func TestArrayKeys(t *testing.T) {
	util := StrUtil{}

	info := map[string]string{
		"a": "a",
		"b": "b",
	}
	res := util.ArrayKeys(info)
	fmt.Println(res)
	if len(res) != 2 {
		t.Fatalf("Str ArrayKeys was err")
	}
}
