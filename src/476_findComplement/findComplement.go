package _76_findComplement


/*
给定一个正整数，输出它的补数。补数是对该数的二进制表示取反。

注意:

给定的整数保证在32位带符号整数的范围内。
你可以假定二进制数不包含前导零位。
示例 1:

输入: 5
输出: 2
解释: 5的二进制表示为101（没有前导零位），其补数为010。所以你需要输出2。
示例 2:

输入: 1
输出: 0
解释: 1的二进制表示为1（没有前导零位），其补数为0。所以你需要输出0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-complement
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func findComplement(num int) int {
	if num == 0 {
		return 1
	}
	ans := 0
	bit := 1
	for num != 0 {
		if num & 1 == 0  {
			ans += bit
		}
		num = num >> 1
		bit *= 2
	}
	return ans
}



func FindComplement(num int) int {
	return findComplement(num)
}