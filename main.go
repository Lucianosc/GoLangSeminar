package main

import (
	"fmt"
	"root/model"
	"strconv"
	"unicode"
)

func main() {

	str1 := "TX03ABC"
	str2 := "NN055B634"

	result1 := parseStruct(str1)
	result2 := parseStruct(str2)

	fmt.Println(result1)
	fmt.Println(result2)

}

func parseStruct(str string) *model.Result {
	typ := str[0:2]
	length := parseInt(str[2:4])
	value := str[4:]

	if len(value) == length {
		if typ == "NN" && IsNumber(value) {
			return model.NewResult(typ, length, value)
		}
		if typ == "TX" && IsLetter(value) {
			return model.NewResult(typ, length, value)
		}
	}

	obj := new(model.Result)
	return obj
}

func IsLetter(str string) bool {
	for _, char := range str {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func IsNumber(str string) bool {
	for _, char := range str {
		if !unicode.IsNumber(char) {
			return false
		}
	}
	return true
}

func parseInt(str string) int {
	number, _ := strconv.Atoi(str)
	return number
}
