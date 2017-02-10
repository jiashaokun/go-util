package Util

func InArray(s interface{}, d map[string]string) int {
	for _, v := range d {
		if s == v {
			return 1
		}
	}
	return 0
}

func ArrayKeys(s string, d map[string]string) int {
	for k, _ := range d {
		if s == d {
			return true
		}
	}
	return false
}

func ArrayColumn(d map[int]map[string]string, column_key string, index_key string) map[int]map[string]string {
	nd := make(map[int]map[string]string)
	for k, v := range d {
		for e, q := range v {
			if e == index_key {
				nd[k][index_key] = q
			}
		}
	}
	return nd
}
