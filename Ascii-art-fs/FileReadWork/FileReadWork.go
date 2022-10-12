package FileReadWork

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func FileReadWork(fileName string) ([]string, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	text := ""
	counter, counterForappendSlice := 0, 0
	var massivString []string
	for scanner.Scan() {

		// fmt.Println(len(scanner.Text()))
		if len(scanner.Text()) == 0 {
			continue
		}
		counter++
		counterForappendSlice++
		text += string(scanner.Text())
		if counter != 760 {
			text += "\n"
		}
		if counterForappendSlice%8 == 0 {
			massivString = append(massivString, text)
			text = ""
		}

		// fmt.Println(scanner.Text())

	}
	// fmt.Println(massivString)

	if err := scanner.Err(); err != nil {
		fmt.Println("SCANNER ERROR")
		return massivString, -1
		//	log.Fatal(err)
	}

	return massivString, 1
}
