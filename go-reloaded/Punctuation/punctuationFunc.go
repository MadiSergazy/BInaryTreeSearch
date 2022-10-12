package Punctuation

func Punctuations(s1 []string) string {
	str := ""
	count := 0
	for index, val := range s1 {
		if val == "A" || val == "a" {
			upper := 1
			if index+upper < len(s1) {
				for i := 1; i <= upper; i++ {
					if s1[index+i] == "\n" || s1[index+i] == "" {
						upper++
						continue
					}
					if isVowel(s1[index+i]) {
						if s1[index] == "A" {
							s1[index] = "An"
							continue
						} else {
							s1[index] = "an"
							continue
						}
					}
				}
			}
		}
		if val == "An" || val == "an" {
			upper := 1
			if index+upper < len(s1) {
				for i := 1; i <= upper; i++ {
					if s1[index+i] == "\n" || s1[index+i] == "" {
						upper++
						continue
					}
					if !isVowel(s1[index+i]) {
						if s1[index] == "An" {
							s1[index] = "A"
							continue
						} else {
							s1[index] = "a"
							continue
						}
					}
				}
			}
		}
		if val == "'" {
			count++
			if count%2 == 1 {
				str += " " + val
				continue
			} else {
				str += val
				continue
			}
		}
		if len(val) == 1 {
			if val >= string(33) && val <= string(47) || val >= string(58) && val <= string(64) || val >= string(91) &&
				val == string(96) || val >= string(123) && val == string(126) {
				str += val
				continue
			}
		}
		if val == "" {
			continue
		}
		if val == "\n" {
			str += "\n"
			continue
		}
		/*to fix*/
		if len(str) >= 1 || index == len(s1)-1 {
			if str != "" {
				if str[len(str)-1:] == "\n" || str[len(str)-1:] == "'" {
					str += s1[index]
				} else {
					str += " " + s1[index]
				}
			} else {
				str += s1[index]
			}
		} else {
			str += " " + s1[index]
		}
	}
	if str == "" {
		return str
	} else {
		if str[0] == ' ' {
			return str[1:]
		}
		return str
	}
}
func isVowel(str string) bool {
	arrVowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'h', 'H'}
	for _, val := range arrVowels {
		if val == str[0] {
			return true
		}
	}
	return false
}
