package CheckModeficators

import (
	"fmt"
	BinHexFunction "piscine/Mods/BinHexFunc"
	CapLowUpFunction "piscine/Mods/CapLowUpFunc"
	"strconv"
)

func error(text []string, pos int) {
	fmt.Println("NOT VALID MODIF: %v in this position: %v", text[pos], pos)
}

type toDo struct { // what do you do doing
	comand         int
	number         int
	positionOnText int
}

func ModifierList(textSlice []string, indexSlice []int) bool {
	arrLessFive := []string{"(up)", "(low)", "(cap)", "(hex)", "(bin)"}
	arrMoreFive := []string{"(up,", "(low", "(cap"}
	for _, val := range indexSlice {
		link := &toDo{}
		if len(textSlice[val]) > 5 {
			findcommand := false
			for indexOfCommand, nameOfCommand := range arrMoreFive {
				if nameOfCommand == textSlice[val][:4] && (textSlice[val][4] == ' ' || textSlice[val][4] == ',') {
					var start int
					var flag bool
					findcommand = true
					/*to understand it is (up, <number>) or other*/
					if indexOfCommand == 0 {
						start = 4
					} else {
						start = 5
					}
					/*check whether the modifier is valid or not*/
					if textSlice[val][start] == ',' || textSlice[val][start] == ' ' {
						start++
						toNumber := ""
						/*modifier number processing*/
						for !flag {
							if textSlice[val][len(textSlice[val])-1] != ')' {
								error(textSlice, val)
								return false
							}
							if textSlice[val][start] != ')' {
								toNumber += string(textSlice[val][start])
								start++
							} else {
								flag = true
							}
						}
						number, err := strconv.Atoi(toNumber)
						if err != nil {
							error(textSlice, val)
							return false
						}
						if number <= 0 {
							error(textSlice, val)
							return false
						}
						/*number, command, index*/
						link = &toDo{indexOfCommand, number, val}
					} else {
						error(textSlice, val)
						return false
					}
				}
			}
			if !findcommand {
				error(textSlice, val)
				return false
			}
		} else {
			count := 0
			for indexOfCommand, nameOfCommand := range arrLessFive {
				if nameOfCommand == textSlice[val] {
					link = &toDo{indexOfCommand, 1, val}
					count++
					break
				}
			}
			if count == 0 {
				error(textSlice, val)
				return false
			}
		}
		if !switchboard(textSlice, link) {
			return false
		}
	}
	return true
}
func switchboard(text []string, link *toDo) bool {
	switch link.comand {
	case 0:
		CapLowUpFunction.Up(text, link.positionOnText, link.number)
	case 1:
		CapLowUpFunction.Low(text, link.positionOnText, link.number)
	case 2:
		CapLowUpFunction.Cap(text, link.positionOnText, link.number)
	case 3:
		if BinHexFunction.Hex(text, link.positionOnText, link.number) {
			return false
		}
	case 4:
		if BinHexFunction.Bin(text, link.positionOnText, link.number) {
			return false
		}
	}
	return true
}

/*
func ModifierList(textSlice []string, indexSlice []int) bool {
	arrFiveModificator := []string{"(up)", "(low)", "(bin)", "(hex)", "(cap)"}
	arrFiveMoreModificator := []string{"(up,", "(low,", "(cap,"}

	for _, val := range indexSlice {
		link := &toDo{}
		if len(textSlice[val]) > 5 { // It is exzacly (comand,
			findcommand := false
			for indexOfCommand, nameOfCommand := range arrFiveMoreModificator {
				if nameOfCommand == textSlice[val][:4] && (textSlice[val][4] == ' ' || textSlice[val][4] == ',') {
					var start int
					var flag bool
					findcommand = true
					//to understand it is (up, <number>) or other
					if indexOfCommand == 0 { // if up
						start = 4
					} else { // if other
						start = 5
					}

					//check whether the modifier is valid or not//
					if textSlice[val][start] == ',' || textSlice[val][start] == ' ' {
						start++
						toNumber := ""

						//modifier number processing//
						for !flag {
							if textSlice[val][len(textSlice[val])-1] != ')' {
								error(textSlice, val)
								return false
							}
							if textSlice[val][start] != ')' {
								toNumber += string(textSlice[val][start])
								start++
							} else {
								flag = true
							}
						}
						number, err := strconv.Atoi(toNumber)
						if err != nil {
							error(textSlice, val)
							return false
						}
						if number <= 0 {
							error(textSlice, val)
							return false
						}
						//number, command, index//
						link = &toDo{indexOfCommand, number, val}
					} else {
						error(textSlice, val)
						return false
					}
				}
			}
			if !findcommand {
				error(textSlice, val)
				return false
			}
		} else {
			count := 0
			for indexOfCommand, nameOfCommand := range arrFiveModificator {
				if nameOfCommand == textSlice[val] {
					link = &toDo{indexOfCommand, 1, val}
					count++
					break
				}
			}
			if count == 0 {
				error(textSlice, val)
				return false
			}
		}
		if !switchboard(textSlice, link) {
			return false
		}
	}
	return true
}

func switchboard(text []string, link *toDo) bool { // call function
	switch link.comand {
	case 0:

		CapLowUpFunction.Up(text, link.positionOnText, link.number)
	case 1:
		CapLowUpFunction.Low(text, link.positionOnText, link.number)
	case 2:
		CapLowUpFunction.Cap(text, link.positionOnText, link.number)

	case 3:
		if BinHexFunction.Hex(text, link.positionOnText, link.number) {
			return false
		}
	case 4:
		if BinHexFunction.Bin(text, link.positionOnText, link.number) {
			return false
		}
	}
	return true
}
*/
