package src

func romanToInt(s string) (res int) {
	for i, c := range s {
		switch c {
		case 'I':
			res += 1
		case 'V':
			if i != 0 && s[i-1] == 'I' {
				res += 3
			} else {
				res += 5
			}

		case 'X':
			if i != 0 && s[i-1] == 'I' {
				res += 8
			} else {
				res += 10
			}

		case 'L':
			if i != 0 && s[i-1] == 'X' {
				res += 30
			} else {
				res += 50
			}

		case 'C':
			if i != 0 && s[i-1] == 'X' {
				res += 80
			} else {
				res += 100
			}

		case 'D':
			if i != 0 && s[i-1] == 'C' {
				res += 300
			} else {
				res += 500
			}

		case 'M':
			if i != 0 && s[i-1] == 'C' {
				res += 800
			} else {
				res += 1000
			}
		default:
		}
	}
	return
}
