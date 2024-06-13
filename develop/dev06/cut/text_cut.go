package cut

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
TextCut structure
*/
type TextCut struct {
	dto *TextCutFlagDTO
}

/*
NewTextCut constructor
*/
func NewTextCut(dto *TextCutFlagDTO) *TextCut {
	return &TextCut{dto}
}

/*
Cut method
*/
func (receiver *TextCut) Cut(inputRows []string) []string {
	fields := receiver.convertValuesOfFieldsFlag()

	var outputRows []string

	for _, inputRow := range inputRows {
		splitRows := strings.Split(inputRow, receiver.dto.DelimiterFlag)

		splitRowsLength := len(splitRows)

		var outputRow string

		if splitRowsLength > 1 {
			for index, fieldNumber := range fields {
				if fieldNumber >= splitRowsLength {
					continue
				}

				if index == 0 {
					outputRow = splitRows[fieldNumber]
					continue
				}
				outputRow += receiver.dto.DelimiterFlag + splitRows[fieldNumber]
			}
			outputRows = append(outputRows, outputRow)
			continue
		}

		if !receiver.dto.SeparatedFlag {
			outputRows = append(outputRows, inputRow)
		}
	}

	return outputRows
}

func (receiver *TextCut) convertValuesOfFieldsFlag() []int {
	stringArray := strings.Split(receiver.dto.FieldsFlag, ",")

	var integerArray []int

	for _, value := range stringArray {
		number, err := strconv.Atoi(value)

		if err != nil {
			fmt.Println("parameter -f error")
			os.Exit(1)
		}

		integerArray = append(integerArray, number-1)
	}

	slices.Sort(integerArray)

	return integerArray
}
