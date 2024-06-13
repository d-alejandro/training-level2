package main

import (
	"bufio"
	"d-alejandro/training-level2/develop/dev06/cut"
	"flag"
	"fmt"
	"os"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fieldsFlag := flag.String("f", "1", "select only these fields; by default selects by field 1")
	delimiterFlag := flag.String("d", "\t", "use field delimiter")
	separatedFlag := flag.Bool("s", false, "do not print lines not containing delimiters")

	fmt.Println("Please enter your text:")

	var inputRows []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		inputRows = append(inputRows, text)
	}

	if len(inputRows) == 0 {
		return
	}

	flag.Parse()

	textCutFlagDTO := cut.TextCutFlagDTO{
		FieldsFlag:    *fieldsFlag,
		DelimiterFlag: *delimiterFlag,
		SeparatedFlag: *separatedFlag,
	}

	textCut := cut.NewTextCut(&textCutFlagDTO)
	response := textCut.Cut(inputRows)

	fmt.Println("Response:")

	for _, row := range response {
		fmt.Println(row)
	}
}
