package main

import (
	"strconv"
	"strings"
	"unicode"
)

type Parsing struct {
	array         []rune
	arrayLength   int
	stringBuilder *strings.Builder
	startCopy     int
	endCopy       int
}

func NewParsing(array []rune) *Parsing {
	return &Parsing{
		array:         array,
		arrayLength:   len(array),
		stringBuilder: &strings.Builder{},
	}
}

func (receiver *Parsing) ParseArrayRune() string {
	for index := receiver.startCopy + 1; index < receiver.arrayLength; index++ {
		numberStartPosition, isStartOk := receiver.searchNumberStartPosition(index)

		if !isStartOk {
			receiver.endCopy = receiver.arrayLength
			break
		}

		receiver.endCopy = numberStartPosition - 1
		receiver.copyAndWriteString()

		numberEndPosition := receiver.searchNumberEndPosition(numberStartPosition + 1)

		if receiver.array[receiver.endCopy] == '\\' {

		} else {
			receiver.convertAndRepeatSymbol(numberStartPosition, numberEndPosition)
		}

		receiver.startCopy = numberEndPosition
		index = numberEndPosition - 1
	}

	if receiver.startCopy < receiver.endCopy {
		receiver.copyAndWriteString()
	}

	return receiver.stringBuilder.String()
}

func (receiver *Parsing) searchNumberStartPosition(searchFrom int) (startPosition int, isOk bool) {
	for index := searchFrom; index < receiver.arrayLength; index++ {
		if unicode.IsDigit(receiver.array[index]) {
			return index, true
		}
	}
	return
}

func (receiver *Parsing) copyAndWriteString() {
	copyString := string(receiver.array[receiver.startCopy:receiver.endCopy])
	receiver.stringBuilder.WriteString(copyString)
}

func (receiver *Parsing) searchNumberEndPosition(searchFrom int) int {
	for index := searchFrom; index < receiver.arrayLength; index++ {
		if !unicode.IsDigit(receiver.array[index]) {
			return index
		}
	}
	return receiver.arrayLength
}

func (receiver *Parsing) convertAndRepeatSymbol(numberStartPosition int, numberEndPosition int) {
	numberFromString := string(receiver.array[numberStartPosition:numberEndPosition])
	number, err := strconv.Atoi(numberFromString)

	if err != nil {
		panic(err)
	}

	unpackedString := strings.Repeat(string(receiver.array[receiver.endCopy]), number)
	receiver.stringBuilder.WriteString(unpackedString)
}
