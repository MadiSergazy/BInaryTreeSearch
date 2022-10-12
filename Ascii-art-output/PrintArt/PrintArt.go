package PrintArt

import (
	"fmt"
	"strings"
)

func PrintArt(massivString []string, args string) (string, int) {
	splitArgs := strings.Split(args, "\\n") // sozder mojno vmesto args
	str := ""
	for _, arg := range splitArgs {

		if arg == "" {
			fmt.Println()
			str += "\n"
			continue
		}

		for i := 0; i < 8; i++ {
			for index, v := range arg {
				if v < 32 || v > 126 {
					fmt.Println("DONT CORRECT ARGS IN %d", index)
					return str, -1
				}
				//fmt.Print(strings.Split(massivString[v-32], "\n")[i])
				str += strings.Split(massivString[v-32], "\n")[i]
			}
			//fmt.Println()
			str += "\n"
		}

	}
	//fmt.Println(str)
	return str, 1
}
