func findContentChildren(g []int, s []int) int {
	// 从小到大排序
	sort.Ints(g)
	sort.Ints(s)
	i, j, result := 0, 0, 0
	for i < len(g) && j < len(s) {
		// 如果s[j]满足，就继续遍历下一个
		if s[j] >= g[i] {
			result = result + 1
			j++
			i++
		} else { // 如果s[j]不满足，则拿较大的和g[i]比较
			j++
		}
	}
	return result
}
