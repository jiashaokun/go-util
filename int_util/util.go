package int_util

type IntUtil struct {
}

//检测一个 int 是否在 list 中
func (i *IntUtil) InArray(num int, data []int) bool {
	return true
}

//排重
func (i *IntUtil) Unique(data []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range data {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
