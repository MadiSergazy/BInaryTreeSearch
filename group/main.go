package main

import (
	"fmt"
)

func main() {
	//s := "aaaaaa"
	s := "abbaca"
	fmt.Println(removeDuplicates(s))
}

func removeDuplicates(s string) string {
	b := []byte(s)
	for i := len(b) - 1; i > 0; i-- {
		if i == len(b) {
			continue
		}
		if b[i] == b[i-1] {
			b = append(b[:i-1], b[i+1:]...)
		}
	}
	return string(b)
}

/*
func removeDuplicates(s string) string {
	for i := len(s) - 1; i >= 1; i-- {
		if s[i] == s[i-1] {
			//ab := s[:i-1]
			s = s[:i]
		}

	}
	return s
}*/

/*
	urls := URLs{}
	GetData("https://groupietrackers.herokuapp.com/api", &urls)

	fmt.Println(urls)
}

func GetData(url string, in interface{}) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fmt.Println("response")
	fmt.Println(response)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	/*  []bytes
	fmt.Println("responseData")
	fmt.Println(responseData)
*/
//
//	fmt.Println("STRING responseData")
//	fmt.Println(string(responseData))
//
//	err = json.Unmarshal(responseData, in)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}
//
//type URLs struct {
//	ArtistsUrl   string `json:"artists"`
//	LocationsUrl string `json:"locations"`
//	DatesUrl     string `json:"dates"`
//	RelationUrl  string `json:"relation"`
//}*/
