package main

import (
	"Mado/Checker"
	"Mado/FileReadWork"
	"Mado/PrintArt"
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	//fmt.Println(arguments)
	args := arguments[0]
	//fmt.Println(args)

	fileName := arguments[1]

	switch fileName {
	case "standard":
		fileName = fileName + ".txt"
	case "shadow":
		fileName = fileName + ".txt"
	case "thinkertoy":
		fileName = fileName + ".txt"
	default:
		fmt.Printf("Incorect args %v\n", fileName)
		return
	}

	massivString, errF := FileReadWork.FileReadWork(fileName) // reading file
	if errF == -1 {
		return
	}

	if !Checker.Check(args) { // check correct args
		return
	}

	errNewLine, count := Checker.ChekcOnlyN(args) // if we have only \n
	if !errNewLine {
		for i := 0; i < count; i++ {
			fmt.Println()
		}
		return
	}
	//var text string
	errP := PrintArt.PrintArt(massivString, args) // print and check err
	//fmt.Print(text)
	if errP == -1 {
		return
	}

}

/*
	//2 variant
		scanner2 := bufio.NewScanner(os.Stdin) // READING CONSOLE
		sozder := ""
		for {
			fmt.Print("Enter Text: ")
			// reads user input until \n by default
			scanner2.Scan()
			// Holds the string that was scanned
			text := scanner2.Text()
			if len(text) != 0 {
				// fmt.Println(text)
				sozder += text
			} else {
				// exit if user entered an empty string
				break
			}

		}

		// handle error
		if scanner2.Err() != nil {
			fmt.Println("Error: ", scanner2.Err())
		}
*/
// fmt.Println(args)
