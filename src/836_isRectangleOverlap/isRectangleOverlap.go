package _36_isRectangleOverlap

import "math"

/*
	矩形以列表 [x1, y1, x2, y2] 的形式表示，其中 (x1, y1) 为左下角的坐标，(x2, y2) 是右上角的坐标。

	如果相交的面积为正，则称两矩形重叠。需要明确的是，只在角或边接触的两个矩形不构成重叠。

	给出两个矩形，判断它们是否重叠并返回结果。

	示例 1：

	输入：rec1 = [0,0,2,2], rec2 = [1,1,3,3]
	输出：true
	示例 2：

	输入：rec1 = [0,0,1,1], rec2 = [1,0,2,1]
	输出：false
	说明：

	两个矩形 rec1 和 rec2 都以含有四个整数的列表的形式给出。
	矩形中的所有坐标都处于 -10^9 和 10^9 之间

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/rectangle-overlap
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if len(rec1) < 4 || len(rec2) < 4 {
		return false
	}
	center1_x := (rec1[2] + rec1[0]) / 2
	center1_y := (rec1[3] + rec1[1]) / 2
	half1_x := (rec1[2] - rec1[0]) / 2
	half1_y := (rec1[3] - rec1[1]) / 2

	center2_x := (rec2[2] + rec2[0]) / 2
	center2_y := (rec2[3] + rec2[1]) / 2
	half2_x := (rec2[2] - rec2[0]) / 2
	half2_y := (rec2[3] - rec2[1]) / 2
	return math.Sqrt(float64((center1_x - center2_x) * (center1_x - center2_x) + (center1_y - center2_y) *(center1_y - center2_y))) <
			math.Sqrt(float64(half2_x * half2_x + half2_y * half2_y)) +math.Sqrt(float64(half1_x * half1_x + half1_y * half1_y))
}