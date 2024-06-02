package main

import (
	"strconv"
	"strings"
	"unicode"
)

func parseArrayRune(array []rune, arrayLength int) string {
	var stringBuilder strings.Builder

	startCopy := 0
	endCopy := 0

	for index := startCopy + 1; index < arrayLength; index++ {
		numberStartPosition, isStartOk := searchNumberStartPosition(array, arrayLength, index)

		if !isStartOk {
			endCopy = arrayLength
			break
		}

		endCopy = numberStartPosition - 1
		copyAndWriteString(&stringBuilder, array, startCopy, endCopy)

		numberEndPosition := searchNumberEndPosition(array, arrayLength, numberStartPosition+1)

		if array[endCopy] == '\\' {

		} else {
			convertAndRepeatSymbol(&stringBuilder, array, numberStartPosition, numberEndPosition, endCopy)
		}

		startCopy = numberEndPosition
		index = numberEndPosition - 1
	}

	if startCopy < endCopy {
		copyAndWriteString(&stringBuilder, array, startCopy, endCopy)
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

func copyAndWriteString(stringBuilder *strings.Builder, array []rune, startCopy int, endCopy int) {
	copyString := string(array[startCopy:endCopy])
	stringBuilder.WriteString(copyString)
}

func searchNumberEndPosition(array []rune, arrayLength int, searchFrom int) int {
	for index := searchFrom; index < arrayLength; index++ {
		if !unicode.IsDigit(array[index]) {
			return index
		}
	}
	return arrayLength
}

func convertAndRepeatSymbol(
	stringBuilder *strings.Builder,
	array []rune,
	numberStartPosition int,
	numberEndPosition int,
	endCopy int,
) {
	numberFromString := string(array[numberStartPosition:numberEndPosition])
	number, err := strconv.Atoi(numberFromString)

	if err != nil {
		panic(err)
	}

	unpackedString := strings.Repeat(string(array[endCopy]), number)
	stringBuilder.WriteString(unpackedString)
}
