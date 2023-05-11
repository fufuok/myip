package myip

// SearchString 搜索字符串位置(左, 第一个)
func SearchString(ss []string, s string) int {
	for i := range ss {
		if s == ss[i] {
			return i
		}
	}
	return -1
}

// InStrings 检查字符串是否存在于 slice
func InStrings(ss []string, s string) bool {
	return SearchString(ss, s) != -1
}
