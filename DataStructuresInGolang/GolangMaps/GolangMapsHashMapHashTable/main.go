package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Println(m == nil)
	m = map[string]string{}
	fmt.Println(m == nil)
	fmt.Println(m)
	fmt.Println(len(m))
	m = make(map[string]string, 5)
	fmt.Println(len(m))
	m = map[string]string{"Walse": "Programmer"}
	fmt.Println(m)
	fmt.Println(len(m))
	m["Jonny"] = "PROGER"
	fmt.Println(m)
	m["Jonny"] = " NO PROGER"
	fmt.Println(m)
	fmt.Println(m["Jonny"])
	delete(m, "Jonny")
	fmt.Println(m)
	m["Walse"] += "Ninga"
	fmt.Println(m)

	for name, title := range m {
		fmt.Println(name, title)
	}

	fmt.Println(m["Jonny"])

	title, ok := m["Jonny"]
	if ok {
		fmt.Println(title)
	} else {
		fmt.Println("Did not find Jonny")
	}

}
