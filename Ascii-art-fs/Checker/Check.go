package Checker

import "fmt"

func Check(args string) bool {
	if len(args) == 0 {
		// fmt.Println("EMPTY ARGS")
		return false
	}
	for index, v := range args {
		if !(v >= 32 && v <= 126) {
			fmt.Printf("INVALID ARGS in index %d\n", index)
			return false
		}
	}
	return true
}

func ChekcOnlyN(args string) (bool, int) {
	counter := 0
	flag := true
	// if len(args) != 1 {

	for i := 0; i < len(args)-1; i++ {
		if args[i] == '\\' && args[i+1] == 'n' {
			counter += 2
		}
		if args[i] != '\\' && args[i] != 'n' {
			flag = false
		}
	}

	if counter == len(args) && flag {
		return false, counter / 2
	} else {
		return true, counter
	}
} // else {
//	return true, counter
//}
//}
