package __StringToInteger

func myAtoi(str string) int {
	const (
		MaxInt32  = 1<<31 - 1
		MinInt32  = -1 << 31
	)
	var (
		result    int
		spaceFlag,numFlag bool
		fixFlag	  byte
		char 	  byte
		charMap   map[byte]int
	)
	charMap = map[byte]int{'0':0,'1':1,'2':2,'3':3,'4':4,'5':5,'6':6,'7':7,'8':8,'9':9}
	spaceFlag = false
	fixFlag = 0
	result = 0
	for _, char = range []byte(str) {
		if char == ' ' {
			if spaceFlag {
				break
			}
			continue
		} else if char == '-'{
			if fixFlag != 0 || numFlag {
				break
			}
			fixFlag = 1
			spaceFlag = true
			continue
		} else if char == '+'{
			if fixFlag != 0  || numFlag{
				break
			}
			fixFlag = 2
			spaceFlag = true
			continue
		} else if char > '9' || char < '0' {
			break
		}
		spaceFlag = true
		numFlag = true
		if fixFlag == 1 && (-result < MinInt32 / 10 || (-result == MinInt32 / 10 && charMap[char] >= 8)){
			return MinInt32
		}else if fixFlag != 1 && result > MaxInt32 / 10 || (result == MaxInt32 / 10 && charMap[char] > 7){
			return MaxInt32
		}
		result = result * 10 + charMap[char]

	}
	if fixFlag == 1 {
		result = -result
	}
	return result
}

func MyAtoi(str string) int  {
	return myAtoi(str)
}

