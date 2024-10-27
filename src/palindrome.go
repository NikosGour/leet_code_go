package src

import (
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	i := 0

	for i < len(str)/2 {
		if str[i] != str[len(str)-1-i] {
			return false
		}
		i++
	}

	return true
}
