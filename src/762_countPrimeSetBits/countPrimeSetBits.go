package _62_countPrimeSetBits







var primes []bool = []bool{false,false,true,true,false,true, false ,true ,false ,false ,false ,true ,false ,true ,false ,false ,false ,true ,false ,true ,false ,false}

func isPrime(a int) bool {
	return primes[a]
}


func countPrimeSetBits(L int, R int) int {
	ans := 0
	for i := L; i <=R; i++ {
		tmp := i
		count := 0
		for tmp != 0 {
			if tmp & 1 == 1 {
				count++
			}
			tmp = tmp >> 1
		}
		if isPrime(count) {
			ans++
		}
	}
	return ans
}