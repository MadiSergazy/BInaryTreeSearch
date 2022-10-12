package main

import (
	"Mado/Checker"
	"Mado/FileReadWork"
	"Mado/PrintArt"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	text, errP := PrintArt.PrintArt(massivString, args) // print and check err
	//fmt.Print(text)
	if errP == -1 {
		return
	}

	FileNameToWrite := strings.Split(arguments[2], "=")
	if FileNameToWrite[0] != "--output" {
		fmt.Println("Incorrect args in writeFile")
		return
	}

	fileExtensionWrite := filepath.Ext(FileNameToWrite[1])
	if fileExtensionWrite != ".txt" { // проверка формата файла
		fmt.Println("INCORRECT FILEANME for write")
		return
	}

	f, err := os.Create(FileNameToWrite[1])
	if err != nil {
		fmt.Println("ERROR WHEN CREATE FILE")
		return
		// log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(text)

	if err2 != nil {
		fmt.Println("ERROR WHEN write FILE")
		return
		// log.Fatal(err2)
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
