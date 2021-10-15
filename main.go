package main

import (
	"errors"
	"root/model"
	"strconv"
	"unicode"
)

func main() {

	// str1 := "TX03ABC"
	// str2 := "NN0556784"

	// result1, err1 := ParseStruct(str1)
	// result2, err2 := ParseStruct(str2)

	// fmt.Println(result1, err1)
	// fmt.Println(result2, err2)

}

func ParseStruct(str string) (*model.Result, error) {
	typ := str[0:2]
	length := parseInt(str[2:4])
	value := str[4:]

	if len(value) == length {
		if typ == "NN" {
			if IsNumber(value) {
				return model.NewResult(typ, length, value), nil
			} else {
				return new(model.Result), errors.New("cadena tipo NN con caracteres no numericos")
			}
		} else if typ == "TX" {
			if IsLetter(value) {
				return model.NewResult(typ, length, value), nil
			} else {
				return new(model.Result), errors.New("cadena tipo TX con caracteres no alfabeticos")
			}
		} else {
			return new(model.Result), errors.New("tipo de cadena invalida")
		}
	}
	return new(model.Result), errors.New("largo de cadena invalida")
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
