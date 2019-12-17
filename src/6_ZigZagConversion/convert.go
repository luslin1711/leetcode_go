package __ZigZagConversion

import "bytes"

func convert(s string, numRows int) string {
	var (
		index, res_index, res_flag int
		char byte
		res [][]byte
		buff *bytes.Buffer
	)
	if numRows <= 1 || numRows > len(s){
		return s
	}
	res_index = 0
	res_flag = 1
	res = make([][]byte,numRows,numRows)
	for index,char = range []byte(s) {
		res[res_index] = append(res[res_index], char)
		if index % (numRows - 1) == 0 && index != 0{
			res_flag = -res_flag
		}
		res_index += res_flag
	}
	buff = bytes.NewBuffer(nil)
	for _, v := range res {
		buff.Write(v)
	}
	return buff.String()
}

func Convert(s string, numRows int) string {
	return convert(s, numRows)
}