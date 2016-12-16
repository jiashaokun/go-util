package util

//type Table struct{}

func Intable(v interface{}, c []int) bool {
	if len(c) == 0 {
		return false
	}
	for _, vv := range c {
		if v == vv {
			return true
		}
	}
	return false
}
