package parsing

import (
	"strconv"
	"strings"
	"unicode"
)

/*
Parser structure for unpack string
*/
type Parser struct {
	array         []rune
	stringBuilder *strings.Builder
}

/*
NewParser is the Parser constructor
*/
func NewParser(array []rune) *Parser {
	return &Parser{
		array:         array,
		stringBuilder: &strings.Builder{},
	}
}

/*
Parse returns the unpacked string
*/
func (receiver *Parser) Parse() string {
	arrayLength := len(receiver.array)

	for index := 0; index < arrayLength; index++ {
		switch {
		case receiver.isEscapeSymbol(index):
			nextIndex := index + 1

			if receiver.isEscapeSymbol(nextIndex) {
				receiver.convertCountAndRepeatSymbol(receiver.array[index+2]+1, receiver.array[nextIndex])
				index += 2
			} else {
				receiver.stringBuilder.WriteRune(receiver.array[nextIndex])
				index++
			}
		case index > 0 && unicode.IsDigit(receiver.array[index]):
			receiver.convertCountAndRepeatSymbol(receiver.array[index], receiver.array[index-1])
		default:
			receiver.stringBuilder.WriteRune(receiver.array[index])
		}
	}

	return receiver.stringBuilder.String()
}

func (receiver *Parser) isEscapeSymbol(index int) bool {
	return receiver.array[index] == '\\'
}

func (receiver *Parser) convertCountAndRepeatSymbol(count, symbol rune) {
	countValue, err := strconv.Atoi(string(count))

	if err != nil {
		panic(err)
	}

	output := strings.Repeat(string(symbol), countValue-1)
	receiver.stringBuilder.WriteString(output)
}
