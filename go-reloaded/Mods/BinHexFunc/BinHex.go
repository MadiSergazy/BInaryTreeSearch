package BinHexFunction

import (
	"fmt"
	"strconv"
)

func Hex(text []string, index, num int) bool {
	if index < num {
		num = index
	}
	for i := 1; i <= num; i++ {
		if index-i >= 0 {
			if numPP(text[index-i], 0) {
				num++
				continue
			}
			number, err := strconv.ParseInt(text[index-i], 16, 64)
			if err != nil {
				fmt.Printf("Not valied arguments change it - %v before (hex)\nPosition in your text - "+
					"%v", text[index-i], index-i)
				fmt.Println()
				return true
			}
			text[index-i] = strconv.FormatInt(number, 10)
		}
	}
	text[index] = ""
	return false
}
func Bin(text []string, index, num int) bool {
	if index < num {
		num = index
	}
	for i := 1; i <= num; i++ {
		if index-i >= 0 {
			if numPP(text[index-i], 0) {
				num++
				continue
			}
			number, err := strconv.ParseInt(text[index-i], 2, 64)
			if err != nil {
				fmt.Printf("Not valied arguments change it - %v before (bin)\nPosition in your text - "+
					"%v", text[index-i], index-i)
				fmt.Println()
				return true
			}
			text[index-i] = strconv.FormatInt(number, 10)
		}
	}
	text[index] = ""
	return false
}
func numPP(s1 string, n int) bool {
	switch n {
	case 0:
		if s1 == "" || s1 == "\n" || s1 >= string(33) && s1 <= string(47) || s1 >= string(58) &&
			s1 <= string(64) || s1 >= string(91) && s1 <= string(96) || s1 >= string(123) && s1 <= string(126) {
			return true
		} else {
			return false
		}
	case 1:
		if s1 == "" || s1 == "\n" || s1 >= string(33) && s1 <= string(64) || s1 >= string(91) && s1 <= string(96) ||
			s1 >= string(123) && s1 <= string(126) {
			return true
		} else {
			return false
		}
	}
	return false
}
