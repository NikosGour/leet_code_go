package src

import (
	"math"
)

func longestCommonPrefix(strs []string) string {
	res := ""
	min_len := math.MaxInt
	for _, str := range strs {
		min_len = min(min_len, len(str))
	}

outer:
	for i := range min_len {
		tmp := strs[0][0 : i+1]
		for _, str := range strs {
			if str[0:i+1] != tmp {
				break outer
			}
		}
		res = tmp
	}

	return res
}
