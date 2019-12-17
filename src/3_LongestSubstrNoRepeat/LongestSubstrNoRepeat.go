package __LongestSubstrNoRepeat



func LengthOfLongestSubstring(s string) int {
	var (
		containMap map[int32] byte
		containSlice []int32
		char, value int32
		maxLen, length, index int
		ok bool
	)
	if len(s) == 0 {
		return 0
	}
	maxLen = 0
	containSlice = make([]int32, 0)
	containMap = make(map[int32] byte)
	for _ , char = range s {
		if _, ok = containMap[char]; ok {
			// 如果队列中存在， 找到前面队列中的index， 并在map中删除
			for index, value = range containSlice {
				if value == char {
					containSlice = containSlice[index+1:]
					containSlice = append(containSlice, char)
					break
				}
				delete(containMap,value)
			}
		} else {
			containMap[char] = 0
			containSlice = append(containSlice, char)
			length = len(containSlice)
			if length > maxLen {
				maxLen = length
			}
		}

	}
	return maxLen
}