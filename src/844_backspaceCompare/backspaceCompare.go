package _44_backspaceCompare



func backspaceCompare(S string, T string) bool {
	s_q := len(S) - 1
	t_q := len(T) - 1
	var s_char byte
	var t_char byte
	s_j_count := 0
	t_j_count := 0
	i := s_q
	j := t_q
	for {
		for ; i >=0; i-- {
			if S[i] != '#' {
				if  s_j_count == 0 {
					s_char = S[i]
					i--
					break
				} else {
					i -= s_j_count
					s_j_count = 0
				}
			} else {
				s_j_count++
			}

		}
		for ; j >=0; j-- {
			if T[j] != '#' {
				if  t_j_count == 0 {
					t_char = T[j]
					j--
					break
				} else {
					j++
					j -= t_j_count
					t_j_count = 0
				}
			} else {
				t_j_count++
			}

		}
		if i <=0 && j <= 0 {
			return true
		}
		if s_char != t_char {
			return false
		}

	}
}