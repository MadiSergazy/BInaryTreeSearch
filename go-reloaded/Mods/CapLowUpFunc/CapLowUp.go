package CapLowUpFunction

import "strings"

func Up(text []string, index, num int) {
	if index < num {
		num = index
	}
	for i := 1; i <= num; i++ {
		if index-i >= 0 {
			if text[index-i] == "" || text[index-i] == "\n" || (text[index-i] >= string(33) &&
				text[index-i] <= string(64)) || text[index-i] >= string(91) && text[index-i] <= string(96) ||
				(text[index-i] >= string(123) && text[index-i] <= string(126)) {
				num++
			}
			text[index-i] = strings.ToUpper(text[index-i])
		}
	}
	text[index] = ""
}

func Low(text []string, index, num int) {
	if index < num {
		num = index
	}
	for i := 1; i <= num; i++ {
		if index-i >= 0 {
			if text[index-i] == "" || text[index-i] == "\n" || (text[index-i] >= string(33) &&
				text[index-i] <= string(64)) || text[index-i] >= string(91) && text[index-i] <= string(96) ||
				(text[index-i] >= string(123) && text[index-i] <= string(126)) {
				num++
			}
			text[index-i] = strings.ToLower(text[index-i])
		}
	}
	text[index] = ""
}

func Cap(text []string, index, num int) {
	if index < num {
		num = index
	}
	for i := 1; i <= num; i++ {
		if index-i >= 0 {
			if text[index-i] == "" || text[index-i] == "\n" || (text[index-i] >= string(33) &&
				text[index-i] <= string(64)) || text[index-i] >= string(91) && text[index-i] <= string(96) ||
				(text[index-i] >= string(123) && text[index-i] <= string(126)) {
				num++
			}
			text[index-i] = strings.ToLower(text[index-i])
			text[index-i] = strings.Title(text[index-i])
		}
	}
	text[index] = ""
}
