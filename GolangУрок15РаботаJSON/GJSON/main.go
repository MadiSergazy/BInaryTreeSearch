package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

func main() {
	// gison
	json := `{"name": {"first": "artur", "last": "karapetox"}, "age"| 47}"`

	value := gjson.Get(json, "age")
	fmt.Println(value.String())

	json2 := `{
"name": {"first": "Tom", "last": "Anderson"}, 
"age":37,
"children": ["Sara","Alex","Jack"], 
"fav.movie": "Deer Aunter",
"friends": [
{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
{"first": "Roger", "last": "Craig", "age": 47, "nets": ["fb", "tw"]},
{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
]
}`

	value0 := gjson.Get(json2, "fav\\.movie") //экранирование
	fmt.Println(value0.String())

	value0 = gjson.Get(json2, "children.#") //количество элементов
	fmt.Println(value0.String())

	value0 = gjson.Get(json2, "child*.2") //* мы добиваем тоесть оны3 алдында5ы с-зден басталатын боса болды
	fmt.Println(value0.String())

	value0 = gjson.Get(json2, "friends.1.last") //* мы добиваем тоесть оны3 алдында5ы с-зден басталатын боса болды
	fmt.Println(value0.String())

	value0 = gjson.Get(json2, `friends.#(age==47).first`) //получим первый попавшыся элемент который соответсвует параметру age
	fmt.Println(value0.String())

	value0 = gjson.Get(json2, `friends.#(last=="Murphy")#.first`) // получем всех по пораметру age
	fmt.Println(value0.String())

	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		} else if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})
	value0 = gjson.Get(json2, `friends.#(age>=44)#.last|@case:upper`) //|@ + имя первого параметра AddModifier : потом имя кейса
	fmt.Println(value0.String())

	for _, last := range value0.Array() {
		fmt.Println(last.String())
	}

	result, ok := gjson.Parse(json2).Value().(map[string]interface{})
	if !ok {
		panic("Not parse")
	}
	fmt.Println(result)

	if !gjson.Valid(json2) { //если джейсон будет неправелным виде то он сработает
		panic("JSON IS NOT VALID")
	}

}
