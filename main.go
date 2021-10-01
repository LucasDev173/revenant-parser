package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Result struct {
	Type   string
	Value  string
	Length int
}

func parse(input string) (Result, error) {
	var parsedInput Result

	inputLenght := len([]rune(input))
	if inputLenght < 5 {
		return parsedInput, errors.New("cadena no valida: es demasiado corta para generar el struct (tamaño minimo: 5)")
	}
	parsedInput.Type = input[:2]
	x, e := strconv.Atoi(input[2:4])
	if e != nil {
		return parsedInput, errors.New("cadena no valida: 3 y 4to caracter debe representar largo")
	}
	parsedInput.Length = x
	parsedInput.Value = input[4:]

	fmt.Printf("%+v\n", parsedInput)
	return parsedInput, nil
}
