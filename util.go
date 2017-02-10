package Util

func InArray(s interface{}, d map[string]string) int {
	for _, v := range d {
		if s == v {
			return true
		}
	}
	return false
}

func ArrayKeys(s string, d map[string]string) int {
	for k, _ := range d {
		if s == d {
			return true
		}
	}
	return false
}
