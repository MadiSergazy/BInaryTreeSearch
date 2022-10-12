package main

import (
	"fmt"
	"os"
	CheckModeficators "piscine/CheckModificators"
	"piscine/Punctuation"
	"strings"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 2 || arguments[0] != "sample.txt" || arguments[1] != "result.txt" {
		fmt.Println("Incorrect arguments")
		return
	}
	if !strings.HasSuffix(arguments[1], ".txt") || !strings.HasSuffix(arguments[0], ".txt") { // check format file
		fmt.Println("Incorrect format in file don't txt")
		return
	}
	dat, err := os.ReadFile(arguments[0])
	check(err)

	var splText []string
	var positionModificators []int
	splText, positionModificators = splitText(string(dat))
	// fmt.Println(splText)
	// fmt.Println(positionModificators)
	if CheckModeficators.ModifierList(splText, positionModificators) {
		str := Punctuation.Punctuations(splText)
		err = os.WriteFile(arguments[1], []byte(str), 0644)
		check(err)
	}
	// fmt.Println(Article(splText))
	// ar := Article(splText)
	// fmt.Println(Punctuation(ar))
}

func splitText(s string) ([]string, []int) {
	var text []string
	var position []int
	var flag bool
	word := ""

	for index, value := range s {

		if value == '(' { // work with comand
			word += string(value)
			flag = true
			continue
		}
		if value == ')' {
			word += string(value)
			text = append(text, word)
			flag = false
			word = ""
			position = append(position, len(text)-1)
			continue
		}
		if value == 10 { // probel
			text = append(text, word)
			word = ""
			continue
		}
		if flag { //  add value inside ()
			word += string(value)
			continue
		}
		if value == '\n' {
			text = append(text, word) // add '\n'
			text = append(text, "\n")
			word = ""
			continue
		}
		// check ' except 0-9
		if value == '\'' || (value >= 58 && value <= 64) || (value >= 32 && value <= 47) || (value >= 123 && value <= 126) || (value >= 91 && value <= 96) { // add spetc simvol
			text = append(text, word)
			text = append(text, string(value))
			word = ""
			continue
		}
		if index == len(s)-1 { // add last letter
			word += string(value)
			text = append(text, word)
			continue
		}

		word += string(value) // default
	}

	return text, position
}

func check(err error) { // chec errors
	if err != nil {
		fmt.Println("Unable to create file or something went wrong:", err)
		fmt.Printf(err.Error())
		return
	}
}

/*
func Punctuation(s string) []string {
	// separation of commas
	for i := 0; i < len(s); i++ {
		if s[i] == '.' || s[i] == ',' || s[i] == '!' || s[i] == '?' || s[i] == ':' || s[i] == ';' || s[i] == '\'' &&
			(i != 0 && !('a' <= s[i-1] && s[i-1] <= 'z' || 'A' <= s[i-1] && s[i-1] <= 'Z') ||
				i != len(s)-1 && !('a' <= s[i+1] && s[i+1] <= 'z' || 'A' <= s[i+1] && s[i+1] <= 'Z')) {
			s = s[:i] + " " + string(s[i]) + " " + s[i+1:]
			i++
		}
	}
	// own strings.Fields with \n
	res := []string{}
	str := ""
	for i, v := range s {
		if v != ' ' {
			str += string(v)
		}
		if str != "" && (v == ' ' || i == len(s)-1 || v == '\n') {
			res = append(res, str)
			str = ""
		}
	}
	for i := 0; i < len(res)-1; i++ {
		if res[i] == "" {
			res = append(res[:i], res[i+1:]...)
		}
	}
	// punctuation
	res2 := []string{}
	comma := true
	for i := 0; i < len(res); i++ {
		if i != 0 && res[i] == "." || res[i] == "," || res[i] == "!" || res[i] == "?" || res[i] == ":" || res[i] == ";" {
			res2[len(res2)-1] += res[i]
		} else if res[i] == "'" {
			if comma {
				comma = false
				if i != len(res)-1 {
					res[i+1] = "'" + res[i+1]
				}
			} else {
				comma = true
				if i != 0 {
					res2[len(res2)-1] = res2[len(res2)-1] + "'"
				}
			}
		} else {
			res2 = append(res2, res[i])
		}
	}
	return res2
}*/
