package main

import (
	"strconv"
	"strings"
	"unicode"
)

func parseArrayRune(array []rune) string {
	var stringBuilder strings.Builder

	const StartSearchPosition = 1
	arrayLength := len(array)

	for index := StartSearchPosition; index < arrayLength; index++ {
		numberStartPosition, isStartOk := searchNumberStartPosition(array, arrayLength, index)

		if !isStartOk {
			break
		}

		numberEndPosition := searchNumberEndPosition(array, arrayLength, numberStartPosition+1)

		if array[index-1] == '\\' {

		} else {
			numberFromString := string(array[numberStartPosition : numberEndPosition+1])

			number, err := strconv.Atoi(numberFromString)

			if err != nil {
				panic(err)
			}

			unpackedString := strings.Repeat(string(array[index-1]), number)

			stringBuilder.WriteString(unpackedString)
		}

		index = numberEndPosition
	}

	if stringBuilder.Len() == 0 {
		return string(array)
	}

	return stringBuilder.String()
}

func searchNumberStartPosition(array []rune, arrayLength int, searchFrom int) (startPosition int, isOk bool) {
	for index := searchFrom; index < arrayLength; index++ {
		if unicode.IsDigit(array[index]) {
			return index, true
		}
	}
	return
}

func searchNumberEndPosition(array []rune, arrayLength int, searchFrom int) int {
	for index := searchFrom; index < arrayLength; index++ {
		if !unicode.IsDigit(array[index]) {
			return index
		}
	}
	return arrayLength - 1
}
