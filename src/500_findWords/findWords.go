package _00_findWords

/*

给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。键盘如下图所示。


示例：

输入: ["Hello", "Alaska", "Dad", "Peace"]
输出: ["Alaska", "Dad"]
 

注意：

你可以重复使用键盘上同一字符。
你可以假设输入的字符串将只包含字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/keyboard-row
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


func findWords(words []string) []string {
	maps := map[byte]byte{'q':1,'w':1,'e':1,'r':1,'t':1,'y':1,'u':1,'i':1,'o':1,'p':1,
		'a':2,'s':2,'d':2,'f':2,'g':2,'h':2,'j':2,'k':2,'l':2,
		'z':3,'x':3,'c':3,'v':3,'b':3,'n':3,'m':3}
	ans := make([]string,0)
	for _, word := range words {
		if word == "" {
			continue
		}
		var tmp byte
		if word[0] < 'a' {
			tmp = maps[word[0] + 32]
		} else  {
			tmp = maps[word[0]]
		}

		for i := 0; i < len(word); i++ {
			byte_i := word[i]
			if byte_i < 'a' {
				byte_i += 32
			}
			if maps[byte_i] != tmp {
				goto end
			}
		}
		ans = append(ans,word)
	end:
	}
	return ans
}

func FindWords(words []string) []string {
	return findWords(words)
}
