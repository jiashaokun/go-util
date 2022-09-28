package int_util

import (
	"fmt"
	"testing"
)

func TestInArray(t *testing.T) {
	util := IntUtil{}
	if ok := util.InArray(5, []int{1, 2, 3, 5}); !ok {
		t.Fatalf("InArray was err")
	}
}

func TestUnique(t *testing.T) {
	util := IntUtil{}

	res := util.Unique([]int{1, 2, 1})
	if len(res) != 2 {
		t.Fatalf("int Unique was err")
	}
}

func TestMaxMin(t *testing.T) {
	util := IntUtil{}
	res := util.MaxMin([]int{2, 1, 2, 5, 4})
	fmt.Println(res)
	if res.Min != 1 || res.Max != 5 {
		t.Fatalf("Int MaxMin was err")
	}
}

func TestPlus(t *testing.T) {
	util := IntUtil{}
	res := util.Plus([]int{1, 2, 3})
	fmt.Println(res)
	if res != 6 {
		t.Fatalf("Int Plus was err")
	}
}

func Test(t *testing.T) {
	util := IntUtil{}
	res := util.RepeatNum(2, []int{1, 2, 2})
	fmt.Println(res)
	if res != 2 {
		t.Fatalf("Int RepeatNum was err")
	}
}

func TestReduce(t *testing.T) {
	util := IntUtil{}
	res := util.Reduce([]int{2, 3, 1, 4, 5})
	if res != -5 {
		fmt.Println("Reduce was err")
	}
}
